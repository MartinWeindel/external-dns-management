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

func NewInsertOrUpdateCommandRecordSet(obj *RecordSet) ddlog.Command {
	rec := NewRecordRecordSet(obj)
	return ddlog.NewInsertOrUpdateCommand(relTableIDRecordSet, rec)
}

func NewDeleteValCommandRecordSet(obj *RecordSet) ddlog.Command {
	rec := NewRecordRecordSet(obj)
	return ddlog.NewDeleteValCommand(relTableIDRecordSet, rec)
}
func NewDeleteKeyCommandRecordSet(obj *RecordSet) ddlog.Command {
	rec := NewKeyRecordRecordSet(obj)
	return ddlog.NewDeleteKeyCommand(relTableIDRecordSet, rec)
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

func GetRelTableIDRecordSet() ddlog.TableID {
	return relTableIDRecordSet
}
