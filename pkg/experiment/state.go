package experiment

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"sort"
	"sync"

	corev1 "k8s.io/api/core/v1"

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

type Context interface {
	logger.LogContext

	GetContext() context.Context

	GetClusterId(name string) string

	IsReady() bool
	GetByExample(runtime.Object) (resources.Interface, error)

	GetStringOption(name string) (string, error)
	GetIntOption(name string) (int, error)

	Synchronize(log logger.LogContext, name string, initiator resources.Object) (bool, error)

	Enqueue(obj resources.Object) error
	EnqueueCommand(cmd string) error
	EnqueueKey(key resources.ClusterObjectKey) error

	SetFinalizer(resources.Object) error
	RemoveFinalizer(resources.Object) error
	HasFinalizer(resources.Object) bool

	GetSecretPropertiesByRef(src resources.ResourcesSource, ref *corev1.SecretReference) (utils.Properties, *resources.SecretObject, error)
}

type State struct {
	context          Context
	classes          *controller.Classes
	numDDlogWorkers  uint
	ddlogProgram     *ddlog.Program
	outRecordHandler ddlog.OutRecordHandler

	credentials map[string]utils.Properties
}

type OutRecordPrinter struct {
	changesMutex sync.Mutex
	clusterId    string
	context      Context
}

const PROVIDER_CLUSTER = "provider"

var ownerGroupKind = resources.NewGroupKind(api.GroupName, api.DNSOwnerKind)
var providerGroupKind = resources.NewGroupKind(api.GroupName, api.DNSProviderKind)
var entryGroupKind = resources.NewGroupKind(api.GroupName, api.DNSEntryKind)

func (p *OutRecordPrinter) Handle(tableID ddlog.TableID, r ddlog.Record, weight int64) {
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
		key := resources.NewClusterKey(p.context.GetClusterId(PROVIDER_CLUSTER), providerGroupKind, o.Key.Arg0, o.Key.Arg1)
		err := p.context.EnqueueKey(key)
		if err != nil {
			fmt.Printf("enqueue key %s failed: %s\n", key, err)
		} else {
			fmt.Printf("enqueued key %s\n", key)
		}
	}
	s, _ := json.MarshalIndent(obj, "", "  ")
	fmt.Printf("%s[%s] %d: %s\n%s\n", meta.TableName, meta.RecordName, weight, s, r.Dump())
}

func NewState(context Context, classes *controller.Classes) *State {
	return &State{
		context:         context,
		numDDlogWorkers: 1,
		classes:         classes,
		credentials:     map[string]utils.Properties{},
	}
}

func (s *State) Setup() error {
	// Ensures that DDlog will use our own logger (klog) to print error messages.
	log := func(msg string) {
		klog.Errorf(msg)
	}

	ddlog.SetErrMsgPrinter(log)

	s.outRecordHandler = &OutRecordPrinter{context: s.context}

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

func (s *State) setupFor(obj runtime.Object, msg string, exec func(resources.Object), processors int) {
	s.context.Infof("### setup %s", msg)
	res, _ := s.context.GetByExample(obj)
	list, _ := res.ListCached(labels.Everything())
	dnsutils.ProcessElements(list, func(e resources.Object) {
		if s.IsResponsibleFor(s.context, e) {
			exec(e)
		}
	}, processors)
}

func (s *State) IsResponsibleFor(logger logger.LogContext, obj resources.Object) bool {
	return s.classes.IsResponsibleFor(logger, obj)
}

func (s *State) hashAndCache(secretProps utils.Properties) string {
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
