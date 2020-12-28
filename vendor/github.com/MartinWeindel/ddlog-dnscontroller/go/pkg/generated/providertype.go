package generated

// Generated code for typedef ProviderType
// DO NOT CHANGE MANUALLY

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/vmware/differential-datalog/go/pkg/ddlog"
)

type ProviderType interface {
	internalProviderType()
	Name() string
	NewRecord() ddlog.Record
}

func NewRecordProviderType(obj ProviderType) ddlog.Record {
	if obj == nil {
		return ddlog.NewRecordNull() 
	}
	return obj.NewRecord()
}

func ProviderTypeFromRecord(record ddlog.Record) (ProviderType, error) {
	rs, err := record.AsStructSafe()
	if err != nil {
		return nil, errors.Wrap(err, "enum ProviderType")
	}
	
	if rs.Name() == "AWSRoute53" {
		return &AWSRoute53{T_: "AWSRoute53"}, nil
	}
	
	if rs.Name() == "AzureDNS" {
		return &AzureDNS{T_: "AzureDNS"}, nil
	}
	
	if rs.Name() == "OpenstackDesignate" {
		return &OpenstackDesignate{T_: "OpenstackDesignate"}, nil
	}
	
	if rs.Name() == "CloudflareDNS" {
		return &CloudflareDNS{T_: "CloudflareDNS"}, nil
	}	
	return nil, errors.Wrap(fmt.Errorf("unexpected record name %s", rs.Name()), "enum ProviderType")
}


var (
	// memory will never be freed, which is fine
	relConstructorAWSRoute53 = ddlog.NewCString("AWSRoute53")
)

type ProviderType_AWSRoute53 interface {
	ProviderType
	internalAWSRoute53()
}

var _ ProviderType = &AWSRoute53{}
var _ ProviderType_AWSRoute53 = &AWSRoute53{}


func NewRecordAWSRoute53(obj *AWSRoute53) ddlog.Record {
	return ddlog.NewRecordStructStatic(relConstructorAWSRoute53)
}

type AWSRoute53 struct {
	T_ string	
}

func (x *AWSRoute53) Name() string {
	return "AWSRoute53"
}

func (x *AWSRoute53) NewRecord() ddlog.Record {
	return NewRecordAWSRoute53(x)
}

func (x *AWSRoute53) internalProviderType() {}

func (x *AWSRoute53) internalAWSRoute53() {}


var (
	// memory will never be freed, which is fine
	relConstructorAzureDNS = ddlog.NewCString("AzureDNS")
)

type ProviderType_AzureDNS interface {
	ProviderType
	internalAzureDNS()
}

var _ ProviderType = &AzureDNS{}
var _ ProviderType_AzureDNS = &AzureDNS{}


func NewRecordAzureDNS(obj *AzureDNS) ddlog.Record {
	return ddlog.NewRecordStructStatic(relConstructorAzureDNS)
}

type AzureDNS struct {
	T_ string	
}

func (x *AzureDNS) Name() string {
	return "AzureDNS"
}

func (x *AzureDNS) NewRecord() ddlog.Record {
	return NewRecordAzureDNS(x)
}

func (x *AzureDNS) internalProviderType() {}

func (x *AzureDNS) internalAzureDNS() {}


var (
	// memory will never be freed, which is fine
	relConstructorOpenstackDesignate = ddlog.NewCString("OpenstackDesignate")
)

type ProviderType_OpenstackDesignate interface {
	ProviderType
	internalOpenstackDesignate()
}

var _ ProviderType = &OpenstackDesignate{}
var _ ProviderType_OpenstackDesignate = &OpenstackDesignate{}


func NewRecordOpenstackDesignate(obj *OpenstackDesignate) ddlog.Record {
	return ddlog.NewRecordStructStatic(relConstructorOpenstackDesignate)
}

type OpenstackDesignate struct {
	T_ string	
}

func (x *OpenstackDesignate) Name() string {
	return "OpenstackDesignate"
}

func (x *OpenstackDesignate) NewRecord() ddlog.Record {
	return NewRecordOpenstackDesignate(x)
}

func (x *OpenstackDesignate) internalProviderType() {}

func (x *OpenstackDesignate) internalOpenstackDesignate() {}


var (
	// memory will never be freed, which is fine
	relConstructorCloudflareDNS = ddlog.NewCString("CloudflareDNS")
)

type ProviderType_CloudflareDNS interface {
	ProviderType
	internalCloudflareDNS()
}

var _ ProviderType = &CloudflareDNS{}
var _ ProviderType_CloudflareDNS = &CloudflareDNS{}


func NewRecordCloudflareDNS(obj *CloudflareDNS) ddlog.Record {
	return ddlog.NewRecordStructStatic(relConstructorCloudflareDNS)
}

type CloudflareDNS struct {
	T_ string	
}

func (x *CloudflareDNS) Name() string {
	return "CloudflareDNS"
}

func (x *CloudflareDNS) NewRecord() ddlog.Record {
	return NewRecordCloudflareDNS(x)
}

func (x *CloudflareDNS) internalProviderType() {}

func (x *CloudflareDNS) internalCloudflareDNS() {}
