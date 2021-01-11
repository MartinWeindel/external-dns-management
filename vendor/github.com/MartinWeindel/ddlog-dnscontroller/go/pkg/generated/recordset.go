package generated

// Generated code for typedef RecordSet
// DO NOT CHANGE MANUALLY

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/vmware/differential-datalog/go/pkg/ddlog"
)


var (
	// memory will never be freed, which is fine
	relConstructorRecordSet = ddlog.NewCString("RecordSet")
)


func NewRecordRecordSet(obj *RecordSet) ddlog.Record {
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
	arg5 := func() ddlog.Record {
	    return ddlog.NewRecordU32(obj.Ttl)
    }()
	arg6 := func() ddlog.Record {
	    vec := make([]ddlog.Record, len(obj.Records))
	for i, item := range obj.Records {
		vec[i] = ddlog.NewRecordString(item)
	}
    return ddlog.NewRecordVector(vec...)
    }()
	return ddlog.NewRecordStructStatic(relConstructorRecordSet, arg0, arg1, arg2, arg3, arg4, arg5, arg6)
}


func RecordSetFromRecord(record ddlog.Record) (*RecordSet, error) {
	rs, err := record.AsStructSafe()
	if err != nil {
		return nil, errors.Wrapf(err, "no struct (RecordSet)")
	}
	if rs.Name() != "RecordSet" {
		return nil, fmt.Errorf("unexpected record %s != RecordSet", rs.Name())
	}
	arg0, err := rs.At(0).ToStringSafe()
	if err != nil {
		return nil, errors.Wrapf(err, "Field owner")
	}
	arg1, err := ProviderTypeFromRecord(rs.At(1))
	if err != nil {
		return nil, errors.Wrapf(err, "Field ptype")
	}
	arg2, err := rs.At(2).ToStringSafe()
	if err != nil {
		return nil, errors.Wrapf(err, "Field zoneid")
	}
	arg3, err := rs.At(3).ToStringSafe()
	if err != nil {
		return nil, errors.Wrapf(err, "Field domain")
	}
	arg4, err := RecordTypeFromRecord(rs.At(4))
	if err != nil {
		return nil, errors.Wrapf(err, "Field rtype")
	}
	arg5, err := rs.At(5).ToU32Safe()
	if err != nil {
		return nil, errors.Wrapf(err, "Field ttl")
	}
	arg6, err := 
	func() ([]string, error) {
		rv0 := rs.At(6)
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
	obj := &RecordSet{	
		Owner:arg0,	
		Ptype:arg1,	
		Zoneid:arg2,	
		Domain:arg3,	
		Rtype:arg4,	
		Ttl:arg5,	
		Records:arg6,
	}
	return obj, nil
}

type RecordSet struct {
    Owner string
    Ptype ProviderType
    Zoneid string
    Domain string
    Rtype RecordType
    Ttl uint32
    Records []string
}
