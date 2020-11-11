package experiment

import (
	"github.com/pkg/errors"

	"github.com/MartinWeindel/ddlog-dnscontroller/go/pkg/generated"
	"github.com/gardener/controller-manager-library/pkg/controllermanager/controller/reconcile"
	"github.com/gardener/controller-manager-library/pkg/logger"
	"github.com/gardener/controller-manager-library/pkg/resources"

	"github.com/gardener/external-dns-management/pkg/apis/dns/v1alpha1"
	dnsutils "github.com/gardener/external-dns-management/pkg/dns/utils"
)

func refineLogger(logger logger.LogContext, ptype string) logger.LogContext {
	if ptype != "" {
		logger = logger.NewContext("type", ptype)
	}
	return logger
}

func (s *State) UpdateProvider(logger logger.LogContext, obj *dnsutils.DNSProviderObject) reconcile.Status {
	logger = refineLogger(logger, obj.TypeCode())
	logger.Infof("reconcile PROVIDER")

	spec := obj.Spec()
	if spec.Type == "cloudflare-dns" {
		ref := spec.SecretRef
		localref := *ref
		ref = &localref
		if ref.Namespace == "" {
			ref.Namespace = obj.GetNamespace()
		}
		secretProps, _, err := s.context.GetSecretPropertiesByRef(obj, ref)
		if err != nil {
			return reconcile.Failed(logger, errors.Wrap(err, "GetSecretPropertiesByRef"))
		}
		hash := s.hashAndCache(secretProps)

		domains := &v1alpha1.DNSSelection{}
		if spec.Domains != nil {
			domains = spec.Domains
		}
		zones := &v1alpha1.DNSSelection{}
		if spec.Zones != nil {
			zones = spec.Zones
		}
		p1 := &generated.DNSProviderSpec{
			Key: generated.ObjectKey{Arg0: obj.GetNamespace(), Arg1: obj.GetName()},
			Spec: generated.ProviderSpec{
				CredentialsHash: hash,
				Domains: generated.IncludeExclude{
					Include: domains.Include,
					Exclude: domains.Exclude,
				},
				ProviderType: &generated.CloudflareDNS{},
				Zones: generated.IncludeExclude{
					Include: zones.Include,
					Exclude: zones.Exclude,
				},
			},
		}
		logger.Infof("Inserting dnsproviderspec")
		cmdInsert := generated.NewInsertOrUpdateCommandDNSProviderSpec(p1)
		// In practice, each transction would likely include more than one command.
		if err := s.ddlogProgram.ApplyUpdatesAsTransaction(cmdInsert); err != nil {
			return reconcile.Failed(logger, errors.Wrap(err, "ApplyUpdatesAsTransaction dnsproviderspec"))
		}
	} else {
		logger.Infof("ignored")
	}
	return reconcile.Succeeded(logger)
}

func (s *State) ProviderDeleted(logger logger.LogContext, key resources.ObjectKey) reconcile.Status {
	return s.deleteProvider(logger, key.ObjectName())
}

func (s *State) RemoveProvider(logger logger.LogContext, obj *dnsutils.DNSProviderObject) reconcile.Status {
	pname := obj.ObjectName()
	return s.deleteProvider(logger, pname)
}

func (s *State) deleteProvider(logger logger.LogContext, name resources.ObjectName) reconcile.Status {
	// TODO delete from ddlog

	return reconcile.Succeeded(logger)
}
