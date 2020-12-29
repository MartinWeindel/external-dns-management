package provider

import (
	"fmt"

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

func (s *expState) addKnownProvider(name resources.ObjectName) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.knownProviders.Add(name)
}

func (s *expState) removeKnownProvider(name resources.ObjectName) bool {
	s.lock.Lock()
	defer s.lock.Unlock()

	if !s.knownProviders.Contains(name) {
		return false
	}
	s.knownProviders.Remove(name)
	return true
}

func (s *expState) UpdateProvider(logger logger.LogContext, obj *dnsutils.DNSProviderObject) reconcile.Status {
	logger = refineLogger(logger, obj.TypeCode())
	logger.Infof("reconcile PROVIDER")
	defer logger.Infof("end - reconcile PROVIDER")

	s.addKnownProvider(obj.ObjectName())
	err := s.context.SetFinalizer(obj)
	if err != nil {
		return reconcile.Delay(logger, fmt.Errorf("cannot set finalizer: %s", err))
	}

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

		account, err := s.GetDNSAccount(logger, obj, secretProps)
		hash := s.hashAndCache(secretProps)
		var accountStatus generated.Status
		var accountZones []generated.Zone
		if err == nil {
			accountStatus = &generated.Ready{}
			zones, err := account.GetZones()
			if err != nil {
				return reconcile.Failed(logger, errors.Wrap(err, "account.GetZones"))
			}
			for _, zone := range zones {
				azone := generated.Zone{
					Id:               zone.Id(),
					Domain:           zone.Domain(),
					ForwardedDomains: zone.ForwardedDomains(),
				}
				accountZones = append(accountZones, azone)
			}
		} else {
			accountStatus = &generated.Error{}
		}

		accountResult := &generated.AccountResult{
			CredentialsHash: hash,
			ProviderType:    &generated.CloudflareDNS{},
			Status:          accountStatus,
			Zones:           accountZones,
		}
		cmdAccountResult := generated.NewInsertOrUpdateCommandAccountResult(accountResult)
		if err != nil {
			if err2 := s.ddlogProgram.ApplyUpdatesAsTransaction(cmdAccountResult); err2 != nil {
				logger.Warn(errors.Wrap(err, "ApplyUpdatesAsTransaction accountresult err"))
			}
			return reconcile.Failed(logger, errors.Wrap(err, "GetDNSAccount"))
		}

		domains := &v1alpha1.DNSSelection{}
		if spec.Domains != nil {
			domains = spec.Domains
		}
		zones := &v1alpha1.DNSSelection{}
		if spec.Zones != nil {
			zones = spec.Zones
		}
		spec := &generated.DNSProviderSpec{
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
		cmdProviderSpec := generated.NewInsertOrUpdateCommandDNSProviderSpec(spec)
		// In practice, each transction would likely include more than one command.
		if err := s.ddlogProgram.ApplyUpdatesAsTransaction(cmdAccountResult, cmdProviderSpec); err != nil {
			return reconcile.Failed(logger, errors.Wrap(err, "ApplyUpdatesAsTransaction dnsproviderspec"))
		}
		list := s.nextProviderStatusesFromQueue(obj.ObjectName())
		for _, item := range list {
			logger.Infof("updating status: %s", item.Status.Name())
			status := obj.Status()
			status.State = item.Status.Name()
			if item.State != nil {
				status.Domains = v1alpha1.DNSSelectionStatus{
					Included: item.State.Domains.Included,
					Excluded: item.State.Domains.Excluded,
				}
				status.Zones = v1alpha1.DNSSelectionStatus{
					Included: item.State.Zoneids.Included,
					Excluded: item.State.Zoneids.Excluded,
				}
			}
		}
		if len(list) > 0 {
			err := obj.UpdateStatus()
			if err != nil {
				return reconcile.Failed(logger, errors.Wrap(err, "UpdateStatus"))
			}
		}
	} else {
		logger.Infof("ignored")
	}
	return reconcile.Succeeded(logger)
}

func (s *expState) ProviderDeleted(logger logger.LogContext, key resources.ObjectKey) reconcile.Status {
	return s.deleteProvider(logger, key.ObjectName())
}

func (s *expState) DeleteProvider(logger logger.LogContext, obj *dnsutils.DNSProviderObject) reconcile.Status {
	err := s.context.RemoveFinalizer(obj)
	if err != nil {
		return reconcile.Delay(logger, fmt.Errorf("cannot remove finalizer: %s", err))
	}
	return s.deleteProvider(logger, obj.ObjectName())
}

func (s *expState) deleteProvider(logger logger.LogContext, name resources.ObjectName) reconcile.Status {
	if s.removeKnownProvider(name) {
		spec := &generated.DNSProviderSpec{
			Key: generated.ObjectKey{Arg0: name.Namespace(), Arg1: name.Name()},
		}
		cmd := generated.NewDeleteKeyCommandDNSProviderSpec(spec)
		if err := s.ddlogProgram.ApplyUpdatesAsTransaction(cmd); err != nil {
			return reconcile.Failed(logger, errors.Wrap(err, "ApplyUpdatesAsTransaction deleteProvider"))
		}
	}
	return reconcile.Succeeded(logger)
}
