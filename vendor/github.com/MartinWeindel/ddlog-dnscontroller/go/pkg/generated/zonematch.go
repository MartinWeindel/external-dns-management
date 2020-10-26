package generated

// Generated code for typedef ZoneMatch
// DO NOT CHANGE MANUALLY

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/vmware/differential-datalog/go/pkg/ddlog"
)


var (
	// memory will never be freed, which is fine
	relConstructorZoneMatch = ddlog.NewCString("ZoneMatch")
)


func NewRecordZoneMatch(obj *ZoneMatch) ddlog.Record {
	arg0 := func() ddlog.Record {
	    return NewRecordStatus(obj.AccountStatus)
    }()
	arg1 := func() ddlog.Record {
	    return NewRecordZone(&obj.Zone)
    }()
	arg2 := func() ddlog.Record {
	    
	var opt ddlog.Record
	if obj.DomainMatch == nil {
		opt = ddlog.NewRecordNone()
	} else {
		rec := ddlog.NewRecordString((*obj.DomainMatch))
		opt = ddlog.NewRecordSome(rec)
	}
	return opt
    }()
	arg3 := func() ddlog.Record {
	    
	var opt ddlog.Record
	if obj.SuperZone == nil {
		opt = ddlog.NewRecordNone()
	} else {
		rec := NewRecordZone(&(*obj.SuperZone))
		opt = ddlog.NewRecordSome(rec)
	}
	return opt
    }()
	return ddlog.NewRecordStructStatic(relConstructorZoneMatch, arg0, arg1, arg2, arg3)
}


func ZoneMatchFromRecord(record ddlog.Record) (*ZoneMatch, error) {
	rs, err := record.AsStructSafe()
	if err != nil {
		return nil, errors.Wrapf(err, "no struct (ZoneMatch)")
	}
	if rs.Name() != "ZoneMatch" {
		return nil, fmt.Errorf("unexpected record %s != ZoneMatch", rs.Name())
	}
	arg0, err := StatusFromRecord(rs.At(0))
	if err != nil {
		return nil, errors.Wrapf(err, "Field accountStatus")
	}
	arg1, err := ZoneFromRecord(rs.At(1))
	if err != nil {
		return nil, errors.Wrapf(err, "Field zone")
	}
	arg2, err := func() (*string, error) {
		rs, err := rs.At(2).AsStructSafe()
		if err != nil {
			return nil, err
		}
		switch rs.Name() {
		case "ddlog_std::None":
			return nil, nil
		case "ddlog_std::Some":
			value, err := rs.At(0).ToStringSafe()
			return &value, err
		default:
			return nil, fmt.Errorf("expected option record: %s", rs.Name())
		}
	}()
	if err != nil {
		return nil, errors.Wrapf(err, "Field domainMatch")
	}
	arg3, err := func() (*Zone, error) {
		rs, err := rs.At(3).AsStructSafe()
		if err != nil {
			return nil, err
		}
		switch rs.Name() {
		case "ddlog_std::None":
			return nil, nil
		case "ddlog_std::Some":
			value, err := ZoneFromRecord(rs.At(0))
			return value, err
		default:
			return nil, fmt.Errorf("expected option record: %s", rs.Name())
		}
	}()
	if err != nil {
		return nil, errors.Wrapf(err, "Field superZone")
	}
	obj := &ZoneMatch{	
		AccountStatus:arg0,	
		Zone:*arg1,	
		DomainMatch:arg2,	
		SuperZone:arg3,
	}
	return obj, nil
}

type ZoneMatch struct {
    AccountStatus Status
    Zone Zone
    DomainMatch *string
    SuperZone *Zone
}
