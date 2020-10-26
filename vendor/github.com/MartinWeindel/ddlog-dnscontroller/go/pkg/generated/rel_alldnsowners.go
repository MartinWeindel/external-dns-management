package generated

import (
	"github.com/vmware/differential-datalog/go/pkg/ddlog"
)

// Generated code for relation AllDNSOwners
// DO NOT CHANGE MANUALLY

// output relation AllDNSOwners [AllDNSOwners]

var (
	relTableIDAllDNSOwners = ddlog.GetTableID("AllDNSOwners")
)

func NewInsertCommandAllDNSOwners(obj *AllDNSOwners) ddlog.Command {
	rec := NewRecordAllDNSOwners(obj)
	return ddlog.NewInsertCommand(relTableIDAllDNSOwners, rec)
}
