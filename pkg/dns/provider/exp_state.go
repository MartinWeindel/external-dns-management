package provider

import (
	"crypto/sha256"
	"fmt"
	"sort"
	"sync"

	"github.com/MartinWeindel/ddlog-dnscontroller/go/pkg/generated"
	"github.com/gardener/controller-manager-library/pkg/controllermanager/controller"
	"github.com/gardener/controller-manager-library/pkg/logger"
	"github.com/gardener/controller-manager-library/pkg/resources"
	"github.com/gardener/controller-manager-library/pkg/utils"
	api "github.com/gardener/external-dns-management/pkg/apis/dns/v1alpha1"
	dnsutils "github.com/gardener/external-dns-management/pkg/dns/utils"
	"github.com/vmware/differential-datalog/go/pkg/ddlog"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/klog"
)

type expState struct {
	setup   *setup
	context Context
	config  Config
	classes *controller.Classes

	credentials map[string]utils.Properties

	accountCache *AccountCache

	zones        map[string]*dnsHostedZone
	zonesAccount map[string]string

	lock              sync.Mutex
	knownProviders    resources.ObjectNameSet
	knownEntries      resources.ObjectNameSet
	knownOwners       resources.ObjectNameSet
	ddlogCommandQueue []ddlog.Command

	numDDlogWorkers   uint
	ddlogProgram      *ddlog.Program
	progData          *generated.ProgData
	ddlogUpdateActive int32
	outRecordHandler  outRecordHandler
}

func newExpState(context Context, classes *controller.Classes, config Config) *expState {
	return &expState{
		setup:             newSetup(),
		context:           context,
		config:            config,
		numDDlogWorkers:   1,
		classes:           classes,
		credentials:       map[string]utils.Properties{},
		accountCache:      NewAccountCache(config.CacheTTL, config.CacheDir, config.Options),
		zones:             map[string]*dnsHostedZone{},
		zonesAccount:      map[string]string{},
		knownProviders:    resources.ObjectNameSet{},
		knownEntries:      resources.ObjectNameSet{},
		knownOwners:       resources.ObjectNameSet{},
		ddlogCommandQueue: []ddlog.Command{},
	}
}

func (s *expState) Setup() error {
	// Ensures that DDlog will use our own logger (klog) to print error messages.
	log := func(msg string) {
		klog.Errorf(msg)
	}

	ddlog.SetErrMsgPrinter(log)

	s.outRecordHandler = newOutRecordHandler()

	klog.Infof("Running new DDlog program")
	var err error
	s.ddlogProgram, err = ddlog.NewProgram(s.numDDlogWorkers, &s.outRecordHandler)
	if err != nil {
		return err
	}
	s.progData = s.outRecordHandler.GetProgData(s.ddlogProgram)

	owner := &generated.DNSOwner{Name: "commandline:identifier", OwnerId: s.config.Ident, Active: true}
	cmd := s.progData.NewInsertOrUpdateCommandDNSOwner(owner)
	s.addToDDLogCommandQueue(cmd)

	s.setupFor(&api.DNSProvider{}, "providers", func(e resources.Object) {
		p := dnsutils.DNSProvider(e)
		s.UpdateProvider(s.context.NewContext("provider", p.ObjectName().String()), p)
	}, 1)

	s.triggerDDLogUpdate()

	return nil
}

func (s *expState) setupFor(obj runtime.Object, msg string, exec func(resources.Object), processors int) {
	s.context.Infof("### setup %s", msg)
	res, _ := s.context.GetByExample(obj)
	list, _ := res.ListCached(labels.Everything())
	dnsutils.ProcessElements(list, func(e resources.Object) {
		if s.IsResponsibleFor(s.context, e) {
			exec(e)
		}
	}, processors)
}

func (s *expState) Start() {
	s.setup.Start(s.context)
	s.setup = nil
}

func (s *expState) IsResponsibleFor(logger logger.LogContext, obj resources.Object) bool {
	return s.classes.IsResponsibleFor(logger, obj)
}

func (s *expState) GetDNSAccount(logger logger.LogContext, provider *dnsutils.DNSProviderObject, props utils.Properties) (*DNSAccount, error) {
	account, new, err := s.accountCache.Get(logger, provider, props, s)
	if err != nil {
		return nil, err
	}
	if new {
		s.triggerAccount(account.Hash())
	}
	return account, err
}

func (s *expState) GetContext() Context {
	return s.context
}

func (s *expState) GetConfig() Config {
	return s.config
}

func (s *expState) hashAndCache(secretProps utils.Properties) string {
	keys := []string{}
	for k := range secretProps {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	h := sha256.New()
	for _, k := range keys {
		h.Write([]byte(k))
		h.Write([]byte(secretProps[k]))
	}
	hash := fmt.Sprintf("%x", h.Sum(nil))

	s.credentials[hash] = secretProps
	return hash
}

func (s *expState) enqueueObject(clusterID string, gk schema.GroupKind, objKey generated.ObjectKey) {
	key := resources.NewClusterKey(s.context.GetClusterId(clusterID), gk, objKey.Arg0, objKey.Arg1)
	s.enqueueKey(key)
}

func (s *expState) enqueueKey(key resources.ClusterObjectKey) {
	err := s.context.EnqueueKey(key)
	if err != nil {
		s.context.Errorf("enqueue key %s failed: %s\n", key, err)
	} else {
		s.context.Infof("enqueued key %s\n", key)
	}
}
