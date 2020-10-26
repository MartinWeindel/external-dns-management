package generated

// Generated code for typedef PState
// DO NOT CHANGE MANUALLY

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/vmware/differential-datalog/go/pkg/ddlog"
)


var (
	// memory will never be freed, which is fine
	relConstructorPState = ddlog.NewCString("PState")
)


func NewRecordPState(obj *PState) ddlog.Record {
	arg0 := func() ddlog.Record {
	    return NewRecordProviderType(obj.Ptype)
    }()
	arg1 := func() ddlog.Record {
	    return NewRecordIncludedExcluded(&obj.Domains)
    }()
	arg2 := func() ddlog.Record {
	    return NewRecordIncludedExcluded(&obj.Zoneids)
    }()
	arg3 := func() ddlog.Record {
	    vec := make([]ddlog.Record, len(obj.Zones))
	for i, item := range obj.Zones {
		vec[i] = NewRecordZone(&item)
	}
    return ddlog.NewRecordVector(vec...)
    }()
	return ddlog.NewRecordStructStatic(relConstructorPState, arg0, arg1, arg2, arg3)
}


func PStateFromRecord(record ddlog.Record) (*PState, error) {
	rs, err := record.AsStructSafe()
	if err != nil {
		return nil, errors.Wrapf(err, "no struct (PState)")
	}
	if rs.Name() != "PState" {
		return nil, fmt.Errorf("unexpected record %s != PState", rs.Name())
	}
	arg0, err := ProviderTypeFromRecord(rs.At(0))
	if err != nil {
		return nil, errors.Wrapf(err, "Field ptype")
	}
	arg1, err := IncludedExcludedFromRecord(rs.At(1))
	if err != nil {
		return nil, errors.Wrapf(err, "Field domains")
	}
	arg2, err := IncludedExcludedFromRecord(rs.At(2))
	if err != nil {
		return nil, errors.Wrapf(err, "Field zoneids")
	}
	arg3, err := 
	func() ([]Zone, error) {
		rv, err := rs.At(3).AsVectorSafe()
		if err != nil {
			return nil, err
		}
		vec := make([]Zone, rv.Size())
		for i := 0; i < len(vec); i++ {
			obj, err := ZoneFromRecord(rv.At(i))
			if err != nil {
				errors.Wrapf(err, "vector index %d", i)
			}
			vec[i] =*obj
		}
		return vec, nil
	}()
	if err != nil {
		return nil, errors.Wrapf(err, "Field zones")
	}
	obj := &PState{	
		Ptype:arg0,	
		Domains:*arg1,	
		Zoneids:*arg2,	
		Zones:arg3,
	}
	return obj, nil
}

type PState struct {
    Ptype ProviderType
    Domains IncludedExcluded
    Zoneids IncludedExcluded
    Zones []Zone
}
