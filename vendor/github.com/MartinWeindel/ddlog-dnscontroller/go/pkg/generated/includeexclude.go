package generated

// Generated code for typedef IncludeExclude
// DO NOT CHANGE MANUALLY

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/vmware/differential-datalog/go/pkg/ddlog"
)


var (
	// memory will never be freed, which is fine
	relConstructorIncludeExclude = ddlog.NewCString("IncludeExclude")
)


func NewRecordIncludeExclude(obj *IncludeExclude) ddlog.Record {
	arg0 := func() ddlog.Record {
	    
	var opt ddlog.Record
	if obj.Include == nil {
		opt = ddlog.NewRecordNone()
	} else {
		rec := func() ddlog.Record {
	    vec := make([]ddlog.Record, len(obj.Include))
	for i, item := range obj.Include {
		vec[i] = ddlog.NewRecordString(item)
	}
    return ddlog.NewRecordVector(vec...)
    }()
		opt = ddlog.NewRecordSome(rec)
	}
	return opt
    }()
	arg1 := func() ddlog.Record {
	    
	var opt ddlog.Record
	if obj.Exclude == nil {
		opt = ddlog.NewRecordNone()
	} else {
		rec := func() ddlog.Record {
	    vec := make([]ddlog.Record, len(obj.Exclude))
	for i, item := range obj.Exclude {
		vec[i] = ddlog.NewRecordString(item)
	}
    return ddlog.NewRecordVector(vec...)
    }()
		opt = ddlog.NewRecordSome(rec)
	}
	return opt
    }()
	return ddlog.NewRecordStructStatic(relConstructorIncludeExclude, arg0, arg1)
}


func IncludeExcludeFromRecord(record ddlog.Record) (*IncludeExclude, error) {
	rs, err := record.AsStructSafe()
	if err != nil {
		return nil, errors.Wrapf(err, "no struct (IncludeExclude)")
	}
	if rs.Name() != "IncludeExclude" {
		return nil, fmt.Errorf("unexpected record %s != IncludeExclude", rs.Name())
	}
	arg0, err := func() ([]string, error) {
		rs, err := rs.At(0).AsStructSafe()
		if err != nil {
			return nil, err
		}
		switch rs.Name() {
		case "ddlog_std::None":
			return nil, nil
		case "ddlog_std::Some":
			value, err := 
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
			return value, err
		default:
			return nil, fmt.Errorf("expected option record: %s", rs.Name())
		}
	}()
	if err != nil {
		return nil, errors.Wrapf(err, "Field include")
	}
	arg1, err := func() ([]string, error) {
		rs, err := rs.At(1).AsStructSafe()
		if err != nil {
			return nil, err
		}
		switch rs.Name() {
		case "ddlog_std::None":
			return nil, nil
		case "ddlog_std::Some":
			value, err := 
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
			return value, err
		default:
			return nil, fmt.Errorf("expected option record: %s", rs.Name())
		}
	}()
	if err != nil {
		return nil, errors.Wrapf(err, "Field exclude")
	}
	obj := &IncludeExclude{	
		Include:arg0,	
		Exclude:arg1,
	}
	return obj, nil
}

type IncludeExclude struct {
    Include []string
    Exclude []string
}
