package generated

// Generated code for typedef DNSOwner
// DO NOT CHANGE MANUALLY

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/vmware/differential-datalog/go/pkg/ddlog"
)


var (
	// memory will never be freed, which is fine
	relConstructorDNSOwner = ddlog.NewCString("DNSOwner")
)


func NewRecordDNSOwner(obj *DNSOwner) ddlog.Record {
	arg0 := func() ddlog.Record {
	    return ddlog.NewRecordString(obj.Name)
    }()
	return ddlog.NewRecordStructStatic(relConstructorDNSOwner, arg0)
}


func DNSOwnerFromRecord(record ddlog.Record) (*DNSOwner, error) {
	rs, err := record.AsStructSafe()
	if err != nil {
		return nil, errors.Wrapf(err, "no struct (DNSOwner)")
	}
	if rs.Name() != "DNSOwner" {
		return nil, fmt.Errorf("unexpected record %s != DNSOwner", rs.Name())
	}
	arg0, err := rs.At(0).ToStringSafe()
	if err != nil {
		return nil, errors.Wrapf(err, "Field name")
	}
	obj := &DNSOwner{	
		Name:arg0,
	}
	return obj, nil
}

type DNSOwner struct {
    Name string
}
