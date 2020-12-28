package generated

import (
	"github.com/vmware/differential-datalog/go/pkg/ddlog"
)

// Generated code for relation AccountResult
// DO NOT CHANGE MANUALLY

// input relation AccountResult [AccountResult]

var (
	relTableIDAccountResult ddlog.TableID = ddlog.GetTableID("AccountResult")
)

func init() {
	relTableIDAccountResult = ddlog.GetTableID("AccountResult")
	meta := &TableMetaData{
		TableID: relTableIDAccountResult,
		TableName: "AccountResult", 
		RecordName: "AccountResult",
		Unmarshaller: func(record ddlog.Record) (interface{}, error) {
			obj, err := AccountResultFromRecord(record)
			return obj, err
		},
	}
	registerTableMetaData(relTableIDAccountResult, meta)
}

func NewInsertCommandAccountResult(obj *AccountResult) ddlog.Command {
	rec := NewRecordAccountResult(obj)
	return ddlog.NewInsertCommand(relTableIDAccountResult, rec)
}

func NewInsertOrUpdateCommandAccountResult(obj *AccountResult) ddlog.Command {
	rec := NewRecordAccountResult(obj)
	return ddlog.NewInsertOrUpdateCommand(relTableIDAccountResult, rec)
}

func NewDeleteValCommandAccountResult(obj *AccountResult) ddlog.Command {
	rec := NewRecordAccountResult(obj)
	return ddlog.NewDeleteValCommand(relTableIDAccountResult, rec)
}
func NewDeleteKeyCommandAccountResult(obj *AccountResult) ddlog.Command {
	rec := NewKeyRecordAccountResult(obj)
	return ddlog.NewDeleteKeyCommand(relTableIDAccountResult, rec)
}


func NewKeyRecordAccountResult(obj *AccountResult) ddlog.Record {
	return func() ddlog.Record {
	    return ddlog.NewRecordString(obj.CredentialsHash)
    }()
}

func GetRelTableIDAccountResult() ddlog.TableID {
	return relTableIDAccountResult
}
