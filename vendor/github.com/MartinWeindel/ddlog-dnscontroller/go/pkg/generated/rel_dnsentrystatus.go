package generated

import (
	"github.com/vmware/differential-datalog/go/pkg/ddlog"
)

// Generated code for relation DNSEntryStatus
// DO NOT CHANGE MANUALLY

// output relation DNSEntryStatus [DNSEntryStatus]

var (
	relTableIDDNSEntryStatus ddlog.TableID = ddlog.GetTableID("DNSEntryStatus")
)

func init() {
	relTableIDDNSEntryStatus = ddlog.GetTableID("DNSEntryStatus")
	meta := &TableMetaData{
		TableID: relTableIDDNSEntryStatus,
		TableName: "DNSEntryStatus", 
		RecordName: "DNSEntryStatus",
		Unmarshaller: func(record ddlog.Record) (interface{}, error) {
			obj, err := DNSEntryStatusFromRecord(record)
			return obj, err
		},
	}
	registerTableMetaData(relTableIDDNSEntryStatus, meta)
}

func NewInsertCommandDNSEntryStatus(obj *DNSEntryStatus) ddlog.Command {
	rec := NewRecordDNSEntryStatus(obj)
	return ddlog.NewInsertCommand(relTableIDDNSEntryStatus, rec)
}

func NewInsertOrUpdateCommandDNSEntryStatus(obj *DNSEntryStatus) ddlog.Command {
	rec := NewRecordDNSEntryStatus(obj)
	return ddlog.NewInsertOrUpdateCommand(relTableIDDNSEntryStatus, rec)
}

func NewDeleteValCommandDNSEntryStatus(obj *DNSEntryStatus) ddlog.Command {
	rec := NewRecordDNSEntryStatus(obj)
	return ddlog.NewDeleteValCommand(relTableIDDNSEntryStatus, rec)
}

func GetRelTableIDDNSEntryStatus() ddlog.TableID {
	return relTableIDDNSEntryStatus
}
