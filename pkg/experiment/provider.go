package experiment

import (
	"github.com/MartinWeindel/ddlog-dnscontroller/go/pkg/generated"
	"github.com/gardener/controller-manager-library/pkg/controllermanager/controller/reconcile"
	"github.com/gardener/controller-manager-library/pkg/logger"

	"github.com/gardener/external-dns-management/pkg/apis/dns/v1alpha1"
	dnsutils "github.com/gardener/external-dns-management/pkg/dns/utils"
)

func refineLogger(logger logger.LogContext, ptype string) logger.LogContext {
	if ptype != "" {
		logger = logger.NewContext("type", ptype)
	}
	return logger
}

func UpdateProvider(logger logger.LogContext, obj *dnsutils.DNSProviderObject) reconcile.Status {
	logger = refineLogger(logger, obj.TypeCode())
	logger.Infof("reconcile PROVIDER")

	spec := obj.Spec()
	if spec.Type == "aws-route53" {
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
				CredentialsHash: "ch1",
				Domains: generated.IncludeExclude{
					Include: domains.Include,
					Exclude: domains.Exclude,
				},
				ProviderType: &generated.AWSRoute53{},
				Zones: generated.IncludeExclude{
					Include: zones.Include,
					Exclude: zones.Exclude,
				},
			},
		}
		logger.Infof("Inserting dnsproviderspec")
		cmdInsert := generated.NewInsertCommandDNSProviderSpec(p1)
		// In practice, each transction would likely include more than one command.
		if err := ddlogProgram.ApplyUpdatesAsTransaction(cmdInsert); err != nil {
			logger.Errorf("Error during transaction: %v", err)
		}
	} else {
		logger.Infof("ignored")
	}
	return reconcile.Succeeded(logger)
}
