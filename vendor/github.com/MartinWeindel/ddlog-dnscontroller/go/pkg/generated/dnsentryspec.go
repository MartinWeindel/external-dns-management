package generated

// Generated code for typedef DNSEntrySpec
// DO NOT CHANGE MANUALLY

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/vmware/differential-datalog/go/pkg/ddlog"
)


var (
	// memory will never be freed, which is fine
	relConstructorDNSEntrySpec = ddlog.NewCString("DNSEntrySpec")
)


func NewRecordDNSEntrySpec(obj *DNSEntrySpec) ddlog.Record {
	arg0 := func() ddlog.Record {
	    return NewRecordObjectKey(&obj.Key)
    }()
	arg1 := func() ddlog.Record {
	    return NewRecordEntrySpec(&obj.Spec)
    }()
	return ddlog.NewRecordStructStatic(relConstructorDNSEntrySpec, arg0, arg1)
}


func DNSEntrySpecFromRecord(record ddlog.Record) (*DNSEntrySpec, error) {
	rs, err := record.AsStructSafe()
	if err != nil {
		return nil, errors.Wrapf(err, "no struct (DNSEntrySpec)")
	}
	if rs.Name() != "DNSEntrySpec" {
		return nil, fmt.Errorf("unexpected record %s != DNSEntrySpec", rs.Name())
	}
	arg0, err := ObjectKeyFromRecord(rs.At(0))
	if err != nil {
		return nil, errors.Wrapf(err, "Field key")
	}
	arg1, err := EntrySpecFromRecord(rs.At(1))
	if err != nil {
		return nil, errors.Wrapf(err, "Field spec")
	}
	obj := &DNSEntrySpec{	
		Key:*arg0,	
		Spec:*arg1,
	}
	return obj, nil
}

type DNSEntrySpec struct {
    Key ObjectKey
    Spec EntrySpec
}
