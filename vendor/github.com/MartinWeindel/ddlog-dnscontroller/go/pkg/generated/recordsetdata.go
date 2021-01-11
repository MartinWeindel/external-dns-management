package generated

// Generated code for typedef RecordSetData
// DO NOT CHANGE MANUALLY

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/vmware/differential-datalog/go/pkg/ddlog"
)


var (
	// memory will never be freed, which is fine
	relConstructorRecordSetData = ddlog.NewCString("RecordSetData")
)


func NewRecordRecordSetData(obj *RecordSetData) ddlog.Record {
	arg0 := func() ddlog.Record {
	    return ddlog.NewRecordString(obj.Owner)
    }()
	arg1 := func() ddlog.Record {
	    return ddlog.NewRecordString(obj.Domain)
    }()
	arg2 := func() ddlog.Record {
	    return NewRecordRecordType(obj.Rtype)
    }()
	arg3 := func() ddlog.Record {
	    return ddlog.NewRecordU32(obj.Ttl)
    }()
	arg4 := func() ddlog.Record {
	    vec := make([]ddlog.Record, len(obj.Records))
	for i, item := range obj.Records {
		vec[i] = ddlog.NewRecordString(item)
	}
    return ddlog.NewRecordVector(vec...)
    }()
	return ddlog.NewRecordStructStatic(relConstructorRecordSetData, arg0, arg1, arg2, arg3, arg4)
}


func RecordSetDataFromRecord(record ddlog.Record) (*RecordSetData, error) {
	rs, err := record.AsStructSafe()
	if err != nil {
		return nil, errors.Wrapf(err, "no struct (RecordSetData)")
	}
	if rs.Name() != "RecordSetData" {
		return nil, fmt.Errorf("unexpected record %s != RecordSetData", rs.Name())
	}
	arg0, err := rs.At(0).ToStringSafe()
	if err != nil {
		return nil, errors.Wrapf(err, "Field owner")
	}
	arg1, err := rs.At(1).ToStringSafe()
	if err != nil {
		return nil, errors.Wrapf(err, "Field domain")
	}
	arg2, err := RecordTypeFromRecord(rs.At(2))
	if err != nil {
		return nil, errors.Wrapf(err, "Field rtype")
	}
	arg3, err := rs.At(3).ToU32Safe()
	if err != nil {
		return nil, errors.Wrapf(err, "Field ttl")
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
	obj := &RecordSetData{	
		Owner:arg0,	
		Domain:arg1,	
		Rtype:arg2,	
		Ttl:arg3,	
		Records:arg4,
	}
	return obj, nil
}

type RecordSetData struct {
    Owner string
    Domain string
    Rtype RecordType
    Ttl uint32
    Records []string
}
