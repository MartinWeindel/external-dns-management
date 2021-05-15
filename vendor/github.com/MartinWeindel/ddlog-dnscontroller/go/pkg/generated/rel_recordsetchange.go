package generated

import (
	"github.com/vmware/differential-datalog/go/pkg/ddlog"
)

// Generated code for relation RecordSetChange
// DO NOT CHANGE MANUALLY

// output relation RecordSetChange [RecordSetChange]

func init() {
	meta := &TableMetaData{
		TableName: "RecordSetChange", 
		RecordName: "RecordSetChange",
		Unmarshaller: func(record ddlog.Record) (interface{}, error) {
			obj, err := RecordSetChangeFromRecord(record)
			return obj, err
		},
	}
	registerTableMetaData(meta)
}

func (pd *ProgData) NewInsertCommandRecordSetChange(obj *RecordSetChange) ddlog.Command {
	rec := NewRecordRecordSetChange(obj)
	tableID := pd.LookupTableID("RecordSetChange")
	return ddlog.NewInsertCommand(tableID, rec)
}

func (pd *ProgData) NewInsertOrUpdateCommandRecordSetChange(obj *RecordSetChange) ddlog.Command {
	rec := NewRecordRecordSetChange(obj)
	tableID := pd.LookupTableID("RecordSetChange")
	return ddlog.NewInsertOrUpdateCommand(tableID, rec)
}

func (pd *ProgData) NewDeleteValCommandRecordSetChange(obj *RecordSetChange) ddlog.Command {
	rec := NewRecordRecordSetChange(obj)
	tableID := pd.LookupTableID("RecordSetChange")
	return ddlog.NewDeleteValCommand(tableID, rec)
}

func (pd *ProgData) GetRelTableIDRecordSetChange() ddlog.TableID {
	return pd.LookupTableID("RecordSetChange")
}
