package generated

import (
	"github.com/vmware/differential-datalog/go/pkg/ddlog"
)

// Generated code for relation AccountResult
// DO NOT CHANGE MANUALLY

// input relation AccountResult [AccountResult]

func init() {
	meta := &TableMetaData{
		TableName: "AccountResult", 
		RecordName: "AccountResult",
		Unmarshaller: func(record ddlog.Record) (interface{}, error) {
			obj, err := AccountResultFromRecord(record)
			return obj, err
		},
	}
	registerTableMetaData(meta)
}

func (pd *ProgData) NewInsertCommandAccountResult(obj *AccountResult) ddlog.Command {
	rec := NewRecordAccountResult(obj)
	tableID := pd.LookupTableID("AccountResult")
	return ddlog.NewInsertCommand(tableID, rec)
}

func (pd *ProgData) NewInsertOrUpdateCommandAccountResult(obj *AccountResult) ddlog.Command {
	rec := NewRecordAccountResult(obj)
	tableID := pd.LookupTableID("AccountResult")
	return ddlog.NewInsertOrUpdateCommand(tableID, rec)
}

func (pd *ProgData) NewDeleteValCommandAccountResult(obj *AccountResult) ddlog.Command {
	rec := NewRecordAccountResult(obj)
	tableID := pd.LookupTableID("AccountResult")
	return ddlog.NewDeleteValCommand(tableID, rec)
}
func (pd *ProgData) NewDeleteKeyCommandAccountResult(obj *AccountResult) ddlog.Command {
	rec := NewKeyRecordAccountResult(obj)
	tableID := pd.LookupTableID("AccountResult")
	return ddlog.NewDeleteKeyCommand(tableID, rec)
}


func NewKeyRecordAccountResult(obj *AccountResult) ddlog.Record {
	return func() ddlog.Record {
	    return ddlog.NewRecordString(obj.CredentialsHash)
    }()
}

func (pd *ProgData) GetRelTableIDAccountResult() ddlog.TableID {
	return pd.LookupTableID("AccountResult")
}
