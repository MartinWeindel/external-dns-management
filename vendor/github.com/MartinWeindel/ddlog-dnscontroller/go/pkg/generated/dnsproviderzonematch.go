package generated

// Generated code for typedef DNSProviderZoneMatch
// DO NOT CHANGE MANUALLY

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/vmware/differential-datalog/go/pkg/ddlog"
)


var (
	// memory will never be freed, which is fine
	relConstructorDNSProviderZoneMatch = ddlog.NewCString("DNSProviderZoneMatch")
)


func NewRecordDNSProviderZoneMatch(obj *DNSProviderZoneMatch) ddlog.Record {
	arg0 := func() ddlog.Record {
	    return NewRecordObjectKey(&obj.Key)
    }()
	arg1 := func() ddlog.Record {
	    return NewRecordZoneMatch(&obj.ZoneMatch)
    }()
	return ddlog.NewRecordStructStatic(relConstructorDNSProviderZoneMatch, arg0, arg1)
}


func DNSProviderZoneMatchFromRecord(record ddlog.Record) (*DNSProviderZoneMatch, error) {
	rs, err := record.AsStructSafe()
	if err != nil {
		return nil, errors.Wrapf(err, "no struct (DNSProviderZoneMatch)")
	}
	if rs.Name() != "DNSProviderZoneMatch" {
		return nil, fmt.Errorf("unexpected record %s != DNSProviderZoneMatch", rs.Name())
	}
	arg0, err := ObjectKeyFromRecord(rs.At(0))
	if err != nil {
		return nil, errors.Wrapf(err, "Field key")
	}
	arg1, err := ZoneMatchFromRecord(rs.At(1))
	if err != nil {
		return nil, errors.Wrapf(err, "Field zoneMatch")
	}
	obj := &DNSProviderZoneMatch{	
		Key:*arg0,	
		ZoneMatch:*arg1,
	}
	return obj, nil
}

type DNSProviderZoneMatch struct {
    Key ObjectKey
    ZoneMatch ZoneMatch
}
