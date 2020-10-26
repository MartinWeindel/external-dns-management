package generated

// Generated code for typedef AccountInUse
// DO NOT CHANGE MANUALLY

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/vmware/differential-datalog/go/pkg/ddlog"
)

type UnmarshalRecord func(record ddlog.Record) (interface{}, error)

type TableMetaData struct {
	TableID ddlog.TableID
	TableName string
	RecordName string
	Unmarshaller UnmarshalRecord
}

var tableMetaDataRegistry map[ddlog.TableID]*TableMetaData

// TODO report/fix bug in record.go
func init() {
	tableMetaDataRegistry = map[ddlog.TableID]*TableMetaData{}

	ddlog.StdSomeConstructor = ddlog.NewCString("ddlog_std::Some")
	ddlog.StdNoneConstructor = ddlog.NewCString("ddlog_std::None")
	ddlog.StdLeftConstructor = ddlog.NewCString("ddlog_std::Left")
	ddlog.StdRightConstructor = ddlog.NewCString("ddlog_std::Right")
}

func registerTableMetaData(tableID ddlog.TableID, meta *TableMetaData) {
	tableMetaDataRegistry[tableID] = meta
}

func LookupTableMetaData(tableID ddlog.TableID) (*TableMetaData, error) {
	meta := tableMetaDataRegistry[tableID]
	if meta == nil {
		return nil, fmt.Errorf("tableID %d not found", tableID)
	}
	return meta, nil
}


var (
	// memory will never be freed, which is fine
	relConstructorAccountInUse = ddlog.NewCString("AccountInUse")
)


func NewRecordAccountInUse(obj *AccountInUse) ddlog.Record {
	arg0 := func() ddlog.Record {
	    return ddlog.NewRecordString(obj.CredentialsHash)
    }()
	arg1 := func() ddlog.Record {
	    return NewRecordProviderType(obj.ProviderType)
    }()
	return ddlog.NewRecordStructStatic(relConstructorAccountInUse, arg0, arg1)
}


func AccountInUseFromRecord(record ddlog.Record) (*AccountInUse, error) {
	rs, err := record.AsStructSafe()
	if err != nil {
		return nil, errors.Wrapf(err, "no struct (AccountInUse)")
	}
	if rs.Name() != "AccountInUse" {
		return nil, fmt.Errorf("unexpected record %s != AccountInUse", rs.Name())
	}
	arg0, err := rs.At(0).ToStringSafe()
	if err != nil {
		return nil, errors.Wrapf(err, "Field credentialsHash")
	}
	arg1, err := ProviderTypeFromRecord(rs.At(1))
	if err != nil {
		return nil, errors.Wrapf(err, "Field providerType")
	}
	obj := &AccountInUse{	
		CredentialsHash:arg0,	
		ProviderType:arg1,
	}
	return obj, nil
}

type AccountInUse struct {
    CredentialsHash string
    ProviderType ProviderType
}
