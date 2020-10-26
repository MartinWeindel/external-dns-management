package generated

import (
	"github.com/vmware/differential-datalog/go/pkg/ddlog"
)

// Generated code for relation RecordSetChange
// DO NOT CHANGE MANUALLY

// output relation RecordSetChange [RecordSetChange]

var (
	relTableIDRecordSetChange ddlog.TableID = ddlog.GetTableID("RecordSetChange")
)

func init() {
	relTableIDRecordSetChange = ddlog.GetTableID("RecordSetChange")
	meta := &TableMetaData{
		TableID: relTableIDRecordSetChange,
		TableName: "RecordSetChange", 
		RecordName: "RecordSetChange",
		Unmarshaller: func(record ddlog.Record) (interface{}, error) {
			obj, err := RecordSetChangeFromRecord(record)
			return obj, err
		},
	}
	registerTableMetaData(relTableIDRecordSetChange, meta)
}

func NewInsertCommandRecordSetChange(obj *RecordSetChange) ddlog.Command {
	rec := NewRecordRecordSetChange(obj)
	return ddlog.NewInsertCommand(relTableIDRecordSetChange, rec)
}
