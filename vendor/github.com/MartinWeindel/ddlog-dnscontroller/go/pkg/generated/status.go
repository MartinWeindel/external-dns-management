package generated

// Generated code for typedef Status
// DO NOT CHANGE MANUALLY

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/vmware/differential-datalog/go/pkg/ddlog"
)

type Status interface {
	internalStatus()
	NewRecord() ddlog.Record
}

func NewRecordStatus(obj Status) ddlog.Record {
	return obj.NewRecord()
}

func StatusFromRecord(record ddlog.Record) (Status, error) {
	rs, err := record.AsStructSafe()
	if err != nil {
		return nil, errors.Wrap(err, "enum Status")
	}
	
	if rs.Name() == "Ready" {
		return &Ready{}, nil
	}
	
	if rs.Name() == "Stale" {
		return &Stale{}, nil
	}
	
	if rs.Name() == "Pending" {
		return &Pending{}, nil
	}
	
	if rs.Name() == "Error" {
		return &Error{}, nil
	}	
	return nil, errors.Wrap(fmt.Errorf("unexpected record name %s", rs.Name()), "enum Status")
}


var (
	// memory will never be freed, which is fine
	relConstructorReady = ddlog.NewCString("Ready")
)

type Status_Ready interface {
	Status
	internalReady()
}

var _ Status = &Ready{}
var _ Status_Ready = &Ready{}


func NewRecordReady(obj *Ready) ddlog.Record {
	return ddlog.NewRecordStructStatic(relConstructorReady)
}

type Ready struct {	
}

func (x *Ready) NewRecord() ddlog.Record {
	return NewRecordReady(x)
}

func (x *Ready) internalStatus() {}

func (x *Ready) internalReady() {}


var (
	// memory will never be freed, which is fine
	relConstructorStale = ddlog.NewCString("Stale")
)

type Status_Stale interface {
	Status
	internalStale()
}

var _ Status = &Stale{}
var _ Status_Stale = &Stale{}


func NewRecordStale(obj *Stale) ddlog.Record {
	return ddlog.NewRecordStructStatic(relConstructorStale)
}

type Stale struct {	
}

func (x *Stale) NewRecord() ddlog.Record {
	return NewRecordStale(x)
}

func (x *Stale) internalStatus() {}

func (x *Stale) internalStale() {}


var (
	// memory will never be freed, which is fine
	relConstructorPending = ddlog.NewCString("Pending")
)

type Status_Pending interface {
	Status
	internalPending()
}

var _ Status = &Pending{}
var _ Status_Pending = &Pending{}


func NewRecordPending(obj *Pending) ddlog.Record {
	return ddlog.NewRecordStructStatic(relConstructorPending)
}

type Pending struct {	
}

func (x *Pending) NewRecord() ddlog.Record {
	return NewRecordPending(x)
}

func (x *Pending) internalStatus() {}

func (x *Pending) internalPending() {}


var (
	// memory will never be freed, which is fine
	relConstructorError = ddlog.NewCString("Error")
)

type Status_Error interface {
	Status
	internalError()
}

var _ Status = &Error{}
var _ Status_Error = &Error{}


func NewRecordError(obj *Error) ddlog.Record {
	return ddlog.NewRecordStructStatic(relConstructorError)
}

type Error struct {	
}

func (x *Error) NewRecord() ddlog.Record {
	return NewRecordError(x)
}

func (x *Error) internalStatus() {}

func (x *Error) internalError() {}
