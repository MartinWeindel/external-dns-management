package generated

import (
	"github.com/vmware/differential-datalog/go/pkg/ddlog"
)

// Generated code for relation AllDNSOwners
// DO NOT CHANGE MANUALLY

// output relation AllDNSOwners [AllDNSOwners]

func init() {
	meta := &TableMetaData{
		TableName: "AllDNSOwners", 
		RecordName: "AllDNSOwners",
		Unmarshaller: func(record ddlog.Record) (interface{}, error) {
			obj, err := AllDNSOwnersFromRecord(record)
			return obj, err
		},
	}
	registerTableMetaData(meta)
}

func (pd *ProgData) NewInsertCommandAllDNSOwners(obj *AllDNSOwners) ddlog.Command {
	rec := NewRecordAllDNSOwners(obj)
	tableID := pd.LookupTableID("AllDNSOwners")
	return ddlog.NewInsertCommand(tableID, rec)
}

func (pd *ProgData) NewInsertOrUpdateCommandAllDNSOwners(obj *AllDNSOwners) ddlog.Command {
	rec := NewRecordAllDNSOwners(obj)
	tableID := pd.LookupTableID("AllDNSOwners")
	return ddlog.NewInsertOrUpdateCommand(tableID, rec)
}

func (pd *ProgData) NewDeleteValCommandAllDNSOwners(obj *AllDNSOwners) ddlog.Command {
	rec := NewRecordAllDNSOwners(obj)
	tableID := pd.LookupTableID("AllDNSOwners")
	return ddlog.NewDeleteValCommand(tableID, rec)
}

func (pd *ProgData) GetRelTableIDAllDNSOwners() ddlog.TableID {
	return pd.LookupTableID("AllDNSOwners")
}
