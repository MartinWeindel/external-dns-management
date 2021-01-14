/*
 * Copyright 2020 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *       http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 *
 */

package provider

import (
	"encoding/json"
	"sync"
	"sync/atomic"

	"github.com/MartinWeindel/ddlog-dnscontroller/go/pkg/generated"
	"github.com/gardener/controller-manager-library/pkg/logger"
	"github.com/gardener/controller-manager-library/pkg/resources"
	"github.com/gardener/controller-manager-library/pkg/utils"
	"github.com/vmware/differential-datalog/go/pkg/ddlog"
)

type weightedDNSProviderStatus struct {
	item   *generated.DNSProviderStatus
	weight int64
}

type weightedDNSProviderZone struct {
	item   *generated.DNSProviderZone
	weight int64
}

////////////////////////////////////////////////////////////////////////////////
// state handling for ddlog reconcilation
////////////////////////////////////////////////////////////////////////////////

func (s *expState) triggerDDLogUpdate() {
	if s.context.IsReady() {
		s.context.EnqueueCommand(CMD_DDLOG_UPDATE)
	} else {
		s.setup.AddCommand(CMD_DDLOG_UPDATE)
	}
}

func (s *expState) isDDLogUpdateActive() bool {
	return atomic.LoadInt32(&s.ddlogUpdateActive) != 0
}

func (s *expState) setDDLogUpdate(active bool) bool {
	var old, new int32
	if active {
		new = 1
	}
	old = 1 - new
	return atomic.CompareAndSwapInt32(&s.ddlogUpdateActive, old, new)
}

func (s *expState) addToDDLogCommandQueue(cmds ...ddlog.Command) {
	s.lock.Lock()
	defer s.lock.Unlock()
	for _, cmd := range cmds {
		s.ddlogCommandQueue = append(s.ddlogCommandQueue, cmd)
	}
}

func (s *expState) nextDDLogCommandQueue() []ddlog.Command {
	s.lock.Lock()
	defer s.lock.Unlock()
	queue := s.ddlogCommandQueue
	s.ddlogCommandQueue = nil
	return queue
}

func (s *expState) updateAccounts(list []string) {
	s.lock.Lock()
	defer s.lock.Unlock()

	for _, hash := range list {
		s.triggerAccount(hash)
	}
}

func (s *expState) UpdateDDLog(logger logger.LogContext) {
	if !s.setDDLogUpdate(true) {
		return
	}
	defer s.setDDLogUpdate(false)

	if s.outRecordHandler.getOutstandingCount() > 0 {
		logger.Infof("ddlog-update: skipped because of outstanding objects")
		return
	}

	queue := s.nextDDLogCommandQueue()
	if len(queue) == 0 {
		return
	}

	logger.Infof("ddlog-update: transaction with %d commands", len(queue))
	if err := s.ddlogProgram.ApplyUpdatesAsTransaction(queue...); err != nil {
		logger.Errorf("ddlog-update transaction failed: %s", err)
	}

	s.updateAccounts(s.outRecordHandler.nextOutstandingAccounts())

	for _, zoneid := range s.outRecordHandler.nextOutstandingZones() {
		s.triggerHostedZone(zoneid)
	}
	for _, name := range s.outRecordHandler.nextOutstandingProviders() {
		key := resources.NewClusterKey(s.context.GetClusterId(PROVIDER_CLUSTER), providerGroupKind, name.Namespace(), name.Name())
		s.enqueueKey(key)
	}
	for _, name := range s.outRecordHandler.nextOutstandingEntries() {
		key := resources.NewClusterKey(s.context.GetClusterId(TARGET_CLUSTER), entryGroupKind, name.Namespace(), name.Name())
		s.enqueueKey(key)
	}
}

func (s *expState) nextProviderStatusesFromQueue(name resources.ObjectName) []weightedDNSProviderStatus {
	return s.outRecordHandler.nextProviderStatusesFromQueue(name)
}

