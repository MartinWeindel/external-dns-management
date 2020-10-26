package generated

// Generated code for typedef LightZone
// DO NOT CHANGE MANUALLY

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/vmware/differential-datalog/go/pkg/ddlog"
)


var (
	// memory will never be freed, which is fine
	relConstructorLightZone = ddlog.NewCString("LightZone")
)


func NewRecordLightZone(obj *LightZone) ddlog.Record {
	arg0 := func() ddlog.Record {
	    return ddlog.NewRecordString(obj.Id)
    }()
	arg1 := func() ddlog.Record {
	    return ddlog.NewRecordString(obj.Domain)
    }()
	return ddlog.NewRecordStructStatic(relConstructorLightZone, arg0, arg1)
}


func LightZoneFromRecord(record ddlog.Record) (*LightZone, error) {
	rs, err := record.AsStructSafe()
	if err != nil {
		return nil, errors.Wrapf(err, "no struct (LightZone)")
	}
	if rs.Name() != "LightZone" {
		return nil, fmt.Errorf("unexpected record %s != LightZone", rs.Name())
	}
	arg0, err := rs.At(0).ToStringSafe()
	if err != nil {
		return nil, errors.Wrapf(err, "Field id")
	}
	arg1, err := rs.At(1).ToStringSafe()
	if err != nil {
		return nil, errors.Wrapf(err, "Field domain")
	}
	obj := &LightZone{	
		Id:arg0,	
		Domain:arg1,
	}
	return obj, nil
}

type LightZone struct {
    Id string
    Domain string
}
