package provider

import (
	"github.com/MartinWeindel/ddlog-dnscontroller/go/pkg/generated"
	"github.com/gardener/controller-manager-library/pkg/controllermanager/controller/reconcile"
	"github.com/gardener/controller-manager-library/pkg/logger"
	"github.com/gardener/controller-manager-library/pkg/resources"

	dnsutils "github.com/gardener/external-dns-management/pkg/dns/utils"
)

func (s *expState) addKnownOwner(name resources.ObjectName) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.knownOwners.Add(name)
}

func (s *expState) removeKnownOwner(name resources.ObjectName) bool {
	s.lock.Lock()
	defer s.lock.Unlock()

	if !s.knownOwners.Contains(name) {
		return false
	}
	s.knownOwners.Remove(name)
	return true
}

func (s *expState) UpdateOwner(logger logger.LogContext, obj *dnsutils.DNSOwnerObject) reconcile.Status {
	logger.Infof("reconcile OWNER")
	defer logger.Infof("end - reconcile OWNER")

	s.addKnownOwner(obj.ObjectName())
	spec := obj.Spec()
	active := true
	if spec.Active != nil {
		active = *spec.Active
	}
	owner := &generated.DNSOwner{Name: obj.GetName(), OwnerId: spec.OwnerId, Active: active}
	logger.Infof("Inserting dnsowner")
	cmd := generated.NewInsertOrUpdateCommandDNSOwner(owner)
	s.addToDDLogCommandQueue(cmd)
	return reconcile.Succeeded(logger)
}

func (s *expState) OwnerDeleted(logger logger.LogContext, key resources.ObjectKey) reconcile.Status {
	return s.deleteOwner(logger, key.ObjectName())
}

func (s *expState) deleteOwner(logger logger.LogContext, name resources.ObjectName) reconcile.Status {
	if s.removeKnownOwner(name) {
		pk := &generated.DNSOwner{Name: name.Name()}
		cmd := generated.NewDeleteKeyCommandDNSOwner(pk)
		s.addToDDLogCommandQueue(cmd)

	}
	return reconcile.Succeeded(logger)
}

func (s *expState) UpdateOwnerCounts(log logger.LogContext) {
	// TODO migrate
	/*
		if !this.initialized {
			return
		}
		log.Infof("update owner statistic")
		statistic := statistic.NewEntryStatistic()
		this.UpdateStatistic(statistic)
		types := this.GetHandlerFactory().TypeCodes()
		metrics.UpdateOwnerStatistic(statistic, types)
		changes := this.ownerCache.UpdateCountsWith(statistic.Owners, types)
		if len(changes) > 0 {
			log.Infof("found %d changes for owner usages", len(changes))
			this.ownerupd <- changes
		}
	*/
}