func (s *expState) nextEntryStatusesFromQueue(name resources.ObjectName) []*generated.DNSEntryStatus {
	return s.outRecordHandler.nextEntryStatusesFromQueue(name)
}

func (s *expState) nextEntryZonesFromQueue(name resources.ObjectName) []*generated.MatchedEntryToZoneInfo {
	return s.outRecordHandler.nextEntryZonesFromQueue(name)
}

func (s *expState) nextRecordSetChange(zoneid string) []*generated.RecordSetChange {
	return s.outRecordHandler.nextRecordSetChange(zoneid)
}

///////////////////////////////////////////////

type outRecordHandler struct {
	lock                 sync.Mutex
	providerUpdateQueue  map[resources.ObjectName][]weightedDNSProviderStatus
	entryUpdateQueue     map[resources.ObjectName][]*generated.DNSEntryStatus
	entryZoneQueue       map[resources.ObjectName][]*generated.MatchedEntryToZoneInfo
	recordChangeQueue    map[string][]*generated.RecordSetChange
	providerZoneChanges  []weightedDNSProviderZone
	outstandingProviders resources.ObjectNameSet
	outstandingEntries   resources.ObjectNameSet
	outstandingZones     utils.StringSet
	outstandingAccounts  utils.StringSet
}

func newOutRecordHandler() outRecordHandler {
	return outRecordHandler{
		providerUpdateQueue:  map[resources.ObjectName][]weightedDNSProviderStatus{},
		entryUpdateQueue:     map[resources.ObjectName][]*generated.DNSEntryStatus{},
		entryZoneQueue:       map[resources.ObjectName][]*generated.MatchedEntryToZoneInfo{},
		recordChangeQueue:    map[string][]*generated.RecordSetChange{},
		outstandingProviders: resources.ObjectNameSet{},
		outstandingEntries:   resources.ObjectNameSet{},
		outstandingZones:     utils.StringSet{},
		outstandingAccounts:  utils.StringSet{},
	}
}

func (h *outRecordHandler) Handle(tableID ddlog.TableID, r ddlog.Record, weight int64) {
	h.lock.Lock()
	defer h.lock.Unlock()

	meta, err := generated.LookupTableMetaData(tableID)
	if err != nil {
		logger.Errorf("ddlog-outrecord: lookup failed tableID=%s weight=%d err=%s dump%=s", tableID, weight, err, r.Dump())
		return
	}
	obj, err := meta.Unmarshaller(r)
	if err != nil {
		logger.Errorf("ddlog-outrecord: unmarshal failed tableID=%s weight=%d err=%s dump%=s", tableID, weight, err, r.Dump())
		return
	}
	if tableID == generated.GetRelTableIDDNSProviderStatus() {
		o := obj.(*generated.DNSProviderStatus)
		name := resources.NewObjectName(o.Key.Arg0, o.Key.Arg1)
		h.outstandingProviders.Add(name)
		h.providerUpdateQueue[name] = append(h.providerUpdateQueue[name], weightedDNSProviderStatus{item: o, weight: weight})
	} else if tableID == generated.GetRelTableIDDNSEntryStatus() && weight == 1 {
		o := obj.(*generated.DNSEntryStatus)
		name := resources.NewObjectName(o.Key.Arg0, o.Key.Arg1)
		h.outstandingEntries.Add(name)
		h.entryUpdateQueue[name] = append(h.entryUpdateQueue[name], o)
	} else if tableID == generated.GetRelTableIDMatchedEntryToZoneInfo() && weight == 1 {
		o := obj.(*generated.MatchedEntryToZoneInfo)
		name := resources.NewObjectName(o.EntryKey.Arg0, o.EntryKey.Arg1)
		h.entryZoneQueue[name] = append(h.entryZoneQueue[name], o)
	} else if tableID == generated.GetRelTableIDDNSProviderZone() {
		o := obj.(*generated.DNSProviderZone)
		h.outstandingZones.Add(o.Zoneid)
		h.providerZoneChanges = append(h.providerZoneChanges, weightedDNSProviderZone{item: o, weight: weight})
	} else if tableID == generated.GetRelTableIDAccountInUse() && weight == 1 {
		o := obj.(*generated.AccountInUse)
		h.outstandingAccounts.Add(o.CredentialsHash)
	} else if tableID == generated.GetRelTableIDRecordSetChange() && weight == 1 {
		o := obj.(*generated.RecordSetChange)
		h.outstandingZones.Add(o.Zoneid)
		h.recordChangeQueue[o.Zoneid] = append(h.recordChangeQueue[o.Zoneid], o)
	}
	s, _ := json.MarshalIndent(obj, "", "  ")
	logger.Infof("ddlog-outrecord: %s[%s] %d: %s\n%s\n", meta.TableName, meta.RecordName, weight, s, r.Dump())
}

