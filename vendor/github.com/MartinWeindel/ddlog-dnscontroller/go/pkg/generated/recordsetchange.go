package generated

// Generated code for typedef RecordSetChange
// DO NOT CHANGE MANUALLY

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/vmware/differential-datalog/go/pkg/ddlog"
)


var (
	// memory will never be freed, which is fine
	relConstructorRecordSetChange = ddlog.NewCString("RecordSetChange")
)


func NewRecordRecordSetChange(obj *RecordSetChange) ddlog.Record {
	arg0 := func() ddlog.Record {
	    return NewRecordProviderType(obj.Ptype)
    }()
	arg1 := func() ddlog.Record {
	    return ddlog.NewRecordString(obj.Zoneid)
    }()
	arg2 := func() ddlog.Record {
	    vec := make([]ddlog.Record, len(obj.Insert))
	for i, item := range obj.Insert {
		vec[i] = NewRecordRecordSetData(&item)
	}
    return ddlog.NewRecordVector(vec...)
    }()
	arg3 := func() ddlog.Record {
	    vec := make([]ddlog.Record, len(obj.Update))
	for i, item := range obj.Update {
		vec[i] = NewRecordRecordSetData(&item)
	}
    return ddlog.NewRecordVector(vec...)
    }()
	arg4 := func() ddlog.Record {
	    vec := make([]ddlog.Record, len(obj.Delete))
	for i, item := range obj.Delete {
		vec[i] = NewRecordRecordSetData(&item)
	}
    return ddlog.NewRecordVector(vec...)
    }()
	arg5 := func() ddlog.Record {
	    vec := make([]ddlog.Record, 0, len(obj.EntryStatus))
	for k, v := range obj.EntryStatus {
		rec_k := NewRecordObjectKey(&k)
		rec_v := NewRecordEntryStatus(v)
		rec := ddlog.NewRecordPair(rec_k, rec_v)
		vec = append(vec, rec)
	}
	return ddlog.NewRecordVector(vec...)
    }()
	return ddlog.NewRecordStructStatic(relConstructorRecordSetChange, arg0, arg1, arg2, arg3, arg4, arg5)
}


func RecordSetChangeFromRecord(record ddlog.Record) (*RecordSetChange, error) {
	rs, err := record.AsStructSafe()
	if err != nil {
		return nil, errors.Wrapf(err, "no struct (RecordSetChange)")
	}
	if rs.Name() != "RecordSetChange" {
		return nil, fmt.Errorf("unexpected record %s != RecordSetChange", rs.Name())
	}
	arg0, err := ProviderTypeFromRecord(rs.At(0))
	if err != nil {
		return nil, errors.Wrapf(err, "Field ptype")
	}
	arg1, err := rs.At(1).ToStringSafe()
	if err != nil {
		return nil, errors.Wrapf(err, "Field zoneid")
	}
	arg2, err := 
	func() ([]RecordSetData, error) {
		rv0 := rs.At(2)
		rv, err := rv0.AsVectorSafe()
		if err != nil {
			return nil, err
		}
		vec := make([]RecordSetData, rv.Size())
		for i := 0; i < len(vec); i++ {
			obj, err := RecordSetDataFromRecord(rv.At(i))
			if err != nil {
				errors.Wrapf(err, "vector index %d", i)
			}
			vec[i] =*obj
		}
		return vec, nil
	}()
	if err != nil {
		return nil, errors.Wrapf(err, "Field insert")
	}
	arg3, err := 
	func() ([]RecordSetData, error) {
		rv0 := rs.At(3)
		rv, err := rv0.AsVectorSafe()
		if err != nil {
			return nil, err
		}
		vec := make([]RecordSetData, rv.Size())
		for i := 0; i < len(vec); i++ {
			obj, err := RecordSetDataFromRecord(rv.At(i))
			if err != nil {
				errors.Wrapf(err, "vector index %d", i)
			}
			vec[i] =*obj
		}
		return vec, nil
	}()
	if err != nil {
		return nil, errors.Wrapf(err, "Field update")
	}
	arg4, err := 
	func() ([]RecordSetData, error) {
		rv0 := rs.At(4)
		rv, err := rv0.AsVectorSafe()
		if err != nil {
			return nil, err
		}
		vec := make([]RecordSetData, rv.Size())
		for i := 0; i < len(vec); i++ {
			obj, err := RecordSetDataFromRecord(rv.At(i))
			if err != nil {
				errors.Wrapf(err, "vector index %d", i)
			}
			vec[i] =*obj
		}
		return vec, nil
	}()
	if err != nil {
		return nil, errors.Wrapf(err, "Field delete")
	}
	arg5, err := func() (map[ObjectKey]EntryStatus, error) {
		rv0 := rs.At(5)
		rv, err := rv0.AsMapSafe()
		if err != nil {
			return nil, err
		}
		result := map[ObjectKey]EntryStatus{}
		for i := 0; i < rv.Size(); i++ {
			key := rv.KeyAt(i)
			value := rv.ValueAt(i)
			k, err := ObjectKeyFromRecord(key)
			if err != nil {
				errors.Wrapf(err, "map key(%d)", i)
			}
			v, err := EntryStatusFromRecord(value)
			if err != nil {
				errors.Wrapf(err, "map value(%d)", i)
			}
			result[*k] =v
		}
		return result, nil
	}()
	if err != nil {
		return nil, errors.Wrapf(err, "Field entryStatus")
	}
	obj := &RecordSetChange{	
		Ptype:arg0,	
		Zoneid:arg1,	
		Insert:arg2,	
		Update:arg3,	
		Delete:arg4,	
		EntryStatus:arg5,
	}
	return obj, nil
}

type RecordSetChange struct {
    Ptype ProviderType
    Zoneid string
    Insert []RecordSetData
    Update []RecordSetData
    Delete []RecordSetData
    EntryStatus map[ObjectKey]EntryStatus
}
