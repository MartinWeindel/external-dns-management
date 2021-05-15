package generated

// Generated code for typedef AccountInUse
// DO NOT CHANGE MANUALLY

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/vmware/differential-datalog/go/pkg/ddlog"
)

// TODO report/fix bug in record.go
func init() {
	ddlog.StdSomeConstructor = ddlog.NewCString("ddlog_std::Some")
	ddlog.StdNoneConstructor = ddlog.NewCString("ddlog_std::None")
	ddlog.StdLeftConstructor = ddlog.NewCString("ddlog_std::Left")
	ddlog.StdRightConstructor = ddlog.NewCString("ddlog_std::Right")
}

type UnmarshalRecord func(record ddlog.Record) (interface{}, error)

type TableMetaData struct {
	TableID      ddlog.TableID
	TableName    string
	RecordName   string
	Unmarshaller UnmarshalRecord
}

var progDataTableMetaDataRegistry []*TableMetaData

type ProgData struct {
	prog          *ddlog.Program
	tableToID     map[string]ddlog.TableID
	tableMetaData map[ddlog.TableID]*TableMetaData
}

func NewProgData(prog *ddlog.Program) *ProgData {
	pd := &ProgData{
		prog:          prog,
		tableToID:     map[string]ddlog.TableID{},
	    tableMetaData: map[ddlog.TableID]*TableMetaData{},
	}
	for _, meta := range progDataTableMetaDataRegistry {
		meta.TableID = prog.GetTableID(meta.TableName)
		pd.tableToID[meta.TableName] = meta.TableID
		pd.tableMetaData[meta.TableID] = meta
	}
	return pd
}

func registerTableMetaData(meta *TableMetaData) {
	progDataTableMetaDataRegistry = append(progDataTableMetaDataRegistry, meta)
}

func (pd *ProgData) LookupTableID(tableName string) ddlog.TableID {
	if tableID, ok := pd.tableToID[tableName]; ok {
		return tableID
	}
	panic(fmt.Sprintf("no tableID for %s", tableName))
}

func (pd *ProgData) LookupTableMetaData(tableID ddlog.TableID) (*TableMetaData, error) {
	meta := pd.tableMetaData[tableID]
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
