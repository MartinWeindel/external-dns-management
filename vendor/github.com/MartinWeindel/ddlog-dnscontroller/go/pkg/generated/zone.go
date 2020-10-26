package generated

// Generated code for typedef Zone
// DO NOT CHANGE MANUALLY

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/vmware/differential-datalog/go/pkg/ddlog"
)


var (
	// memory will never be freed, which is fine
	relConstructorZone = ddlog.NewCString("Zone")
)


func NewRecordZone(obj *Zone) ddlog.Record {
	arg0 := func() ddlog.Record {
	    return ddlog.NewRecordString(obj.Id)
    }()
	arg1 := func() ddlog.Record {
	    return ddlog.NewRecordString(obj.Domain)
    }()
	arg2 := func() ddlog.Record {
	    vec := make([]ddlog.Record, len(obj.ForwardedDomains))
	for i, item := range obj.ForwardedDomains {
		vec[i] = ddlog.NewRecordString(item)
	}
    return ddlog.NewRecordVector(vec...)
    }()
	return ddlog.NewRecordStructStatic(relConstructorZone, arg0, arg1, arg2)
}


func ZoneFromRecord(record ddlog.Record) (*Zone, error) {
	rs, err := record.AsStructSafe()
	if err != nil {
		return nil, errors.Wrapf(err, "no struct (Zone)")
	}
	if rs.Name() != "Zone" {
		return nil, fmt.Errorf("unexpected record %s != Zone", rs.Name())
	}
	arg0, err := rs.At(0).ToStringSafe()
	if err != nil {
		return nil, errors.Wrapf(err, "Field id")
	}
	arg1, err := rs.At(1).ToStringSafe()
	if err != nil {
		return nil, errors.Wrapf(err, "Field domain")
	}
	arg2, err := 
	func() ([]string, error) {
		rv, err := rs.At(2).AsVectorSafe()
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
		return nil, errors.Wrapf(err, "Field forwardedDomains")
	}
	obj := &Zone{	
		Id:arg0,	
		Domain:arg1,	
		ForwardedDomains:arg2,
	}
	return obj, nil
}

type Zone struct {
    Id string
    Domain string
    ForwardedDomains []string
}
