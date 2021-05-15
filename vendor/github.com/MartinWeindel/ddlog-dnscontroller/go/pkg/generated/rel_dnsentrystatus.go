package generated

import (
	"github.com/vmware/differential-datalog/go/pkg/ddlog"
)

// Generated code for relation DNSEntryStatus
// DO NOT CHANGE MANUALLY

// output relation DNSEntryStatus [DNSEntryStatus]

func init() {
	meta := &TableMetaData{
		TableName: "DNSEntryStatus", 
		RecordName: "DNSEntryStatus",
		Unmarshaller: func(record ddlog.Record) (interface{}, error) {
			obj, err := DNSEntryStatusFromRecord(record)
			return obj, err
		},
	}
	registerTableMetaData(meta)
}

func (pd *ProgData) NewInsertCommandDNSEntryStatus(obj *DNSEntryStatus) ddlog.Command {
	rec := NewRecordDNSEntryStatus(obj)
	tableID := pd.LookupTableID("DNSEntryStatus")
	return ddlog.NewInsertCommand(tableID, rec)
}

func (pd *ProgData) NewInsertOrUpdateCommandDNSEntryStatus(obj *DNSEntryStatus) ddlog.Command {
	rec := NewRecordDNSEntryStatus(obj)
	tableID := pd.LookupTableID("DNSEntryStatus")
	return ddlog.NewInsertOrUpdateCommand(tableID, rec)
}

func (pd *ProgData) NewDeleteValCommandDNSEntryStatus(obj *DNSEntryStatus) ddlog.Command {
	rec := NewRecordDNSEntryStatus(obj)
	tableID := pd.LookupTableID("DNSEntryStatus")
	return ddlog.NewDeleteValCommand(tableID, rec)
}

func (pd *ProgData) GetRelTableIDDNSEntryStatus() ddlog.TableID {
	return pd.LookupTableID("DNSEntryStatus")
}
