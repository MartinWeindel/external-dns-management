package generated

// Generated code for typedef DNSProviderStatus
// DO NOT CHANGE MANUALLY

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/vmware/differential-datalog/go/pkg/ddlog"
)


var (
	// memory will never be freed, which is fine
	relConstructorDNSProviderStatus = ddlog.NewCString("DNSProviderStatus")
)


func NewRecordDNSProviderStatus(obj *DNSProviderStatus) ddlog.Record {
	arg0 := func() ddlog.Record {
	    return NewRecordObjectKey(&obj.Key)
    }()
	arg1 := func() ddlog.Record {
	    return NewRecordStatus(obj.Status)
    }()
	arg2 := func() ddlog.Record {
	    
	var opt ddlog.Record
	if obj.State == nil {
		opt = ddlog.NewRecordNone()
	} else {
		rec := NewRecordPState(&(*obj.State))
		opt = ddlog.NewRecordSome(rec)
	}
	return opt
    }()
	return ddlog.NewRecordStructStatic(relConstructorDNSProviderStatus, arg0, arg1, arg2)
}


func DNSProviderStatusFromRecord(record ddlog.Record) (*DNSProviderStatus, error) {
	rs, err := record.AsStructSafe()
	if err != nil {
		return nil, errors.Wrapf(err, "no struct (DNSProviderStatus)")
	}
	if rs.Name() != "DNSProviderStatus" {
		return nil, fmt.Errorf("unexpected record %s != DNSProviderStatus", rs.Name())
	}
	arg0, err := ObjectKeyFromRecord(rs.At(0))
	if err != nil {
		return nil, errors.Wrapf(err, "Field key")
	}
	arg1, err := StatusFromRecord(rs.At(1))
	if err != nil {
		return nil, errors.Wrapf(err, "Field status")
	}
	arg2, err := func() (*PState, error) {
		rs, err := rs.At(2).AsStructSafe()
		if err != nil {
			return nil, err
		}
		switch rs.Name() {
		case "ddlog_std::None":
			return nil, nil
		case "ddlog_std::Some":
			value, err := PStateFromRecord(rs.At(0))
			return value, err
		default:
			return nil, fmt.Errorf("expected option record: %s", rs.Name())
		}
	}()
	if err != nil {
		return nil, errors.Wrapf(err, "Field state")
	}
	obj := &DNSProviderStatus{	
		Key:*arg0,	
		Status:arg1,	
		State:arg2,
	}
	return obj, nil
}

type DNSProviderStatus struct {
    Key ObjectKey
    Status Status
    State *PState
}
