package provider

import (
	"crypto/sha256"
	"encoding/json"
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
	"k8s.io/klog"
)

type expState struct {
	context          Context
	config           Config
	classes          *controller.Classes
	numDDlogWorkers  uint
	ddlogProgram     *ddlog.Program
	outRecordHandler ddlog.OutRecordHandler

	credentials map[string]utils.Properties

	accountCache *AccountCache

	lock                sync.Mutex
	providerUpdateQueue map[resources.ObjectName][]*generated.DNSProviderStatus
}

type expOutRecordPrinter struct {
	changesMutex sync.Mutex
	clusterId    string
	state        *expState
}

func (p *expOutRecordPrinter) Handle(tableID ddlog.TableID, r ddlog.Record, weight int64) {
	p.changesMutex.Lock()
	defer p.changesMutex.Unlock()

	meta, err := generated.LookupTableMetaData(tableID)
	if err != nil {
		println("lookup failed", tableID, err, r.Dump(), weight)
	}
	obj, err := meta.Unmarshaller(r)
	if err != nil {
		println("unmarshal failed", tableID, err, r.Dump(), weight)
	}
	if tableID == generated.GetRelTableIDDNSProviderStatus() {
		o := obj.(*generated.DNSProviderStatus)
		p.state.AddProviderStatusToQueueAndEnqueue(o)
	}
	s, _ := json.MarshalIndent(obj, "", "  ")
	fmt.Printf("%s[%s] %d: %s\n%s\n", meta.TableName, meta.RecordName, weight, s, r.Dump())
}

func newExpState(context Context, classes *controller.Classes, config Config) *expState {
	return &expState{
		context:             context,
		config:              config,
		numDDlogWorkers:     1,
		classes:             classes,
		credentials:         map[string]utils.Properties{},
		providerUpdateQueue: map[resources.ObjectName][]*generated.DNSProviderStatus{},
		accountCache:        NewAccountCache(config.CacheTTL, config.CacheDir, config.Options),
	}
}

func (s *expState) Setup() error {
	// Ensures that DDlog will use our own logger (klog) to print error messages.
	log := func(msg string) {
		klog.Errorf(msg)
	}

	ddlog.SetErrMsgPrinter(log)

	s.outRecordHandler = &expOutRecordPrinter{state: s}

	klog.Infof("Running new DDlog program")
	var err error
	s.ddlogProgram, err = ddlog.NewProgram(s.numDDlogWorkers, s.outRecordHandler)
	if err != nil {
		return err
	}

	s.setupFor(&api.DNSProvider{}, "providers", func(e resources.Object) {
		p := dnsutils.DNSProvider(e)
		s.UpdateProvider(s.context.NewContext("provider", p.ObjectName().String()), p)
	}, 1)

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

func (s *expState) IsResponsibleFor(logger logger.LogContext, obj resources.Object) bool {
	return s.classes.IsResponsibleFor(logger, obj)
}

func (s *expState) GetDNSAccount(logger logger.LogContext, provider *dnsutils.DNSProviderObject, props utils.Properties) (*DNSAccount, error) {
	return s.accountCache.Get(logger, provider, props, s)
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

func (s *expState) AddProviderStatusToQueueAndEnqueue(obj *generated.DNSProviderStatus) {
	s.addProviderStatusToQueue(obj)

	key := resources.NewClusterKey(s.context.GetClusterId(PROVIDER_CLUSTER), providerGroupKind, obj.Key.Arg0, obj.Key.Arg1)
	err := s.context.EnqueueKey(key)
	if err != nil {
		s.context.Errorf("enqueue key %s failed: %s\n", key, err)
	} else {
		s.context.Infof("enqueued key %s\n", key)
	}
}

func (s *expState) addProviderStatusToQueue(obj *generated.DNSProviderStatus) {
	name := resources.NewObjectName(obj.Key.Arg0, obj.Key.Arg1)

	s.lock.Lock()
	defer s.lock.Unlock()
	list := s.providerUpdateQueue[name]
	s.providerUpdateQueue[name] = append(list, obj)
}

func (s *expState) nextProviderStatusesFromQueue(name resources.ObjectName) []*generated.DNSProviderStatus {
	s.lock.Lock()
	defer s.lock.Unlock()
	list, ok := s.providerUpdateQueue[name]
	if !ok || len(list) == 0 {
		return nil
	}
	delete(s.providerUpdateQueue, name)
	return list
}
