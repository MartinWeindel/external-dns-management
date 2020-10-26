package experiment

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/MartinWeindel/ddlog-dnscontroller/go/pkg/generated"
	"github.com/vmware/differential-datalog/go/pkg/ddlog"
	"k8s.io/klog"
)

func log(msg string) {
	klog.Errorf(msg)
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

var (
	numDDlogWorkers  uint = 1
	ddlogProgram     *ddlog.Program
	outRecordHandler ddlog.OutRecordHandler
)

func init() {
	// Ensures that DDlog will use our own logger (klog) to print error messages.
	ddlog.SetErrMsgPrinter(log)

	outRecordHandler = &OutRecordPrinter{}

	klog.Infof("Running new DDlog program")
	var err error
	ddlogProgram, err = ddlog.NewProgram(numDDlogWorkers, outRecordHandler)
	if err != nil {
		panic(err)
	}
}
