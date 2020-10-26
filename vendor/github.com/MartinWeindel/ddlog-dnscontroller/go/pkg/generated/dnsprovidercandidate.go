package generated

// Generated code for typedef DNSProviderCandidate
// DO NOT CHANGE MANUALLY

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/vmware/differential-datalog/go/pkg/ddlog"
)


var (
	// memory will never be freed, which is fine
	relConstructorDNSProviderCandidate = ddlog.NewCString("DNSProviderCandidate")
)


func NewRecordDNSProviderCandidate(obj *DNSProviderCandidate) ddlog.Record {
	arg0 := func() ddlog.Record {
	    return NewRecordObjectKey(&obj.EntryKey)
    }()
	arg1 := func() ddlog.Record {
	    return ddlog.NewRecordString(obj.Domain)
    }()
	arg2 := func() ddlog.Record {
	    return NewRecordObjectKey(&obj.ProviderKey)
    }()
	arg3 := func() ddlog.Record {
	    return NewRecordPState(&obj.ProviderState)
    }()
	return ddlog.NewRecordStructStatic(relConstructorDNSProviderCandidate, arg0, arg1, arg2, arg3)
}


func DNSProviderCandidateFromRecord(record ddlog.Record) (*DNSProviderCandidate, error) {
	rs, err := record.AsStructSafe()
	if err != nil {
		return nil, errors.Wrapf(err, "no struct (DNSProviderCandidate)")
	}
	if rs.Name() != "DNSProviderCandidate" {
		return nil, fmt.Errorf("unexpected record %s != DNSProviderCandidate", rs.Name())
	}
	arg0, err := ObjectKeyFromRecord(rs.At(0))
	if err != nil {
		return nil, errors.Wrapf(err, "Field entryKey")
	}
	arg1, err := rs.At(1).ToStringSafe()
	if err != nil {
		return nil, errors.Wrapf(err, "Field domain")
	}
	arg2, err := ObjectKeyFromRecord(rs.At(2))
	if err != nil {
		return nil, errors.Wrapf(err, "Field providerKey")
	}
	arg3, err := PStateFromRecord(rs.At(3))
	if err != nil {
		return nil, errors.Wrapf(err, "Field providerState")
	}
	obj := &DNSProviderCandidate{	
		EntryKey:*arg0,	
		Domain:arg1,	
		ProviderKey:*arg2,	
		ProviderState:*arg3,
	}
	return obj, nil
}

type DNSProviderCandidate struct {
    EntryKey ObjectKey
    Domain string
    ProviderKey ObjectKey
    ProviderState PState
}
