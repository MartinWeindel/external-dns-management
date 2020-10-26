package generated

import (
	"github.com/vmware/differential-datalog/go/pkg/ddlog"
)

// Generated code for relation RecordSet
// DO NOT CHANGE MANUALLY

// input relation RecordSet [RecordSet]

var (
	relTableIDRecordSet ddlog.TableID = ddlog.GetTableID("RecordSet")
)

func init() {
	relTableIDRecordSet = ddlog.GetTableID("RecordSet")
	meta := &TableMetaData{
		TableID: relTableIDRecordSet,
		TableName: "RecordSet", 
		RecordName: "RecordSet",
		Unmarshaller: func(record ddlog.Record) (interface{}, error) {
			obj, err := RecordSetFromRecord(record)
			return obj, err
		},
	}
	registerTableMetaData(relTableIDRecordSet, meta)
}

func NewInsertCommandRecordSet(obj *RecordSet) ddlog.Command {
	rec := NewRecordRecordSet(obj)
	return ddlog.NewInsertCommand(relTableIDRecordSet, rec)
}
