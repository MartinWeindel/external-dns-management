package generated

// Generated code for typedef GroupedEntries
// DO NOT CHANGE MANUALLY

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/vmware/differential-datalog/go/pkg/ddlog"
)


var (
	// memory will never be freed, which is fine
	relConstructorGroupedEntries = ddlog.NewCString("GroupedEntries")
)


func NewRecordGroupedEntries(obj *GroupedEntries) ddlog.Record {
	arg0 := func() ddlog.Record {
	    return NewRecordProviderType(obj.Ptype)
    }()
	arg1 := func() ddlog.Record {
	    return ddlog.NewRecordString(obj.Zoneid)
    }()
	arg2 := func() ddlog.Record {
	    vec := make([]ddlog.Record, len(obj.Entries))
	for i, item := range obj.Entries {
		vec[i] = func() ddlog.Record {
	    
	
	tuple := make([]ddlog.Record, 2)
	tuple[0] = NewRecordObjectKey(&item.Arg0)
	tuple[1] = NewRecordEntrySpec(&item.Arg1)
	return ddlog.NewRecordTuple(tuple...)
    }()
	}
    return ddlog.NewRecordVector(vec...)
    }()
	return ddlog.NewRecordStructStatic(relConstructorGroupedEntries, arg0, arg1, arg2)
}


func GroupedEntriesFromRecord(record ddlog.Record) (*GroupedEntries, error) {
	rs, err := record.AsStructSafe()
	if err != nil {
		return nil, errors.Wrapf(err, "no struct (GroupedEntries)")
	}
	if rs.Name() != "GroupedEntries" {
		return nil, fmt.Errorf("unexpected record %s != GroupedEntries", rs.Name())
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
	func() ([]Tuple_ObjectKey_EntrySpec, error) {
		rv0 := rs.At(2)
		rv, err := rv0.AsVectorSafe()
		if err != nil {
			return nil, err
		}
		vec := make([]Tuple_ObjectKey_EntrySpec, rv.Size())
		for i := 0; i < len(vec); i++ {
			obj, err := 
	func() (*Tuple_ObjectKey_EntrySpec, error) {
		rt, err := rv.At(i).AsTupleSafe()
		if err != nil {
			return nil, err
		}
		obj := &Tuple_ObjectKey_EntrySpec{}
		
		
		a0, err := ObjectKeyFromRecord(rt.At(0))
		if err != nil {
			return nil, errors.Wrap(err, "tuple index 0")
		}
		obj.Arg0 =*a0
		
		
		a1, err := EntrySpecFromRecord(rt.At(1))
		if err != nil {
			return nil, errors.Wrap(err, "tuple index 1")
		}
		obj.Arg1 =*a1
		return obj, nil
	}()
			if err != nil {
				errors.Wrapf(err, "vector index %d", i)
			}
			vec[i] =*obj
		}
		return vec, nil
	}()
	if err != nil {
		return nil, errors.Wrapf(err, "Field entries")
	}
	obj := &GroupedEntries{	
		Ptype:arg0,	
		Zoneid:arg1,	
		Entries:arg2,
	}
	return obj, nil
}

type GroupedEntries struct {
    Ptype ProviderType
    Zoneid string
    Entries []Tuple_ObjectKey_EntrySpec
}
