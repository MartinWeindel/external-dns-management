package generated

import (
	"github.com/vmware/differential-datalog/go/pkg/ddlog"
)

// Generated code for relation DNSProviderSpec
// DO NOT CHANGE MANUALLY

// input relation DNSProviderSpec [DNSProviderSpec]

func init() {
	meta := &TableMetaData{
		TableName: "DNSProviderSpec", 
		RecordName: "DNSProviderSpec",
		Unmarshaller: func(record ddlog.Record) (interface{}, error) {
			obj, err := DNSProviderSpecFromRecord(record)
			return obj, err
		},
	}
	registerTableMetaData(meta)
}

func (pd *ProgData) NewInsertCommandDNSProviderSpec(obj *DNSProviderSpec) ddlog.Command {
	rec := NewRecordDNSProviderSpec(obj)
	tableID := pd.LookupTableID("DNSProviderSpec")
	return ddlog.NewInsertCommand(tableID, rec)
}

func (pd *ProgData) NewInsertOrUpdateCommandDNSProviderSpec(obj *DNSProviderSpec) ddlog.Command {
	rec := NewRecordDNSProviderSpec(obj)
	tableID := pd.LookupTableID("DNSProviderSpec")
	return ddlog.NewInsertOrUpdateCommand(tableID, rec)
}

func (pd *ProgData) NewDeleteValCommandDNSProviderSpec(obj *DNSProviderSpec) ddlog.Command {
	rec := NewRecordDNSProviderSpec(obj)
	tableID := pd.LookupTableID("DNSProviderSpec")
	return ddlog.NewDeleteValCommand(tableID, rec)
}
func (pd *ProgData) NewDeleteKeyCommandDNSProviderSpec(obj *DNSProviderSpec) ddlog.Command {
	rec := NewKeyRecordDNSProviderSpec(obj)
	tableID := pd.LookupTableID("DNSProviderSpec")
	return ddlog.NewDeleteKeyCommand(tableID, rec)
}


func NewKeyRecordDNSProviderSpec(obj *DNSProviderSpec) ddlog.Record {
	return func() ddlog.Record {
	    return NewRecordObjectKey(&obj.Key)
    }()
}

func (pd *ProgData) GetRelTableIDDNSProviderSpec() ddlog.TableID {
	return pd.LookupTableID("DNSProviderSpec")
}
