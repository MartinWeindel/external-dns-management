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
	"fmt"
	"strings"

	"github.com/MartinWeindel/ddlog-dnscontroller/go/pkg/generated"
	"github.com/gardener/controller-manager-library/pkg/controllermanager/controller/reconcile"
	"github.com/gardener/controller-manager-library/pkg/logger"
	"github.com/gardener/external-dns-management/pkg/dns"
)

////////////////////////////////////////////////////////////////////////////////
// state handling for zone reconcilation
////////////////////////////////////////////////////////////////////////////////

type expZoneReconciliation struct {
	zone      *dnsHostedZone
	providers DNSProviders
	//	entries   Entries
	//	stale     DNSNames
	//	dedicated bool
	//	deleting  bool
}

func (s *expState) triggerHostedZone(name string) {
	cmd := CMD_HOSTEDZONE_PREFIX + name
	if s.context.IsReady() {
		s.context.EnqueueCommand(cmd)
	} else {
		s.setup.AddCommand(cmd)
	}
}

func (s *expState) triggerAccount(hash string) {
	cmd := CMD_ACCOUNT_PREFIX + hash
	if s.context.IsReady() {
		s.context.EnqueueCommand(cmd)
	} else {
		s.setup.AddCommand(cmd)
	}
}

func (s *expState) DecodeZoneCommand(cmd string) string {
	if strings.HasPrefix(cmd, CMD_HOSTEDZONE_PREFIX) {
		return cmd[len(CMD_HOSTEDZONE_PREFIX):]
	}
	return ""
}

func (s *expState) DecodeAccountCommand(cmd string) string {
	if strings.HasPrefix(cmd, CMD_ACCOUNT_PREFIX) {
		return cmd[len(CMD_ACCOUNT_PREFIX):]
	}
	return ""
}

func (s *expState) ReconcileAccount(logger logger.LogContext, hash string) reconcile.Status {
	logger.Infof("reconcile ACCOUNT")
	defer logger.Infof("ACCOUNT done")
	s.lock.Lock()
	defer s.lock.Unlock()

	account := s.accountCache.GetByHash(hash)
	if account == nil {
		logger.Info("account not found")
		return reconcile.Succeeded(logger)
	}
	zones, err := account.GetZones()
	if err != nil {
		logger.Warnf("GetZones failed with %s", err)
		return reconcile.Succeeded(logger)
	}
	for _, z := range zones {
		zone := s.zones[z.Id()]
		if zone == nil {
			zone = newDNSHostedZone(s.config.RescheduleDelay, z)
			s.zones[z.Id()] = zone
			logger.Infof("adding hosted zone %q (%s)", z.Id(), z.Domain())
			s.zonesAccount[z.Id()] = hash
			s.triggerHostedZone(zone.Id())
		}
	}
	return reconcile.Succeeded(logger)
}

func (s *expState) lookupAccountAndZone(zoneid string) (*DNSAccount, *dnsHostedZone, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	hash := s.zonesAccount[zoneid]
	if hash == "" {
		return nil, nil, fmt.Errorf("account hash not found")
	}
	account := s.accountCache.GetByHash(hash)
	if account == nil {
		return nil, nil, fmt.Errorf("account not found for hash %s", hash)
	}
	zone := s.zones[zoneid]
	if zone == nil {
		return nil, nil, fmt.Errorf("zone %s not found", zoneid)
	}
	return account, zone, nil
}

