package generated

import (
	"github.com/vmware/differential-datalog/go/pkg/ddlog"
)

// Generated code for relation DNSOwner
// DO NOT CHANGE MANUALLY

// input relation DNSOwner [DNSOwner]

var (
	relTableIDDNSOwner ddlog.TableID = ddlog.GetTableID("DNSOwner")
)

func init() {
	relTableIDDNSOwner = ddlog.GetTableID("DNSOwner")
	meta := &TableMetaData{
		TableID: relTableIDDNSOwner,
		TableName: "DNSOwner", 
		RecordName: "DNSOwner",
		Unmarshaller: func(record ddlog.Record) (interface{}, error) {
			obj, err := DNSOwnerFromRecord(record)
			return obj, err
		},
	}
	registerTableMetaData(relTableIDDNSOwner, meta)
}

func NewInsertCommandDNSOwner(obj *DNSOwner) ddlog.Command {
	rec := NewRecordDNSOwner(obj)
	return ddlog.NewInsertCommand(relTableIDDNSOwner, rec)
}

func NewInsertOrUpdateCommandDNSOwner(obj *DNSOwner) ddlog.Command {
	rec := NewRecordDNSOwner(obj)
	return ddlog.NewInsertOrUpdateCommand(relTableIDDNSOwner, rec)
}

func NewDeleteValCommandDNSOwner(obj *DNSOwner) ddlog.Command {
	rec := NewRecordDNSOwner(obj)
	return ddlog.NewDeleteValCommand(relTableIDDNSOwner, rec)
}
func NewDeleteKeyCommandDNSOwner(obj *DNSOwner) ddlog.Command {
	rec := NewKeyRecordDNSOwner(obj)
	return ddlog.NewDeleteKeyCommand(relTableIDDNSOwner, rec)
}


func NewKeyRecordDNSOwner(obj *DNSOwner) ddlog.Record {
	return func() ddlog.Record {
	    return ddlog.NewRecordString(obj.Name)
    }()
}

func GetRelTableIDDNSOwner() ddlog.TableID {
	return relTableIDDNSOwner
}