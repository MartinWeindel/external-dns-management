package generated

// Generated code for typedef DNSProviderZone
// DO NOT CHANGE MANUALLY

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/vmware/differential-datalog/go/pkg/ddlog"
)


var (
	// memory will never be freed, which is fine
	relConstructorDNSProviderZone = ddlog.NewCString("DNSProviderZone")
)


func NewRecordDNSProviderZone(obj *DNSProviderZone) ddlog.Record {
	arg0 := func() ddlog.Record {
	    return NewRecordObjectKey(&obj.Key)
    }()
	arg1 := func() ddlog.Record {
	    return NewRecordProviderType(obj.Ptype)
    }()
	arg2 := func() ddlog.Record {
	    return ddlog.NewRecordString(obj.Zoneid)
    }()
	return ddlog.NewRecordStructStatic(relConstructorDNSProviderZone, arg0, arg1, arg2)
}


func DNSProviderZoneFromRecord(record ddlog.Record) (*DNSProviderZone, error) {
	rs, err := record.AsStructSafe()
	if err != nil {
		return nil, errors.Wrapf(err, "no struct (DNSProviderZone)")
	}
	if rs.Name() != "DNSProviderZone" {
		return nil, fmt.Errorf("unexpected record %s != DNSProviderZone", rs.Name())
	}
	arg0, err := ObjectKeyFromRecord(rs.At(0))
	if err != nil {
		return nil, errors.Wrapf(err, "Field key")
	}
	arg1, err := ProviderTypeFromRecord(rs.At(1))
	if err != nil {
		return nil, errors.Wrapf(err, "Field ptype")
	}
	arg2, err := rs.At(2).ToStringSafe()
	if err != nil {
		return nil, errors.Wrapf(err, "Field zoneid")
	}
	obj := &DNSProviderZone{	
		Key:*arg0,	
		Ptype:arg1,	
		Zoneid:arg2,
	}
	return obj, nil
}

type DNSProviderZone struct {
    Key ObjectKey
    Ptype ProviderType
    Zoneid string
}
