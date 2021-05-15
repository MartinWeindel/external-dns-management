package generated

import (
	"github.com/vmware/differential-datalog/go/pkg/ddlog"
)

// Generated code for relation AccountInUse
// DO NOT CHANGE MANUALLY

// output relation AccountInUse [AccountInUse]

func init() {
	meta := &TableMetaData{
		TableName: "AccountInUse", 
		RecordName: "AccountInUse",
		Unmarshaller: func(record ddlog.Record) (interface{}, error) {
			obj, err := AccountInUseFromRecord(record)
			return obj, err
		},
	}
	registerTableMetaData(meta)
}

func (pd *ProgData) NewInsertCommandAccountInUse(obj *AccountInUse) ddlog.Command {
	rec := NewRecordAccountInUse(obj)
	tableID := pd.LookupTableID("AccountInUse")
	return ddlog.NewInsertCommand(tableID, rec)
}

func (pd *ProgData) NewInsertOrUpdateCommandAccountInUse(obj *AccountInUse) ddlog.Command {
	rec := NewRecordAccountInUse(obj)
	tableID := pd.LookupTableID("AccountInUse")
	return ddlog.NewInsertOrUpdateCommand(tableID, rec)
}

func (pd *ProgData) NewDeleteValCommandAccountInUse(obj *AccountInUse) ddlog.Command {
	rec := NewRecordAccountInUse(obj)
	tableID := pd.LookupTableID("AccountInUse")
	return ddlog.NewDeleteValCommand(tableID, rec)
}

func (pd *ProgData) GetRelTableIDAccountInUse() ddlog.TableID {
	return pd.LookupTableID("AccountInUse")
}
