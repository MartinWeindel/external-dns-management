package generated

// Generated code for typedef AccountResult
// DO NOT CHANGE MANUALLY

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/vmware/differential-datalog/go/pkg/ddlog"
)


var (
	// memory will never be freed, which is fine
	relConstructorAccountResult = ddlog.NewCString("AccountResult")
)


func NewRecordAccountResult(obj *AccountResult) ddlog.Record {
	arg0 := func() ddlog.Record {
	    return ddlog.NewRecordString(obj.CredentialsHash)
    }()
	arg1 := func() ddlog.Record {
	    return NewRecordProviderType(obj.ProviderType)
    }()
	arg2 := func() ddlog.Record {
	    return NewRecordStatus(obj.Status)
    }()
	arg3 := func() ddlog.Record {
	    vec := make([]ddlog.Record, len(obj.Zones))
	for i, item := range obj.Zones {
		vec[i] = NewRecordZone(&item)
	}
    return ddlog.NewRecordVector(vec...)
    }()
	return ddlog.NewRecordStructStatic(relConstructorAccountResult, arg0, arg1, arg2, arg3)
}


func AccountResultFromRecord(record ddlog.Record) (*AccountResult, error) {
	rs, err := record.AsStructSafe()
	if err != nil {
		return nil, errors.Wrapf(err, "no struct (AccountResult)")
	}
	if rs.Name() != "AccountResult" {
		return nil, fmt.Errorf("unexpected record %s != AccountResult", rs.Name())
	}
	arg0, err := rs.At(0).ToStringSafe()
	if err != nil {
		return nil, errors.Wrapf(err, "Field credentialsHash")
	}
	arg1, err := ProviderTypeFromRecord(rs.At(1))
	if err != nil {
		return nil, errors.Wrapf(err, "Field providerType")
	}
	arg2, err := StatusFromRecord(rs.At(2))
	if err != nil {
		return nil, errors.Wrapf(err, "Field status")
	}
	arg3, err := 
	func() ([]Zone, error) {
		rv0 := rs.At(3)
		rv, err := rv0.AsVectorSafe()
		if err != nil {
			return nil, err
		}
		vec := make([]Zone, rv.Size())
		for i := 0; i < len(vec); i++ {
			obj, err := ZoneFromRecord(rv.At(i))
			if err != nil {
				errors.Wrapf(err, "vector index %d", i)
			}
			vec[i] =*obj
		}
		return vec, nil
	}()
	if err != nil {
		return nil, errors.Wrapf(err, "Field zones")
	}
	obj := &AccountResult{	
		CredentialsHash:arg0,	
		ProviderType:arg1,	
		Status:arg2,	
		Zones:arg3,
	}
	return obj, nil
}

type AccountResult struct {
    CredentialsHash string
    ProviderType ProviderType
    Status Status
    Zones []Zone
}
