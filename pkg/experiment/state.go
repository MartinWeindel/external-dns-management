package experiment

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"

	"github.com/MartinWeindel/ddlog-dnscontroller/go/pkg/generated"
	"github.com/gardener/controller-manager-library/pkg/logger"
	"github.com/gardener/controller-manager-library/pkg/resources"
	"github.com/vmware/differential-datalog/go/pkg/ddlog"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog"
)

type Context interface {
	logger.LogContext

	GetContext() context.Context

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
}

type State struct {
	context          Context
	numDDlogWorkers  uint
	ddlogProgram     *ddlog.Program
	outRecordHandler ddlog.OutRecordHandler
}

type OutRecordPrinter struct {
	changesMutex sync.Mutex
}

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
	s, _ := json.MarshalIndent(obj, "", "  ")
	fmt.Printf("%s[%s] %d: %s\n", meta.TableName, meta.RecordName, weight, s)
}

func NewState(context Context) *State {
	return &State{
		context:         context,
		numDDlogWorkers: 1,
	}
}

func (s *State) Setup() error {
	// Ensures that DDlog will use our own logger (klog) to print error messages.
	log := func(msg string) {
		klog.Errorf(msg)
	}

	ddlog.SetErrMsgPrinter(log)

	s.outRecordHandler = &OutRecordPrinter{}

	klog.Infof("Running new DDlog program")
	var err error
	s.ddlogProgram, err = ddlog.NewProgram(s.numDDlogWorkers, s.outRecordHandler)
	if err != nil {
		return err
	}
	return nil
}
