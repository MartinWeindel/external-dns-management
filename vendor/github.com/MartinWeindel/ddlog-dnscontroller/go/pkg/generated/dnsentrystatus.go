package generated

// Generated code for typedef DNSEntryStatus
// DO NOT CHANGE MANUALLY

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/vmware/differential-datalog/go/pkg/ddlog"
)


var (
	// memory will never be freed, which is fine
	relConstructorDNSEntryStatus = ddlog.NewCString("DNSEntryStatus")
)


func NewRecordDNSEntryStatus(obj *DNSEntryStatus) ddlog.Record {
	arg0 := func() ddlog.Record {
	    return NewRecordObjectKey(&obj.Key)
    }()
	arg1 := func() ddlog.Record {
	    return NewRecordEntryStatus(obj.Status)
    }()
	arg2 := func() ddlog.Record {
	    
	var opt ddlog.Record
	if obj.State == nil {
		opt = ddlog.NewRecordNone()
	} else {
		rec := NewRecordEState(&(*obj.State))
		opt = ddlog.NewRecordSome(rec)
	}
	return opt
    }()
	return ddlog.NewRecordStructStatic(relConstructorDNSEntryStatus, arg0, arg1, arg2)
}


func DNSEntryStatusFromRecord(record ddlog.Record) (*DNSEntryStatus, error) {
	rs, err := record.AsStructSafe()
	if err != nil {
		return nil, errors.Wrapf(err, "no struct (DNSEntryStatus)")
	}
	if rs.Name() != "DNSEntryStatus" {
		return nil, fmt.Errorf("unexpected record %s != DNSEntryStatus", rs.Name())
	}
	arg0, err := ObjectKeyFromRecord(rs.At(0))
	if err != nil {
		return nil, errors.Wrapf(err, "Field key")
	}
	arg1, err := EntryStatusFromRecord(rs.At(1))
	if err != nil {
		return nil, errors.Wrapf(err, "Field status")
	}
	arg2, err := func() (*EState, error) {
		rs, err := rs.At(2).AsStructSafe()
		if err != nil {
			return nil, err
		}
		switch rs.Name() {
		case "ddlog_std::None":
			return nil, nil
		case "ddlog_std::Some":
			value, err := EStateFromRecord(rs.At(0))
			return value, err
		default:
			return nil, fmt.Errorf("expected option record: %s", rs.Name())
		}
	}()
	if err != nil {
		return nil, errors.Wrapf(err, "Field state")
	}
	obj := &DNSEntryStatus{	
		Key:*arg0,	
		Status:arg1,	
		State:arg2,
	}
	return obj, nil
}

type DNSEntryStatus struct {
    Key ObjectKey
    Status EntryStatus
    State *EState
}
