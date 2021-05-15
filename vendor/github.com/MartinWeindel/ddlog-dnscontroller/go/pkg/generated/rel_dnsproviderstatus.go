package generated

import (
	"github.com/vmware/differential-datalog/go/pkg/ddlog"
)

// Generated code for relation DNSProviderStatus
// DO NOT CHANGE MANUALLY

// output relation DNSProviderStatus [DNSProviderStatus]

func init() {
	meta := &TableMetaData{
		TableName: "DNSProviderStatus", 
		RecordName: "DNSProviderStatus",
		Unmarshaller: func(record ddlog.Record) (interface{}, error) {
			obj, err := DNSProviderStatusFromRecord(record)
			return obj, err
		},
	}
	registerTableMetaData(meta)
}

func (pd *ProgData) NewInsertCommandDNSProviderStatus(obj *DNSProviderStatus) ddlog.Command {
	rec := NewRecordDNSProviderStatus(obj)
	tableID := pd.LookupTableID("DNSProviderStatus")
	return ddlog.NewInsertCommand(tableID, rec)
}

func (pd *ProgData) NewInsertOrUpdateCommandDNSProviderStatus(obj *DNSProviderStatus) ddlog.Command {
	rec := NewRecordDNSProviderStatus(obj)
	tableID := pd.LookupTableID("DNSProviderStatus")
	return ddlog.NewInsertOrUpdateCommand(tableID, rec)
}

func (pd *ProgData) NewDeleteValCommandDNSProviderStatus(obj *DNSProviderStatus) ddlog.Command {
	rec := NewRecordDNSProviderStatus(obj)
	tableID := pd.LookupTableID("DNSProviderStatus")
	return ddlog.NewDeleteValCommand(tableID, rec)
}

func (pd *ProgData) GetRelTableIDDNSProviderStatus() ddlog.TableID {
	return pd.LookupTableID("DNSProviderStatus")
}
