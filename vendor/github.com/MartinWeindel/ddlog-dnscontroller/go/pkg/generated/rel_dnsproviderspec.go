package generated

import (
	"github.com/vmware/differential-datalog/go/pkg/ddlog"
)

// Generated code for relation DNSProviderSpec
// DO NOT CHANGE MANUALLY

// input relation DNSProviderSpec [DNSProviderSpec]

var (
	relTableIDDNSProviderSpec ddlog.TableID = ddlog.GetTableID("DNSProviderSpec")
)

func init() {
	relTableIDDNSProviderSpec = ddlog.GetTableID("DNSProviderSpec")
	meta := &TableMetaData{
		TableID: relTableIDDNSProviderSpec,
		TableName: "DNSProviderSpec", 
		RecordName: "DNSProviderSpec",
		Unmarshaller: func(record ddlog.Record) (interface{}, error) {
			obj, err := DNSProviderSpecFromRecord(record)
			return obj, err
		},
	}
	registerTableMetaData(relTableIDDNSProviderSpec, meta)
}

func NewInsertCommandDNSProviderSpec(obj *DNSProviderSpec) ddlog.Command {
	rec := NewRecordDNSProviderSpec(obj)
	return ddlog.NewInsertCommand(relTableIDDNSProviderSpec, rec)
}
