package generated

import (
	"github.com/vmware/differential-datalog/go/pkg/ddlog"
)

// Generated code for relation DNSEntrySpec
// DO NOT CHANGE MANUALLY

// input relation DNSEntrySpec [DNSEntrySpec]

var (
	relTableIDDNSEntrySpec ddlog.TableID = ddlog.GetTableID("DNSEntrySpec")
)

func init() {
	relTableIDDNSEntrySpec = ddlog.GetTableID("DNSEntrySpec")
	meta := &TableMetaData{
		TableID: relTableIDDNSEntrySpec,
		TableName: "DNSEntrySpec", 
		RecordName: "DNSEntrySpec",
		Unmarshaller: func(record ddlog.Record) (interface{}, error) {
			obj, err := DNSEntrySpecFromRecord(record)
			return obj, err
		},
	}
	registerTableMetaData(relTableIDDNSEntrySpec, meta)
}

func NewInsertCommandDNSEntrySpec(obj *DNSEntrySpec) ddlog.Command {
	rec := NewRecordDNSEntrySpec(obj)
	return ddlog.NewInsertCommand(relTableIDDNSEntrySpec, rec)
}
