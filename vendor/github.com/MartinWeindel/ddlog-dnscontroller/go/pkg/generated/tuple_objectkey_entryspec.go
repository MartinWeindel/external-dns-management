package generated

// Generated code for typedef Tuple_ObjectKey_EntrySpec
// DO NOT CHANGE MANUALLY

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/vmware/differential-datalog/go/pkg/ddlog"
)


var (
	// memory will never be freed, which is fine
	relConstructorTuple_ObjectKey_EntrySpec = ddlog.NewCString("Tuple_ObjectKey_EntrySpec")
)


func NewRecordTuple_ObjectKey_EntrySpec(obj *Tuple_ObjectKey_EntrySpec) ddlog.Record {
	arg0 := func() ddlog.Record {
	    return NewRecordObjectKey(&obj.Arg0)
    }()
	arg1 := func() ddlog.Record {
	    return NewRecordEntrySpec(&obj.Arg1)
    }()
	return ddlog.NewRecordStructStatic(relConstructorTuple_ObjectKey_EntrySpec, arg0, arg1)
}


func Tuple_ObjectKey_EntrySpecFromRecord(record ddlog.Record) (*Tuple_ObjectKey_EntrySpec, error) {
	rs, err := record.AsStructSafe()
	if err != nil {
		return nil, errors.Wrapf(err, "no struct (Tuple_ObjectKey_EntrySpec)")
	}
	if rs.Name() != "Tuple_ObjectKey_EntrySpec" {
		return nil, fmt.Errorf("unexpected record %s != Tuple_ObjectKey_EntrySpec", rs.Name())
	}
	arg0, err := ObjectKeyFromRecord(rs.At(0))
	if err != nil {
		return nil, errors.Wrapf(err, "Field Arg0")
	}
	arg1, err := EntrySpecFromRecord(rs.At(1))
	if err != nil {
		return nil, errors.Wrapf(err, "Field Arg1")
	}
	obj := &Tuple_ObjectKey_EntrySpec{	
		Arg0:*arg0,	
		Arg1:*arg1,
	}
	return obj, nil
}

type Tuple_ObjectKey_EntrySpec struct {
    Arg0 ObjectKey
    Arg1 EntrySpec
}
