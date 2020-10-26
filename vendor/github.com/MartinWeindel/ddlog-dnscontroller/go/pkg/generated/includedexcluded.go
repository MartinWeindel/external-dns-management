package generated

// Generated code for typedef IncludedExcluded
// DO NOT CHANGE MANUALLY

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/vmware/differential-datalog/go/pkg/ddlog"
)


var (
	// memory will never be freed, which is fine
	relConstructorIncludedExcluded = ddlog.NewCString("IncludedExcluded")
)


func NewRecordIncludedExcluded(obj *IncludedExcluded) ddlog.Record {
	arg0 := func() ddlog.Record {
	    vec := make([]ddlog.Record, len(obj.Included))
	for i, item := range obj.Included {
		vec[i] = ddlog.NewRecordString(item)
	}
    return ddlog.NewRecordVector(vec...)
    }()
	arg1 := func() ddlog.Record {
	    vec := make([]ddlog.Record, len(obj.Excluded))
	for i, item := range obj.Excluded {
		vec[i] = ddlog.NewRecordString(item)
	}
    return ddlog.NewRecordVector(vec...)
    }()
	return ddlog.NewRecordStructStatic(relConstructorIncludedExcluded, arg0, arg1)
}


func IncludedExcludedFromRecord(record ddlog.Record) (*IncludedExcluded, error) {
	rs, err := record.AsStructSafe()
	if err != nil {
		return nil, errors.Wrapf(err, "no struct (IncludedExcluded)")
	}
	if rs.Name() != "IncludedExcluded" {
		return nil, fmt.Errorf("unexpected record %s != IncludedExcluded", rs.Name())
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
		return nil, errors.Wrapf(err, "Field included")
	}
	arg1, err := 
	func() ([]string, error) {
		rv, err := rs.At(1).AsVectorSafe()
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
		return nil, errors.Wrapf(err, "Field excluded")
	}
	obj := &IncludedExcluded{	
		Included:arg0,	
		Excluded:arg1,
	}
	return obj, nil
}

type IncludedExcluded struct {
    Included []string
    Excluded []string
}
