package generated

import (
	"github.com/vmware/differential-datalog/go/pkg/ddlog"
)

// Generated code for relation DNSEntrySpec
// DO NOT CHANGE MANUALLY

// input relation DNSEntrySpec [DNSEntrySpec]

func init() {
	meta := &TableMetaData{
		TableName: "DNSEntrySpec", 
		RecordName: "DNSEntrySpec",
		Unmarshaller: func(record ddlog.Record) (interface{}, error) {
			obj, err := DNSEntrySpecFromRecord(record)
			return obj, err
		},
	}
	registerTableMetaData(meta)
}

func (pd *ProgData) NewInsertCommandDNSEntrySpec(obj *DNSEntrySpec) ddlog.Command {
	rec := NewRecordDNSEntrySpec(obj)
	tableID := pd.LookupTableID("DNSEntrySpec")
	return ddlog.NewInsertCommand(tableID, rec)
}

func (pd *ProgData) NewInsertOrUpdateCommandDNSEntrySpec(obj *DNSEntrySpec) ddlog.Command {
	rec := NewRecordDNSEntrySpec(obj)
	tableID := pd.LookupTableID("DNSEntrySpec")
	return ddlog.NewInsertOrUpdateCommand(tableID, rec)
}

func (pd *ProgData) NewDeleteValCommandDNSEntrySpec(obj *DNSEntrySpec) ddlog.Command {
	rec := NewRecordDNSEntrySpec(obj)
	tableID := pd.LookupTableID("DNSEntrySpec")
	return ddlog.NewDeleteValCommand(tableID, rec)
}
func (pd *ProgData) NewDeleteKeyCommandDNSEntrySpec(obj *DNSEntrySpec) ddlog.Command {
	rec := NewKeyRecordDNSEntrySpec(obj)
	tableID := pd.LookupTableID("DNSEntrySpec")
	return ddlog.NewDeleteKeyCommand(tableID, rec)
}


func NewKeyRecordDNSEntrySpec(obj *DNSEntrySpec) ddlog.Record {
	return func() ddlog.Record {
	    return NewRecordObjectKey(&obj.Key)
    }()
}

func (pd *ProgData) GetRelTableIDDNSEntrySpec() ddlog.TableID {
	return pd.LookupTableID("DNSEntrySpec")
}
