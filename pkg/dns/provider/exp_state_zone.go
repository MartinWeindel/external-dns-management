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

	/*
		delay, hasProviders, req := this.GetZoneReconcilation(logger, zoneid)
		if req == nil || req.zone == nil {
			if !hasProviders {
				return reconcile.Succeeded(logger).Stop()
			}
			return reconcile.Failed(logger, fmt.Errorf("zone %s not used anymore -> stop reconciling", zoneid))
		}
		logger = this.RefineLogger(logger, req.zone.ProviderType())
		if delay > 0 {
			logger.Infof("too early (required delay between two reconcilations: %s) -> skip and reschedule", this.config.Delay)
			return reconcile.Succeeded(logger).RescheduleAfter(delay)
		}
		logger.Infof("precondition fulfilled for zone %s", zoneid)
		if done, err := this.StartZoneReconcilation(logger, req); done {
			if err != nil {
				if _, ok := err.(*perrs.NoSuchHostedZone); ok {
					for _, provider := range req.providers {
						// trigger provider reconciliation to update its status
						_ = this.context.Enqueue(provider.Object())
					}
					return reconcile.Succeeded(logger)
				}
				logger.Infof("zone reconcilation failed for %s: %s", req.zone.Id(), err)
				return reconcile.Succeeded(logger).RescheduleAfter(req.zone.RateLimit())
			}
			return reconcile.Succeeded(logger)
		}
		logger.Infof("reconciling zone %q (%s) already busy and skipped", zoneid, req.zone.Domain())
		return reconcile.Succeeded(logger).RescheduleAfter(10 * time.Second)
	*/
}

/*
func (this *expState) StartZoneReconcilation(logger logger.LogContext, req *zoneReconciliation) (bool, error) {
	xxx
	// TODO
	if req.deleting {
		ctxutil.Tick(this.GetContext().GetContext(), controller.DeletionActivity)
	}
	if req.zone.TestAndSetBusy() {
		defer req.zone.Release()

		list := make(EntryList, 0, len(req.stale)+len(req.entries))
		for _, e := range req.entries {
			list = append(list, e)
		}
		for _, e := range req.stale {
			if req.entries[e.ObjectName()] == nil {
				list = append(list, e)
			} else {
				logger.Errorf("???, duplicate entry in stale and entries")
			}
		}
		logger.Infof("locking %d entries for zone reconcilation", len(list))
		list.Lock()
		defer func() {
			logger.Infof("unlocking %d entries", len(list))
			list.Unlock()
			this.triggerStatistic()
		}()
		return true, this.reconcileZone(logger, req)
	}
	return false, nil
}

func (this *expState) GetZoneReconcilation(logger logger.LogContext, zoneid string) (time.Duration, bool, *expZoneReconciliation) {
	req := &expZoneReconciliation{}

	this.lock.RLock()
	defer this.lock.RUnlock()

	hasProviders := this.hasProviders()
	zone := this.zones[zoneid]
	if zone == nil {
		return 0, hasProviders, nil
	}
	now := time.Now()
	req.zone = zone
	if now.Before(zone.next) {
		return zone.next.Sub(now), hasProviders, req
	}
	req.providers = this.getProvidersForZone(zoneid)
	return 0, hasProviders, req
}
*/

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
