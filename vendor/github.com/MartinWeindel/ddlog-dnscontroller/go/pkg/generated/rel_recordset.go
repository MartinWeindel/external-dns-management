package generated

import (
	"github.com/vmware/differential-datalog/go/pkg/ddlog"
)

// Generated code for relation RecordSet
// DO NOT CHANGE MANUALLY

// input relation RecordSet [RecordSet]

func init() {
	meta := &TableMetaData{
		TableName: "RecordSet", 
		RecordName: "RecordSet",
		Unmarshaller: func(record ddlog.Record) (interface{}, error) {
			obj, err := RecordSetFromRecord(record)
			return obj, err
		},
	}
	registerTableMetaData(meta)
}

func (pd *ProgData) NewInsertCommandRecordSet(obj *RecordSet) ddlog.Command {
	rec := NewRecordRecordSet(obj)
	tableID := pd.LookupTableID("RecordSet")
	return ddlog.NewInsertCommand(tableID, rec)
}

func (pd *ProgData) NewInsertOrUpdateCommandRecordSet(obj *RecordSet) ddlog.Command {
	rec := NewRecordRecordSet(obj)
	tableID := pd.LookupTableID("RecordSet")
	return ddlog.NewInsertOrUpdateCommand(tableID, rec)
}

func (pd *ProgData) NewDeleteValCommandRecordSet(obj *RecordSet) ddlog.Command {
	rec := NewRecordRecordSet(obj)
	tableID := pd.LookupTableID("RecordSet")
	return ddlog.NewDeleteValCommand(tableID, rec)
}
func (pd *ProgData) NewDeleteKeyCommandRecordSet(obj *RecordSet) ddlog.Command {
	rec := NewKeyRecordRecordSet(obj)
	tableID := pd.LookupTableID("RecordSet")
	return ddlog.NewDeleteKeyCommand(tableID, rec)
}


func NewKeyRecordRecordSet(obj *RecordSet) ddlog.Record {
	arg0 := func() ddlog.Record {
	    return ddlog.NewRecordString(obj.Owner)
    }()
	arg1 := func() ddlog.Record {
	    return NewRecordProviderType(obj.Ptype)
    }()
	arg2 := func() ddlog.Record {
	    return ddlog.NewRecordString(obj.Zoneid)
    }()
	arg3 := func() ddlog.Record {
	    return ddlog.NewRecordString(obj.Domain)
    }()
	arg4 := func() ddlog.Record {
	    return NewRecordRecordType(obj.Rtype)
    }()
	return ddlog.NewRecordTuple(arg0,arg1,arg2,arg3,arg4)
}

func (pd *ProgData) GetRelTableIDRecordSet() ddlog.TableID {
	return pd.LookupTableID("RecordSet")
}
