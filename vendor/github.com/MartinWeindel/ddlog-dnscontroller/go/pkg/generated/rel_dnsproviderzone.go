package generated

import (
	"github.com/vmware/differential-datalog/go/pkg/ddlog"
)

// Generated code for relation DNSProviderZone
// DO NOT CHANGE MANUALLY

// output relation DNSProviderZone [DNSProviderZone]

var (
	relTableIDDNSProviderZone ddlog.TableID = ddlog.GetTableID("DNSProviderZone")
)

func init() {
	relTableIDDNSProviderZone = ddlog.GetTableID("DNSProviderZone")
	meta := &TableMetaData{
		TableID: relTableIDDNSProviderZone,
		TableName: "DNSProviderZone", 
		RecordName: "DNSProviderZone",
		Unmarshaller: func(record ddlog.Record) (interface{}, error) {
			obj, err := DNSProviderZoneFromRecord(record)
			return obj, err
		},
	}
	registerTableMetaData(relTableIDDNSProviderZone, meta)
}

func NewInsertCommandDNSProviderZone(obj *DNSProviderZone) ddlog.Command {
	rec := NewRecordDNSProviderZone(obj)
	return ddlog.NewInsertCommand(relTableIDDNSProviderZone, rec)
}

func NewInsertOrUpdateCommandDNSProviderZone(obj *DNSProviderZone) ddlog.Command {
	rec := NewRecordDNSProviderZone(obj)
	return ddlog.NewInsertOrUpdateCommand(relTableIDDNSProviderZone, rec)
}

func NewDeleteValCommandDNSProviderZone(obj *DNSProviderZone) ddlog.Command {
	rec := NewRecordDNSProviderZone(obj)
	return ddlog.NewDeleteValCommand(relTableIDDNSProviderZone, rec)
}

func GetRelTableIDDNSProviderZone() ddlog.TableID {
	return relTableIDDNSProviderZone
}
