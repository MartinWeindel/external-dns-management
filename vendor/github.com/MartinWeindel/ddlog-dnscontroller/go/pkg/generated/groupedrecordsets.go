package generated

// Generated code for typedef GroupedRecordSets
// DO NOT CHANGE MANUALLY

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/vmware/differential-datalog/go/pkg/ddlog"
)


var (
	// memory will never be freed, which is fine
	relConstructorGroupedRecordSets = ddlog.NewCString("GroupedRecordSets")
)


func NewRecordGroupedRecordSets(obj *GroupedRecordSets) ddlog.Record {
	arg0 := func() ddlog.Record {
	    return NewRecordProviderType(obj.Ptype)
    }()
	arg1 := func() ddlog.Record {
	    return ddlog.NewRecordString(obj.Zoneid)
    }()
	arg2 := func() ddlog.Record {
	    vec := make([]ddlog.Record, len(obj.RecordSets))
	for i, item := range obj.RecordSets {
		vec[i] = NewRecordRecordSetData(&item)
	}
    return ddlog.NewRecordVector(vec...)
    }()
	return ddlog.NewRecordStructStatic(relConstructorGroupedRecordSets, arg0, arg1, arg2)
}


func GroupedRecordSetsFromRecord(record ddlog.Record) (*GroupedRecordSets, error) {
	rs, err := record.AsStructSafe()
	if err != nil {
		return nil, errors.Wrapf(err, "no struct (GroupedRecordSets)")
	}
	if rs.Name() != "GroupedRecordSets" {
		return nil, fmt.Errorf("unexpected record %s != GroupedRecordSets", rs.Name())
	}
	arg0, err := ProviderTypeFromRecord(rs.At(0))
	if err != nil {
		return nil, errors.Wrapf(err, "Field ptype")
	}
	arg1, err := rs.At(1).ToStringSafe()
	if err != nil {
		return nil, errors.Wrapf(err, "Field zoneid")
	}
	arg2, err := 
	func() ([]RecordSetData, error) {
		rv0 := rs.At(2)
		rv, err := rv0.AsVectorSafe()
		if err != nil {
			return nil, err
		}
		vec := make([]RecordSetData, rv.Size())
		for i := 0; i < len(vec); i++ {
			obj, err := RecordSetDataFromRecord(rv.At(i))
			if err != nil {
				errors.Wrapf(err, "vector index %d", i)
			}
			vec[i] =*obj
		}
		return vec, nil
	}()
	if err != nil {
		return nil, errors.Wrapf(err, "Field recordSets")
	}
	obj := &GroupedRecordSets{	
		Ptype:arg0,	
		Zoneid:arg1,	
		RecordSets:arg2,
	}
	return obj, nil
}

type GroupedRecordSets struct {
    Ptype ProviderType
    Zoneid string
    RecordSets []RecordSetData
}
