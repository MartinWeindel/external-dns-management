package generated

// Generated code for typedef DNSProviderSpec
// DO NOT CHANGE MANUALLY

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/vmware/differential-datalog/go/pkg/ddlog"
)

var (
	// memory will never be freed, which is fine
	relConstructorDNSProviderSpec = ddlog.NewCString("DNSProviderSpec")
)

func NewRecordDNSProviderSpec(obj *DNSProviderSpec) ddlog.Record {
	arg0 := func() ddlog.Record {
		return NewRecordObjectKey(&obj.Key)
	}()
	arg1 := func() ddlog.Record {
		return NewRecordProviderSpec(&obj.Spec)
	}()
	return ddlog.NewRecordStructStatic(relConstructorDNSProviderSpec, arg0, arg1)
}

func NewRecordKey_DNSProviderSpec(obj *DNSProviderSpec) ddlog.Record {
	arg0 := func() ddlog.Record {
		return NewRecordObjectKey(&obj.Key)
	}()
	return arg0
}

func DNSProviderSpecFromRecord(record ddlog.Record) (*DNSProviderSpec, error) {
	rs, err := record.AsStructSafe()
	if err != nil {
		return nil, errors.Wrapf(err, "no struct (DNSProviderSpec)")
	}
	if rs.Name() != "DNSProviderSpec" {
		return nil, fmt.Errorf("unexpected record %s != DNSProviderSpec", rs.Name())
	}
	arg0, err := ObjectKeyFromRecord(rs.At(0))
	if err != nil {
		return nil, errors.Wrapf(err, "Field key")
	}
	arg1, err := ProviderSpecFromRecord(rs.At(1))
	if err != nil {
		return nil, errors.Wrapf(err, "Field spec")
	}
	obj := &DNSProviderSpec{
		Key:  *arg0,
		Spec: *arg1,
	}
	return obj, nil
}

type DNSProviderSpec struct {
	Key  ObjectKey
	Spec ProviderSpec
}
