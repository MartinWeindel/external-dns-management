package generated

import (
	"github.com/vmware/differential-datalog/go/pkg/ddlog"
)

// Generated code for relation DNSProviderZone
// DO NOT CHANGE MANUALLY

// output relation DNSProviderZone [DNSProviderZone]

func init() {
	meta := &TableMetaData{
		TableName: "DNSProviderZone", 
		RecordName: "DNSProviderZone",
		Unmarshaller: func(record ddlog.Record) (interface{}, error) {
			obj, err := DNSProviderZoneFromRecord(record)
			return obj, err
		},
	}
	registerTableMetaData(meta)
}

func (pd *ProgData) NewInsertCommandDNSProviderZone(obj *DNSProviderZone) ddlog.Command {
	rec := NewRecordDNSProviderZone(obj)
	tableID := pd.LookupTableID("DNSProviderZone")
	return ddlog.NewInsertCommand(tableID, rec)
}

func (pd *ProgData) NewInsertOrUpdateCommandDNSProviderZone(obj *DNSProviderZone) ddlog.Command {
	rec := NewRecordDNSProviderZone(obj)
	tableID := pd.LookupTableID("DNSProviderZone")
	return ddlog.NewInsertOrUpdateCommand(tableID, rec)
}

func (pd *ProgData) NewDeleteValCommandDNSProviderZone(obj *DNSProviderZone) ddlog.Command {
	rec := NewRecordDNSProviderZone(obj)
	tableID := pd.LookupTableID("DNSProviderZone")
	return ddlog.NewDeleteValCommand(tableID, rec)
}

func (pd *ProgData) GetRelTableIDDNSProviderZone() ddlog.TableID {
	return pd.LookupTableID("DNSProviderZone")
}
