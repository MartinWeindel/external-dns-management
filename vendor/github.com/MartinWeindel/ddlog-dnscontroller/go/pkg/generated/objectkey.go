package generated

// Generated code for typedef ObjectKey
// DO NOT CHANGE MANUALLY

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/vmware/differential-datalog/go/pkg/ddlog"
)

type ObjectKey struct {
	Arg0 string
	Arg1 string
}

func NewRecordObjectKey(obj *ObjectKey) ddlog.Record {
	arg0 := func() ddlog.Record {
	    return ddlog.NewRecordString(obj.Arg0)
    }()
	arg1 := func() ddlog.Record {
	    return ddlog.NewRecordString(obj.Arg1)
    }()
	return ddlog.NewRecordTuple(arg0,arg1)
}

func ObjectKeyFromRecord(record ddlog.Record) (*ObjectKey, error) {
	rt, err := record.AsTupleSafe()
	if err != nil {
		return nil, fmt.Errorf("%s", err)
	}
	obj := &ObjectKey{}
	obj.Arg0, err = rt.At(0).ToStringSafe()
	if err != nil {
		return nil, errors.Wrapf(err, "Tuple index 0")
	}
	obj.Arg1, err = rt.At(1).ToStringSafe()
	if err != nil {
		return nil, errors.Wrapf(err, "Tuple index 1")
	}
	return obj, nil
}
