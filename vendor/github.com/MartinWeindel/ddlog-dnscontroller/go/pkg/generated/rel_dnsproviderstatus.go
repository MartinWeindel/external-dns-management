package generated

import (
	"github.com/vmware/differential-datalog/go/pkg/ddlog"
)

// Generated code for relation DNSProviderStatus
// DO NOT CHANGE MANUALLY

// output relation DNSProviderStatus [DNSProviderStatus]

var (
	relTableIDDNSProviderStatus ddlog.TableID = ddlog.GetTableID("DNSProviderStatus")
)

func init() {
	relTableIDDNSProviderStatus = ddlog.GetTableID("DNSProviderStatus")
	meta := &TableMetaData{
		TableID: relTableIDDNSProviderStatus,
		TableName: "DNSProviderStatus", 
		RecordName: "DNSProviderStatus",
		Unmarshaller: func(record ddlog.Record) (interface{}, error) {
			obj, err := DNSProviderStatusFromRecord(record)
			return obj, err
		},
	}
	registerTableMetaData(relTableIDDNSProviderStatus, meta)
}

func NewInsertCommandDNSProviderStatus(obj *DNSProviderStatus) ddlog.Command {
	rec := NewRecordDNSProviderStatus(obj)
	return ddlog.NewInsertCommand(relTableIDDNSProviderStatus, rec)
}

func NewInsertOrUpdateCommandDNSProviderStatus(obj *DNSProviderStatus) ddlog.Command {
	rec := NewRecordDNSProviderStatus(obj)
	return ddlog.NewInsertOrUpdateCommand(relTableIDDNSProviderStatus, rec)
}

func NewDeleteValCommandDNSProviderStatus(obj *DNSProviderStatus) ddlog.Command {
	rec := NewRecordDNSProviderStatus(obj)
	return ddlog.NewDeleteValCommand(relTableIDDNSProviderStatus, rec)
}

func GetRelTableIDDNSProviderStatus() ddlog.TableID {
	return relTableIDDNSProviderStatus
}
