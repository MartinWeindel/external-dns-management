package generated

// Generated code for typedef EntrySpec
// DO NOT CHANGE MANUALLY

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/vmware/differential-datalog/go/pkg/ddlog"
)


var (
	// memory will never be freed, which is fine
	relConstructorEntrySpec = ddlog.NewCString("EntrySpec")
)


func NewRecordEntrySpec(obj *EntrySpec) ddlog.Record {
	arg0 := func() ddlog.Record {
	    return ddlog.NewRecordString(obj.Owner)
    }()
	arg1 := func() ddlog.Record {
	    return ddlog.NewRecordString(obj.Domain)
    }()
	arg2 := func() ddlog.Record {
	    return ddlog.NewRecordU32(obj.Ttl)
    }()
	arg3 := func() ddlog.Record {
	    return NewRecordRecordType(obj.Rtype)
    }()
	arg4 := func() ddlog.Record {
	    vec := make([]ddlog.Record, len(obj.Records))
	for i, item := range obj.Records {
		vec[i] = ddlog.NewRecordString(item)
	}
    return ddlog.NewRecordVector(vec...)
    }()
	return ddlog.NewRecordStructStatic(relConstructorEntrySpec, arg0, arg1, arg2, arg3, arg4)
}


func EntrySpecFromRecord(record ddlog.Record) (*EntrySpec, error) {
	rs, err := record.AsStructSafe()
	if err != nil {
		return nil, errors.Wrapf(err, "no struct (EntrySpec)")
	}
	if rs.Name() != "EntrySpec" {
		return nil, fmt.Errorf("unexpected record %s != EntrySpec", rs.Name())
	}
	arg0, err := rs.At(0).ToStringSafe()
	if err != nil {
		return nil, errors.Wrapf(err, "Field owner")
	}
	arg1, err := rs.At(1).ToStringSafe()
	if err != nil {
		return nil, errors.Wrapf(err, "Field domain")
	}
	arg2, err := rs.At(2).ToU32Safe()
	if err != nil {
		return nil, errors.Wrapf(err, "Field ttl")
	}
	arg3, err := RecordTypeFromRecord(rs.At(3))
	if err != nil {
		return nil, errors.Wrapf(err, "Field rtype")
	}
	arg4, err := 
	func() ([]string, error) {
		rv0 := rs.At(4)
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
	obj := &EntrySpec{	
		Owner:arg0,	
		Domain:arg1,	
		Ttl:arg2,	
		Rtype:arg3,	
		Records:arg4,
	}
	return obj, nil
}

type EntrySpec struct {
    Owner string
    Domain string
    Ttl uint32
    Rtype RecordType
    Records []string
}
