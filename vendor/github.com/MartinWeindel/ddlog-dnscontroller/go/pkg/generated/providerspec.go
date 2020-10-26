package generated

// Generated code for typedef ProviderSpec
// DO NOT CHANGE MANUALLY

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/vmware/differential-datalog/go/pkg/ddlog"
)


var (
	// memory will never be freed, which is fine
	relConstructorProviderSpec = ddlog.NewCString("ProviderSpec")
)


func NewRecordProviderSpec(obj *ProviderSpec) ddlog.Record {
	arg0 := func() ddlog.Record {
	    return ddlog.NewRecordString(obj.CredentialsHash)
    }()
	arg1 := func() ddlog.Record {
	    return NewRecordProviderType(obj.ProviderType)
    }()
	arg2 := func() ddlog.Record {
	    return NewRecordIncludeExclude(&obj.Domains)
    }()
	arg3 := func() ddlog.Record {
	    return NewRecordIncludeExclude(&obj.Zones)
    }()
	return ddlog.NewRecordStructStatic(relConstructorProviderSpec, arg0, arg1, arg2, arg3)
}


func ProviderSpecFromRecord(record ddlog.Record) (*ProviderSpec, error) {
	rs, err := record.AsStructSafe()
	if err != nil {
		return nil, errors.Wrapf(err, "no struct (ProviderSpec)")
	}
	if rs.Name() != "ProviderSpec" {
		return nil, fmt.Errorf("unexpected record %s != ProviderSpec", rs.Name())
	}
	arg0, err := rs.At(0).ToStringSafe()
	if err != nil {
		return nil, errors.Wrapf(err, "Field credentialsHash")
	}
	arg1, err := ProviderTypeFromRecord(rs.At(1))
	if err != nil {
		return nil, errors.Wrapf(err, "Field providerType")
	}
	arg2, err := IncludeExcludeFromRecord(rs.At(2))
	if err != nil {
		return nil, errors.Wrapf(err, "Field domains")
	}
	arg3, err := IncludeExcludeFromRecord(rs.At(3))
	if err != nil {
		return nil, errors.Wrapf(err, "Field zones")
	}
	obj := &ProviderSpec{	
		CredentialsHash:arg0,	
		ProviderType:arg1,	
		Domains:*arg2,	
		Zones:*arg3,
	}
	return obj, nil
}

type ProviderSpec struct {
    CredentialsHash string
    ProviderType ProviderType
    Domains IncludeExclude
    Zones IncludeExclude
}
