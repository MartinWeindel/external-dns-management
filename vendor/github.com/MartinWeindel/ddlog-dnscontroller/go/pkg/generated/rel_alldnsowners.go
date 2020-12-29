package generated

import (
	"github.com/vmware/differential-datalog/go/pkg/ddlog"
)

// Generated code for relation AllDNSOwners
// DO NOT CHANGE MANUALLY

// output relation AllDNSOwners [AllDNSOwners]

var (
	relTableIDAllDNSOwners ddlog.TableID = ddlog.GetTableID("AllDNSOwners")
)

func init() {
	relTableIDAllDNSOwners = ddlog.GetTableID("AllDNSOwners")
	meta := &TableMetaData{
		TableID: relTableIDAllDNSOwners,
		TableName: "AllDNSOwners", 
		RecordName: "AllDNSOwners",
		Unmarshaller: func(record ddlog.Record) (interface{}, error) {
			obj, err := AllDNSOwnersFromRecord(record)
			return obj, err
		},
	}
	registerTableMetaData(relTableIDAllDNSOwners, meta)
}

func NewInsertCommandAllDNSOwners(obj *AllDNSOwners) ddlog.Command {
	rec := NewRecordAllDNSOwners(obj)
	return ddlog.NewInsertCommand(relTableIDAllDNSOwners, rec)
}

func NewInsertOrUpdateCommandAllDNSOwners(obj *AllDNSOwners) ddlog.Command {
	rec := NewRecordAllDNSOwners(obj)
	return ddlog.NewInsertOrUpdateCommand(relTableIDAllDNSOwners, rec)
}

func NewDeleteValCommandAllDNSOwners(obj *AllDNSOwners) ddlog.Command {
	rec := NewRecordAllDNSOwners(obj)
	return ddlog.NewDeleteValCommand(relTableIDAllDNSOwners, rec)
}

func GetRelTableIDAllDNSOwners() ddlog.TableID {
	return relTableIDAllDNSOwners
}
