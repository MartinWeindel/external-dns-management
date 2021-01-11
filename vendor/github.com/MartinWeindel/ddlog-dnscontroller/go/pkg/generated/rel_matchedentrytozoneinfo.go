package generated

import (
	"github.com/vmware/differential-datalog/go/pkg/ddlog"
)

// Generated code for relation MatchedEntryToZoneInfo
// DO NOT CHANGE MANUALLY

// output relation MatchedEntryToZoneInfo [MatchedEntryToZoneInfo]

var (
	relTableIDMatchedEntryToZoneInfo ddlog.TableID = ddlog.GetTableID("MatchedEntryToZoneInfo")
)

func init() {
	relTableIDMatchedEntryToZoneInfo = ddlog.GetTableID("MatchedEntryToZoneInfo")
	meta := &TableMetaData{
		TableID: relTableIDMatchedEntryToZoneInfo,
		TableName: "MatchedEntryToZoneInfo", 
		RecordName: "MatchedEntryToZoneInfo",
		Unmarshaller: func(record ddlog.Record) (interface{}, error) {
			obj, err := MatchedEntryToZoneInfoFromRecord(record)
			return obj, err
		},
	}
	registerTableMetaData(relTableIDMatchedEntryToZoneInfo, meta)
}

func NewInsertCommandMatchedEntryToZoneInfo(obj *MatchedEntryToZoneInfo) ddlog.Command {
	rec := NewRecordMatchedEntryToZoneInfo(obj)
	return ddlog.NewInsertCommand(relTableIDMatchedEntryToZoneInfo, rec)
}

func NewInsertOrUpdateCommandMatchedEntryToZoneInfo(obj *MatchedEntryToZoneInfo) ddlog.Command {
	rec := NewRecordMatchedEntryToZoneInfo(obj)
	return ddlog.NewInsertOrUpdateCommand(relTableIDMatchedEntryToZoneInfo, rec)
}

func NewDeleteValCommandMatchedEntryToZoneInfo(obj *MatchedEntryToZoneInfo) ddlog.Command {
	rec := NewRecordMatchedEntryToZoneInfo(obj)
	return ddlog.NewDeleteValCommand(relTableIDMatchedEntryToZoneInfo, rec)
}

func GetRelTableIDMatchedEntryToZoneInfo() ddlog.TableID {
	return relTableIDMatchedEntryToZoneInfo
}
