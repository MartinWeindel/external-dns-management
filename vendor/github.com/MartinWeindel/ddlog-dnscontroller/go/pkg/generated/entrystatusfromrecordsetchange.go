package generated

// Generated code for typedef EntryStatusFromRecordSetChange
// DO NOT CHANGE MANUALLY

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/vmware/differential-datalog/go/pkg/ddlog"
)


var (
	// memory will never be freed, which is fine
	relConstructorEntryStatusFromRecordSetChange = ddlog.NewCString("EntryStatusFromRecordSetChange")
)


func NewRecordEntryStatusFromRecordSetChange(obj *EntryStatusFromRecordSetChange) ddlog.Record {
	arg0 := func() ddlog.Record {
	    return NewRecordProviderType(obj.Ptype)
    }()
	arg1 := func() ddlog.Record {
	    return ddlog.NewRecordString(obj.Zoneid)
    }()
	arg2 := func() ddlog.Record {
	    return NewRecordObjectKey(&obj.Key)
    }()
	arg3 := func() ddlog.Record {
	    return NewRecordEntryStatus(obj.Status)
    }()
	return ddlog.NewRecordStructStatic(relConstructorEntryStatusFromRecordSetChange, arg0, arg1, arg2, arg3)
}


func EntryStatusFromRecordSetChangeFromRecord(record ddlog.Record) (*EntryStatusFromRecordSetChange, error) {
	rs, err := record.AsStructSafe()
	if err != nil {
		return nil, errors.Wrapf(err, "no struct (EntryStatusFromRecordSetChange)")
	}
	if rs.Name() != "EntryStatusFromRecordSetChange" {
		return nil, fmt.Errorf("unexpected record %s != EntryStatusFromRecordSetChange", rs.Name())
	}
	arg0, err := ProviderTypeFromRecord(rs.At(0))
	if err != nil {
		return nil, errors.Wrapf(err, "Field ptype")
	}
	arg1, err := rs.At(1).ToStringSafe()
	if err != nil {
		return nil, errors.Wrapf(err, "Field zoneid")
	}
	arg2, err := ObjectKeyFromRecord(rs.At(2))
	if err != nil {
		return nil, errors.Wrapf(err, "Field key")
	}
	arg3, err := EntryStatusFromRecord(rs.At(3))
	if err != nil {
		return nil, errors.Wrapf(err, "Field status")
	}
	obj := &EntryStatusFromRecordSetChange{	
		Ptype:arg0,	
		Zoneid:arg1,	
		Key:*arg2,	
		Status:arg3,
	}
	return obj, nil
}

type EntryStatusFromRecordSetChange struct {
    Ptype ProviderType
    Zoneid string
    Key ObjectKey
    Status EntryStatus
}
