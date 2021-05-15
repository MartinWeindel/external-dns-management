package generated

import (
	"github.com/vmware/differential-datalog/go/pkg/ddlog"
)

// Generated code for relation DNSOwner
// DO NOT CHANGE MANUALLY

// input relation DNSOwner [DNSOwner]

func init() {
	meta := &TableMetaData{
		TableName: "DNSOwner", 
		RecordName: "DNSOwner",
		Unmarshaller: func(record ddlog.Record) (interface{}, error) {
			obj, err := DNSOwnerFromRecord(record)
			return obj, err
		},
	}
	registerTableMetaData(meta)
}

func (pd *ProgData) NewInsertCommandDNSOwner(obj *DNSOwner) ddlog.Command {
	rec := NewRecordDNSOwner(obj)
	tableID := pd.LookupTableID("DNSOwner")
	return ddlog.NewInsertCommand(tableID, rec)
}

func (pd *ProgData) NewInsertOrUpdateCommandDNSOwner(obj *DNSOwner) ddlog.Command {
	rec := NewRecordDNSOwner(obj)
	tableID := pd.LookupTableID("DNSOwner")
	return ddlog.NewInsertOrUpdateCommand(tableID, rec)
}

func (pd *ProgData) NewDeleteValCommandDNSOwner(obj *DNSOwner) ddlog.Command {
	rec := NewRecordDNSOwner(obj)
	tableID := pd.LookupTableID("DNSOwner")
	return ddlog.NewDeleteValCommand(tableID, rec)
}
func (pd *ProgData) NewDeleteKeyCommandDNSOwner(obj *DNSOwner) ddlog.Command {
	rec := NewKeyRecordDNSOwner(obj)
	tableID := pd.LookupTableID("DNSOwner")
	return ddlog.NewDeleteKeyCommand(tableID, rec)
}


func NewKeyRecordDNSOwner(obj *DNSOwner) ddlog.Record {
	return func() ddlog.Record {
	    return ddlog.NewRecordString(obj.Name)
    }()
}

func (pd *ProgData) GetRelTableIDDNSOwner() ddlog.TableID {
	return pd.LookupTableID("DNSOwner")
}
