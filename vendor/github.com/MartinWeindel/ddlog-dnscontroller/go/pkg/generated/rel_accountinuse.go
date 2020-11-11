package generated

import (
	"github.com/vmware/differential-datalog/go/pkg/ddlog"
)

// Generated code for relation AccountInUse
// DO NOT CHANGE MANUALLY

// output relation AccountInUse [AccountInUse]

var (
	relTableIDAccountInUse ddlog.TableID = ddlog.GetTableID("AccountInUse")
)

func init() {
	relTableIDAccountInUse = ddlog.GetTableID("AccountInUse")
	meta := &TableMetaData{
		TableID: relTableIDAccountInUse,
		TableName: "AccountInUse", 
		RecordName: "AccountInUse",
		Unmarshaller: func(record ddlog.Record) (interface{}, error) {
			obj, err := AccountInUseFromRecord(record)
			return obj, err
		},
	}
	registerTableMetaData(relTableIDAccountInUse, meta)
}

func NewInsertCommandAccountInUse(obj *AccountInUse) ddlog.Command {
	rec := NewRecordAccountInUse(obj)
	return ddlog.NewInsertCommand(relTableIDAccountInUse, rec)
}

func NewInsertOrUpdateCommandAccountInUse(obj *AccountInUse) ddlog.Command {
	rec := NewRecordAccountInUse(obj)
	return ddlog.NewInsertOrUpdateCommand(relTableIDAccountInUse, rec)
}

func NewDeleteValCommandAccountInUse(obj *AccountInUse) ddlog.Command {
	rec := NewRecordAccountInUse(obj)
	return ddlog.NewDeleteValCommand(relTableIDAccountInUse, rec)
}

func GetRelTableIDAccountInUse() ddlog.TableID {
	return relTableIDAccountInUse
}
