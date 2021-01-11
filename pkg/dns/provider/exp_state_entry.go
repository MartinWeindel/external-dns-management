package provider

import (
	"fmt"
	"net"

	"github.com/MartinWeindel/ddlog-dnscontroller/go/pkg/generated"
	"github.com/gardener/controller-manager-library/pkg/controllermanager/controller/reconcile"
	"github.com/gardener/controller-manager-library/pkg/logger"
	"github.com/gardener/controller-manager-library/pkg/resources"
	"github.com/pkg/errors"

	dnsutils "github.com/gardener/external-dns-management/pkg/dns/utils"
)

func (s *expState) addKnownEntry(name resources.ObjectName) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.knownEntries.Add(name)
}

func (s *expState) removeKnownEntry(name resources.ObjectName) bool {
	s.lock.Lock()
	defer s.lock.Unlock()

	if !s.knownEntries.Contains(name) {
		return false
	}
	s.knownEntries.Remove(name)
	return true
}

func (s *expState) UpdateEntry(logger logger.LogContext, obj *dnsutils.DNSEntryObject) reconcile.Status {
	logger.Infof("reconcile ENTRY")
	defer logger.Infof("end - reconcile ENTRY")

	s.addKnownEntry(obj.ObjectName())
	err := s.context.SetFinalizer(obj)
	if err != nil {
		return reconcile.Delay(logger, fmt.Errorf("cannot set finalizer: %s", err))
	}

	spec := obj.Spec()
	gspec := &generated.DNSEntrySpec{
		Key: generated.ObjectKey{Arg0: obj.GetNamespace(), Arg1: obj.GetName()},
		Spec: generated.EntrySpec{
			Domain: spec.DNSName,
		},
	}
	if spec.OwnerId != nil {
		gspec.Spec.Owner = *spec.OwnerId
	} else {
		gspec.Spec.Owner = s.config.Ident
	}
	if spec.TTL != nil {
		gspec.Spec.Ttl = uint32(*spec.TTL)
	}
	if len(spec.Targets) != 0 {
		aCount := 0
		for _, target := range spec.Targets {
			if net.ParseIP(target) != nil {
				aCount++
			}
		}
		if aCount == 0 {
			gspec.Spec.Rtype = &generated.CNAME{}
		} else if aCount == len(spec.Targets) {
			gspec.Spec.Rtype = &generated.A{}
		}
		gspec.Spec.Records = spec.Targets
	} else if len(spec.Text) != 0 {
		gspec.Spec.Rtype = &generated.TXT{}
		gspec.Spec.Records = spec.Text
	}
	logger.Infof("Inserting dnsentryspec")
	cmd := generated.NewInsertOrUpdateCommandDNSEntrySpec(gspec)
	s.addToDDLogCommandQueue(cmd)
	list := s.nextEntryStatusesFromQueue(obj.ObjectName())
	for _, item := range list {
		logger.Infof("updating status: %s", item.Status.Name())
		status := obj.Status()
		switch v := item.Status.(type) {
		case *generated.Unchanged:
			status.State = "Ready"
			status.Message = nil
		case *generated.Inserting:
			status.State = "Pending"
			msg := "inserting entry"
			status.Message = &msg
		case *generated.Updating:
			status.State = "Pending"
			msg := "updating entry"
			status.Message = &msg
		case *generated.ForeignOwner:
			status.State = "Error"
			msg := fmt.Sprintf("entry has foreign owner %q", v.Owner)
			status.Message = &msg
		case *generated.OwnerConflict:
			status.State = "Error"
			msg := fmt.Sprintf("dnsname is already used by owner %q", v.Owner)
			status.Message = &msg
		case *generated.NoProvider:
			status.State = "Error"
			msg := "no matching provider"
			status.Message = &msg
		}
		if item.State != nil {
			ttl := int64(item.State.Ttl)
			status.TTL = &ttl
			status.Zone = &item.State.Zoneid
			ptype := providerTypeFromGenerated(item.State.Ptype)
			status.ProviderType = &ptype
		}
	}
	zlist := s.nextEntryZonesFromQueue(obj.ObjectName())
	for _, item := range zlist {
		status := obj.Status()
		status.Provider = &item.ProviderKey.Arg1
	}
	if len(list)+len(zlist) > 0 {
		err := obj.UpdateStatus()
		if err != nil {
			return reconcile.Failed(logger, errors.Wrap(err, "UpdateStatus"))
		}
	}

	return reconcile.Succeeded(logger)
}

func (s *expState) EntryDeleted(logger logger.LogContext, key resources.ObjectKey) reconcile.Status {
	return s.deleteEntry(logger, key.ObjectName())
}

func (s *expState) DeleteEntry(logger logger.LogContext, obj *dnsutils.DNSEntryObject) reconcile.Status {
	err := s.context.RemoveFinalizer(obj)
	if err != nil {
		return reconcile.Delay(logger, fmt.Errorf("cannot remove finalizer: %s", err))
	}
	return s.deleteEntry(logger, obj.ObjectName())
}

func (s *expState) deleteEntry(logger logger.LogContext, name resources.ObjectName) reconcile.Status {
	if s.removeKnownEntry(name) {
		spec := &generated.DNSEntrySpec{
			Key: generated.ObjectKey{Arg0: name.Namespace(), Arg1: name.Name()},
		}
		cmd := generated.NewDeleteKeyCommandDNSEntrySpec(spec)
		s.addToDDLogCommandQueue(cmd)
	}
	return reconcile.Succeeded(logger)
}
