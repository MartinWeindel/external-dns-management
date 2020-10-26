package generated

// Generated code for typedef EntryStatus
// DO NOT CHANGE MANUALLY

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/vmware/differential-datalog/go/pkg/ddlog"
)

type EntryStatus interface {
	internalEntryStatus()
	NewRecord() ddlog.Record
}

func NewRecordEntryStatus(obj EntryStatus) ddlog.Record {
	return obj.NewRecord()
}

func EntryStatusFromRecord(record ddlog.Record) (EntryStatus, error) {
	rs, err := record.AsStructSafe()
	if err != nil {
		return nil, errors.Wrap(err, "enum EntryStatus")
	}
	
	if rs.Name() == "Unchanged" {
		return &Unchanged{}, nil
	}
	
	if rs.Name() == "Inserting" {
		return &Inserting{}, nil
	}
	
	if rs.Name() == "Updating" {
		return &Updating{}, nil
	}
	
	if rs.Name() == "ForeignOwner" {
		return &ForeignOwner{}, nil
	}
	
	if rs.Name() == "OwnerConflict" {
		return &OwnerConflict{}, nil
	}
	
	if rs.Name() == "NoProvider" {
		return &NoProvider{}, nil
	}	
	return nil, errors.Wrap(fmt.Errorf("unexpected record name %s", rs.Name()), "enum EntryStatus")
}


var (
	// memory will never be freed, which is fine
	relConstructorUnchanged = ddlog.NewCString("Unchanged")
)

type EntryStatus_Unchanged interface {
	EntryStatus
	internalUnchanged()
}

var _ EntryStatus = &Unchanged{}
var _ EntryStatus_Unchanged = &Unchanged{}


func NewRecordUnchanged(obj *Unchanged) ddlog.Record {
	return ddlog.NewRecordStructStatic(relConstructorUnchanged)
}

type Unchanged struct {	
}

func (x *Unchanged) NewRecord() ddlog.Record {
	return NewRecordUnchanged(x)
}

func (x *Unchanged) internalEntryStatus() {}

func (x *Unchanged) internalUnchanged() {}


var (
	// memory will never be freed, which is fine
	relConstructorInserting = ddlog.NewCString("Inserting")
)

type EntryStatus_Inserting interface {
	EntryStatus
	internalInserting()
}

var _ EntryStatus = &Inserting{}
var _ EntryStatus_Inserting = &Inserting{}


func NewRecordInserting(obj *Inserting) ddlog.Record {
	return ddlog.NewRecordStructStatic(relConstructorInserting)
}

type Inserting struct {	
}

func (x *Inserting) NewRecord() ddlog.Record {
	return NewRecordInserting(x)
}

func (x *Inserting) internalEntryStatus() {}

func (x *Inserting) internalInserting() {}


var (
	// memory will never be freed, which is fine
	relConstructorUpdating = ddlog.NewCString("Updating")
)

type EntryStatus_Updating interface {
	EntryStatus
	internalUpdating()
}

var _ EntryStatus = &Updating{}
var _ EntryStatus_Updating = &Updating{}


func NewRecordUpdating(obj *Updating) ddlog.Record {
	return ddlog.NewRecordStructStatic(relConstructorUpdating)
}

type Updating struct {	
}

func (x *Updating) NewRecord() ddlog.Record {
	return NewRecordUpdating(x)
}

func (x *Updating) internalEntryStatus() {}

func (x *Updating) internalUpdating() {}


var (
	// memory will never be freed, which is fine
	relConstructorForeignOwner = ddlog.NewCString("ForeignOwner")
)

type EntryStatus_ForeignOwner interface {
	EntryStatus
	internalForeignOwner()
}

var _ EntryStatus = &ForeignOwner{}
var _ EntryStatus_ForeignOwner = &ForeignOwner{}


func NewRecordForeignOwner(obj *ForeignOwner) ddlog.Record {
	arg0 := func() ddlog.Record {
	    return ddlog.NewRecordString(obj.Owner)
    }()
	return ddlog.NewRecordStructStatic(relConstructorForeignOwner, arg0)
}

type ForeignOwner struct {
    Owner string	
}

func (x *ForeignOwner) NewRecord() ddlog.Record {
	return NewRecordForeignOwner(x)
}

func (x *ForeignOwner) internalEntryStatus() {}

func (x *ForeignOwner) internalForeignOwner() {}


var (
	// memory will never be freed, which is fine
	relConstructorOwnerConflict = ddlog.NewCString("OwnerConflict")
)

type EntryStatus_OwnerConflict interface {
	EntryStatus
	internalOwnerConflict()
}

var _ EntryStatus = &OwnerConflict{}
var _ EntryStatus_OwnerConflict = &OwnerConflict{}


func NewRecordOwnerConflict(obj *OwnerConflict) ddlog.Record {
	arg0 := func() ddlog.Record {
	    return ddlog.NewRecordString(obj.Owner)
    }()
	return ddlog.NewRecordStructStatic(relConstructorOwnerConflict, arg0)
}

type OwnerConflict struct {
    Owner string	
}

func (x *OwnerConflict) NewRecord() ddlog.Record {
	return NewRecordOwnerConflict(x)
}

func (x *OwnerConflict) internalEntryStatus() {}

func (x *OwnerConflict) internalOwnerConflict() {}


var (
	// memory will never be freed, which is fine
	relConstructorNoProvider = ddlog.NewCString("NoProvider")
)

type EntryStatus_NoProvider interface {
	EntryStatus
	internalNoProvider()
}

var _ EntryStatus = &NoProvider{}
var _ EntryStatus_NoProvider = &NoProvider{}


func NewRecordNoProvider(obj *NoProvider) ddlog.Record {
	return ddlog.NewRecordStructStatic(relConstructorNoProvider)
}

type NoProvider struct {	
}

func (x *NoProvider) NewRecord() ddlog.Record {
	return NewRecordNoProvider(x)
}

func (x *NoProvider) internalEntryStatus() {}

func (x *NoProvider) internalNoProvider() {}