func (s *expState) ReconcileZone(logger logger.LogContext, zoneid string) reconcile.Status {
	logger.Infof("reconcile ZONE")
	defer logger.Infof("zone done")

	account, zone, err := s.lookupAccountAndZone(zoneid)
	if err != nil {
		logger.Warn(err.Error())
		return reconcile.Succeeded(logger)
	}

	state, err := account.GetZoneState(zone.zone)
	if err != nil {
		logger.Warnf("GetZoneState failed: %s", err)
		return reconcile.Succeeded(logger)
	}

	var done DoneHandler
	for _, change := range s.nextRecordSetChange(zoneid) {
		var reqs []*ChangeRequest
		for _, data := range change.Insert {
			reqs = addChangeRequests(reqs, true, nil, &data, done)
		}
		for i := 0; i < len(change.Update); i += 2 {
			oldData := change.Update[i]
			newData := change.Update[i+1]
			if oldData.Domain != newData.Domain {
				logger.Errorf("update domain mismatch: %s != %s", oldData.Domain, newData.Domain)
				continue
			}
			changedOwner := oldData.Owner != newData.Owner
			reqs = addChangeRequests(reqs, changedOwner, &oldData, &newData, done)
		}
		for _, data := range change.Delete {
			reqs = addChangeRequests(reqs, true, &data, nil, done)
		}
		err := account.ExecuteRequests(logger, zone.zone, state, reqs)
		if err != nil {
			logger.Warn(err.Error())
		}
	}

	state, err = account.GetZoneState(zone.zone)
	if err != nil {
		logger.Warnf("GetZoneState failed: %s", err)
		return reconcile.Succeeded(logger)
	}

	ptype := providerTypeToGenerated(account.ProviderType())

	{
		rs := &generated.RecordSet{
			Owner:   "",
			Ptype:   ptype,
			Zoneid:  zoneid,
			Domain:  zone.Domain(),
			Rtype:   &generated.ZONEROOT{},
			Ttl:     uint32(0),
			Records: []string{},
		}
		cmd := generated.NewInsertOrUpdateCommandRecordSet(rs)
		s.addToDDLogCommandQueue(cmd)
	}

	for dnsName, set := range state.GetDNSSets() {
		owner := set.GetOwner()
		for recordType, recordset := range set.Sets {
			rtype := recordTypeToGenerated(recordType)
			if rtype != nil {
				records := make([]string, len(recordset.Records))
				for i, record := range recordset.Records {
					records[i] = record.Value
				}
				rs := &generated.RecordSet{
					Owner:   owner,
					Ptype:   ptype,
					Zoneid:  zoneid,
					Domain:  dnsName,
					Rtype:   rtype,
					Ttl:     uint32(recordset.TTL),
					Records: records,
				}
				cmd := generated.NewInsertOrUpdateCommandRecordSet(rs)
				s.addToDDLogCommandQueue(cmd)
			}
		}
	}

	return reconcile.Succeeded(logger)
}

func addChangeRequests(reqs []*ChangeRequest, addMeta bool, old, new *generated.RecordSetData,
	done DoneHandler) []*ChangeRequest {
	var oldset, newset, oldmetaset, newmetaset *dns.DNSSet
	var rtype string
	if new != nil {
		rtype, newset, newmetaset = buildDNSSets(*new, addMeta)
	}
	if old != nil {
		rtype, oldset, oldmetaset = buildDNSSets(*old, addMeta)
	}

	var action string
	if old == nil && new != nil {
		action = R_CREATE
	} else if old != nil && new != nil {
		action = R_UPDATE
	} else if old != nil && new == nil {
		action = R_DELETE
	}
	if action != "" {
		reqs = append(reqs, NewChangeRequest(action, rtype, oldset, newset, done))
		if addMeta {
			reqs = append(reqs, NewChangeRequest(action, dns.RS_META, oldmetaset, newmetaset, done))
		}
	}
	return reqs
}

func buildDNSSets(data generated.RecordSetData, addMeta bool) (rtype string, set *dns.DNSSet, meta *dns.DNSSet) {
	rtype = recordTypeFromGenerated(data.Rtype)
	set = dns.NewDNSSet(data.Domain)
	set.SetRecordSet(rtype, int64(data.Ttl), data.Records...)
	if addMeta {
		meta = dns.NewDNSSet(data.Domain)
		meta.SetRecordSet(dns.RS_META, 600)
		meta.SetOwner(data.Owner)
	}
	return
}

func recordTypeToGenerated(typename string) generated.RecordType {
	switch typename {
	case dns.RS_A:
		return &generated.A{}
	case dns.RS_TXT:
		return &generated.TXT{}
	case dns.RS_CNAME:
		return &generated.CNAME{}
	}
	return nil
}

func recordTypeFromGenerated(rtype generated.RecordType) string {
	switch rtype.(type) {
	case *generated.A:
		return dns.RS_A
	case *generated.TXT:
		return dns.RS_TXT
	case *generated.CNAME:
		return dns.RS_CNAME
	}
	return ""
}
