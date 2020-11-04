package generated

// Generated code for typedef RecordType
// DO NOT CHANGE MANUALLY

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/vmware/differential-datalog/go/pkg/ddlog"
)

type RecordType interface {
	internalRecordType()
	NewRecord() ddlog.Record
}

func NewRecordRecordType(obj RecordType) ddlog.Record {
	return obj.NewRecord()
}

func RecordTypeFromRecord(record ddlog.Record) (RecordType, error) {
	rs, err := record.AsStructSafe()
	if err != nil {
		return nil, errors.Wrap(err, "enum RecordType")
	}
	
	if rs.Name() == "A" {
		return &A{T_: "A"}, nil
	}
	
	if rs.Name() == "CNAME" {
		return &CNAME{T_: "CNAME"}, nil
	}
	
	if rs.Name() == "TXT" {
		return &TXT{T_: "TXT"}, nil
	}
	
	if rs.Name() == "NS" {
		return &NS{T_: "NS"}, nil
	}	
	return nil, errors.Wrap(fmt.Errorf("unexpected record name %s", rs.Name()), "enum RecordType")
}


var (
	// memory will never be freed, which is fine
	relConstructorA = ddlog.NewCString("A")
)

type RecordType_A interface {
	RecordType
	internalA()
}

var _ RecordType = &A{}
var _ RecordType_A = &A{}


func NewRecordA(obj *A) ddlog.Record {
	return ddlog.NewRecordStructStatic(relConstructorA)
}

type A struct {
	T_ string	
}

func (x *A) NewRecord() ddlog.Record {
	return NewRecordA(x)
}

func (x *A) internalRecordType() {}

func (x *A) internalA() {}


var (
	// memory will never be freed, which is fine
	relConstructorCNAME = ddlog.NewCString("CNAME")
)

type RecordType_CNAME interface {
	RecordType
	internalCNAME()
}

var _ RecordType = &CNAME{}
var _ RecordType_CNAME = &CNAME{}


func NewRecordCNAME(obj *CNAME) ddlog.Record {
	return ddlog.NewRecordStructStatic(relConstructorCNAME)
}

type CNAME struct {
	T_ string	
}

func (x *CNAME) NewRecord() ddlog.Record {
	return NewRecordCNAME(x)
}

func (x *CNAME) internalRecordType() {}

func (x *CNAME) internalCNAME() {}


var (
	// memory will never be freed, which is fine
	relConstructorTXT = ddlog.NewCString("TXT")
)

type RecordType_TXT interface {
	RecordType
	internalTXT()
}

var _ RecordType = &TXT{}
var _ RecordType_TXT = &TXT{}


func NewRecordTXT(obj *TXT) ddlog.Record {
	return ddlog.NewRecordStructStatic(relConstructorTXT)
}

type TXT struct {
	T_ string	
}

func (x *TXT) NewRecord() ddlog.Record {
	return NewRecordTXT(x)
}

func (x *TXT) internalRecordType() {}

func (x *TXT) internalTXT() {}


var (
	// memory will never be freed, which is fine
	relConstructorNS = ddlog.NewCString("NS")
)

type RecordType_NS interface {
	RecordType
	internalNS()
}

var _ RecordType = &NS{}
var _ RecordType_NS = &NS{}


func NewRecordNS(obj *NS) ddlog.Record {
	return ddlog.NewRecordStructStatic(relConstructorNS)
}

type NS struct {
	T_ string	
}

func (x *NS) NewRecord() ddlog.Record {
	return NewRecordNS(x)
}

func (x *NS) internalRecordType() {}

func (x *NS) internalNS() {}
