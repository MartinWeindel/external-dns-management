package generated

// Generated code for typedef AllDNSOwners
// DO NOT CHANGE MANUALLY

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/vmware/differential-datalog/go/pkg/ddlog"
)


var (
	// memory will never be freed, which is fine
	relConstructorAllDNSOwners = ddlog.NewCString("AllDNSOwners")
)


func NewRecordAllDNSOwners(obj *AllDNSOwners) ddlog.Record {
	arg0 := func() ddlog.Record {
	    vec := make([]ddlog.Record, len(obj.OwnerIds))
	for i, item := range obj.OwnerIds {
		vec[i] = ddlog.NewRecordString(item)
	}
    return ddlog.NewRecordVector(vec...)
    }()
	return ddlog.NewRecordStructStatic(relConstructorAllDNSOwners, arg0)
}


func AllDNSOwnersFromRecord(record ddlog.Record) (*AllDNSOwners, error) {
	rs, err := record.AsStructSafe()
	if err != nil {
		return nil, errors.Wrapf(err, "no struct (AllDNSOwners)")
	}
	if rs.Name() != "AllDNSOwners" {
		return nil, fmt.Errorf("unexpected record %s != AllDNSOwners", rs.Name())
	}
	arg0, err := 
	func() ([]string, error) {
		rv, err := rs.At(0).AsVectorSafe()
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
		return nil, errors.Wrapf(err, "Field ownerIds")
	}
	obj := &AllDNSOwners{	
		OwnerIds:arg0,
	}
	return obj, nil
}

type AllDNSOwners struct {
    OwnerIds []string
}
