package generated

// Generated code for typedef MatchedEntryToZone
// DO NOT CHANGE MANUALLY

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/vmware/differential-datalog/go/pkg/ddlog"
)


var (
	// memory will never be freed, which is fine
	relConstructorMatchedEntryToZone = ddlog.NewCString("MatchedEntryToZone")
)


func NewRecordMatchedEntryToZone(obj *MatchedEntryToZone) ddlog.Record {
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
	arg4 := func() ddlog.Record {
	    return ddlog.NewRecordString(obj.Zoneid)
    }()
	return ddlog.NewRecordStructStatic(relConstructorMatchedEntryToZone, arg0, arg1, arg2, arg3, arg4)
}


func MatchedEntryToZoneFromRecord(record ddlog.Record) (*MatchedEntryToZone, error) {
	rs, err := record.AsStructSafe()
	if err != nil {
		return nil, errors.Wrapf(err, "no struct (MatchedEntryToZone)")
	}
	if rs.Name() != "MatchedEntryToZone" {
		return nil, fmt.Errorf("unexpected record %s != MatchedEntryToZone", rs.Name())
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
	arg4, err := rs.At(4).ToStringSafe()
	if err != nil {
		return nil, errors.Wrapf(err, "Field zoneid")
	}
	obj := &MatchedEntryToZone{	
		EntryKey:*arg0,	
		Domain:arg1,	
		ProviderKey:*arg2,	
		ProviderState:*arg3,	
		Zoneid:arg4,
	}
	return obj, nil
}

type MatchedEntryToZone struct {
    EntryKey ObjectKey
    Domain string
    ProviderKey ObjectKey
    ProviderState PState
    Zoneid string
}
