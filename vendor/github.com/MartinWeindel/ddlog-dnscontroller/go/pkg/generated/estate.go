package generated

// Generated code for typedef EState
// DO NOT CHANGE MANUALLY

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/vmware/differential-datalog/go/pkg/ddlog"
)


var (
	// memory will never be freed, which is fine
	relConstructorEState = ddlog.NewCString("EState")
)


func NewRecordEState(obj *EState) ddlog.Record {
	arg0 := func() ddlog.Record {
	    return NewRecordProviderType(obj.Ptype)
    }()
	arg1 := func() ddlog.Record {
	    return ddlog.NewRecordString(obj.Zoneid)
    }()
	arg2 := func() ddlog.Record {
	    return ddlog.NewRecordString(obj.Owner)
    }()
	arg3 := func() ddlog.Record {
	    return ddlog.NewRecordString(obj.Domain)
    }()
	arg4 := func() ddlog.Record {
	    return ddlog.NewRecordU32(obj.Ttl)
    }()
	arg5 := func() ddlog.Record {
	    vec := make([]ddlog.Record, len(obj.Records))
	for i, item := range obj.Records {
		vec[i] = ddlog.NewRecordString(item)
	}
    return ddlog.NewRecordVector(vec...)
    }()
	return ddlog.NewRecordStructStatic(relConstructorEState, arg0, arg1, arg2, arg3, arg4, arg5)
}


func EStateFromRecord(record ddlog.Record) (*EState, error) {
	rs, err := record.AsStructSafe()
	if err != nil {
		return nil, errors.Wrapf(err, "no struct (EState)")
	}
	if rs.Name() != "EState" {
		return nil, fmt.Errorf("unexpected record %s != EState", rs.Name())
	}
	arg0, err := ProviderTypeFromRecord(rs.At(0))
	if err != nil {
		return nil, errors.Wrapf(err, "Field ptype")
	}
	arg1, err := rs.At(1).ToStringSafe()
	if err != nil {
		return nil, errors.Wrapf(err, "Field zoneid")
	}
	arg2, err := rs.At(2).ToStringSafe()
	if err != nil {
		return nil, errors.Wrapf(err, "Field owner")
	}
	arg3, err := rs.At(3).ToStringSafe()
	if err != nil {
		return nil, errors.Wrapf(err, "Field domain")
	}
	arg4, err := rs.At(4).ToU32Safe()
	if err != nil {
		return nil, errors.Wrapf(err, "Field ttl")
	}
	arg5, err := 
	func() ([]string, error) {
		rv0 := rs.At(5)
		rv, err := rv0.AsVectorSafe()
		if err != nil {
			return nil, err
		}
		vec := make([]string, rv.Size())
		for i := 0; i < len(vec); i++ {
			obj, err := rv.At(i).ToStringSafe()
			if err != nil {
				errors.Wrapf(err, "vector index %d", i)
			}
			vec[i] =obj
		}
		return vec, nil
	}()
	if err != nil {
		return nil, errors.Wrapf(err, "Field records")
	}
	obj := &EState{	
		Ptype:arg0,	
		Zoneid:arg1,	
		Owner:arg2,	
		Domain:arg3,	
		Ttl:arg4,	
		Records:arg5,
	}
	return obj, nil
}

type EState struct {
    Ptype ProviderType
    Zoneid string
    Owner string
    Domain string
    Ttl uint32
    Records []string
}
