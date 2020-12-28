package provider

import (
	"github.com/pkg/errors"

	"github.com/MartinWeindel/ddlog-dnscontroller/go/pkg/generated"
	"github.com/gardener/controller-manager-library/pkg/controllermanager/controller/reconcile"
	"github.com/gardener/controller-manager-library/pkg/logger"
	"github.com/gardener/controller-manager-library/pkg/resources"

	dnsutils "github.com/gardener/external-dns-management/pkg/dns/utils"
)

func (s *expState) UpdateOwner(logger logger.LogContext, obj *dnsutils.DNSOwnerObject) reconcile.Status {
	logger.Infof("reconcile OWNER")
	defer logger.Infof("end - reconcile OWNER")

	spec := obj.Spec()
	active := true
	if spec.Active != nil {
		active = *spec.Active
	}
	owner := &generated.DNSOwner{Name: obj.GetName(), OwnerId: spec.OwnerId, Active: active}
	logger.Infof("Inserting dnsowner")
	cmd := generated.NewInsertOrUpdateCommandDNSOwner(owner)
	// In practice, each transction would likely include more than one command.
	if err := s.ddlogProgram.ApplyUpdatesAsTransaction(cmd); err != nil {
		return reconcile.Failed(logger, errors.Wrap(err, "ApplyUpdatesAsTransaction dnsowner"))
	}
	return reconcile.Succeeded(logger)
}

func (s *expState) OwnerDeleted(logger logger.LogContext, key resources.ObjectKey) reconcile.Status {
	return s.deleteOwner(logger, key.ObjectName())
}

func (s *expState) deleteOwner(logger logger.LogContext, name resources.ObjectName) reconcile.Status {
	pk := &generated.DNSOwner{Name: name.Name()}
	cmd := generated.NewDeleteKeyCommandDNSOwner(pk)
	if err := s.ddlogProgram.ApplyUpdatesAsTransaction(cmd); err != nil {
		return reconcile.Failed(logger, errors.Wrap(err, "ApplyUpdatesAsTransaction deleteOwner"))
	}
	return reconcile.Succeeded(logger)
}