func (h *outRecordHandler) nextProviderStatusesFromQueue(name resources.ObjectName) []weightedDNSProviderStatus {
	h.lock.Lock()
	defer h.lock.Unlock()
	list, ok := h.providerUpdateQueue[name]
	if !ok || len(list) == 0 {
		return nil
	}
	delete(h.providerUpdateQueue, name)
	return list
}

func (h *outRecordHandler) nextEntryStatusesFromQueue(name resources.ObjectName) []*generated.DNSEntryStatus {
	h.lock.Lock()
	defer h.lock.Unlock()
	list, ok := h.entryUpdateQueue[name]
	if !ok || len(list) == 0 {
		return nil
	}
	delete(h.entryUpdateQueue, name)
	return list
}

func (h *outRecordHandler) nextEntryZonesFromQueue(name resources.ObjectName) []*generated.MatchedEntryToZoneInfo {
	h.lock.Lock()
	defer h.lock.Unlock()
	list, ok := h.entryZoneQueue[name]
	if !ok || len(list) == 0 {
		return nil
	}
	delete(h.entryZoneQueue, name)
	return list
}

func (h *outRecordHandler) nextProviderZonesChanges() []weightedDNSProviderZone {
	h.lock.Lock()
	defer h.lock.Unlock()
	list := h.providerZoneChanges
	h.providerZoneChanges = nil
	return list
}

func (h *outRecordHandler) nextRecordSetChange(zoneid string) []*generated.RecordSetChange {
	h.lock.Lock()
	defer h.lock.Unlock()
	list, ok := h.recordChangeQueue[zoneid]
	if !ok || len(list) == 0 {
		return nil
	}
	delete(h.recordChangeQueue, zoneid)
	return list
}

func (h *outRecordHandler) getOutstandingCount() int {
	h.lock.Lock()
	defer h.lock.Unlock()
	return len(h.outstandingProviders) + len(h.outstandingEntries) + len(h.outstandingZones) + len(h.outstandingAccounts)
}

func (h *outRecordHandler) nextOutstandingProviders() []resources.ObjectName {
	h.lock.Lock()
	defer h.lock.Unlock()

	list := h.outstandingProviders.AsArray()
	h.outstandingProviders = resources.ObjectNameSet{}

	return list
}

func (h *outRecordHandler) nextOutstandingEntries() []resources.ObjectName {
	h.lock.Lock()
	defer h.lock.Unlock()

	list := h.outstandingEntries.AsArray()
	h.outstandingEntries = resources.ObjectNameSet{}

	return list
}

func (h *outRecordHandler) nextOutstandingZones() []string {
	h.lock.Lock()
	defer h.lock.Unlock()

	list := h.outstandingZones.AsArray()
	h.outstandingZones = utils.StringSet{}

	return list
}

func (h *outRecordHandler) nextOutstandingAccounts() []string {
	h.lock.Lock()
	defer h.lock.Unlock()

	list := h.outstandingAccounts.AsArray()
	h.outstandingAccounts = utils.StringSet{}

	return list
}
