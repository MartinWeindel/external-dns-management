package generated

import (
	"github.com/vmware/differential-datalog/go/pkg/ddlog"
)

// Generated code for relation MatchedEntryToZoneInfo
// DO NOT CHANGE MANUALLY

// output relation MatchedEntryToZoneInfo [MatchedEntryToZoneInfo]

func init() {
	meta := &TableMetaData{
		TableName: "MatchedEntryToZoneInfo", 
		RecordName: "MatchedEntryToZoneInfo",
		Unmarshaller: func(record ddlog.Record) (interface{}, error) {
			obj, err := MatchedEntryToZoneInfoFromRecord(record)
			return obj, err
		},
	}
	registerTableMetaData(meta)
}

func (pd *ProgData) NewInsertCommandMatchedEntryToZoneInfo(obj *MatchedEntryToZoneInfo) ddlog.Command {
	rec := NewRecordMatchedEntryToZoneInfo(obj)
	tableID := pd.LookupTableID("MatchedEntryToZoneInfo")
	return ddlog.NewInsertCommand(tableID, rec)
}

func (pd *ProgData) NewInsertOrUpdateCommandMatchedEntryToZoneInfo(obj *MatchedEntryToZoneInfo) ddlog.Command {
	rec := NewRecordMatchedEntryToZoneInfo(obj)
	tableID := pd.LookupTableID("MatchedEntryToZoneInfo")
	return ddlog.NewInsertOrUpdateCommand(tableID, rec)
}

func (pd *ProgData) NewDeleteValCommandMatchedEntryToZoneInfo(obj *MatchedEntryToZoneInfo) ddlog.Command {
	rec := NewRecordMatchedEntryToZoneInfo(obj)
	tableID := pd.LookupTableID("MatchedEntryToZoneInfo")
	return ddlog.NewDeleteValCommand(tableID, rec)
}

func (pd *ProgData) GetRelTableIDMatchedEntryToZoneInfo() ddlog.TableID {
	return pd.LookupTableID("MatchedEntryToZoneInfo")
}
