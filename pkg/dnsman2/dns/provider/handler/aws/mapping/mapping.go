// SPDX-FileCopyrightText: 2024 SAP SE or an SAP affiliate company and Gardener contributors
//
// SPDX-License-Identifier: Apache-2.0
package mapping

import (
	"strings"

	"github.com/gardener/external-dns-management/pkg/dnsman2/dns"
	"github.com/gardener/external-dns-management/pkg/dnsman2/dns/provider/handler/aws/data"
)

var canonicalHostedZones = data.CanonicalHostedZones()

// CanonicalHostedZone returns the matching canonical zone for a given hostname.
func CanonicalHostedZone(hostname string) string {
	for suffix, zone := range canonicalHostedZones {
		if strings.HasSuffix(hostname, suffix) {
			return zone
		}
	}
	return ""
}

// MapTargets maps CNAME records to A/AAAA records for hosted zones used for AWS load balancers.
// Additionally, it adds a TXT record for the CNAME value to query the value via standard DNS queries.
func MapTargets(targets []dns.Target) []dns.Target {
	mapped := make([]dns.Target, 0, len(targets)+1)
	for _, t := range targets {
		switch t.GetRecordType() {
		case dns.TypeCNAME:
			hostedZone := CanonicalHostedZone(t.GetRecordValue())
			if hostedZone != "" {
				switch strings.ToLower(t.GetIPStack()) {
				case dns.AnnotationValueIPStackIPDualStack:
					mapped = append(mapped, dns.NewTarget(dns.TypeAWS_ALIAS_A, t.GetRecordValue(), t.GetTTL()))
					mapped = append(mapped, dns.NewTarget(dns.TypeAWS_ALIAS_AAAA, t.GetRecordValue(), t.GetTTL()))
				case dns.AnnotationValueIPStackIPv6:
					mapped = append(mapped, dns.NewTarget(dns.TypeAWS_ALIAS_AAAA, t.GetRecordValue(), t.GetTTL()))
				default:
					mapped = append(mapped, dns.NewTarget(dns.TypeAWS_ALIAS_A, t.GetRecordValue(), t.GetTTL()))
				}
				mapped = append(mapped, dns.NewTarget(dns.TypeTXT, t.GetRecordValue(), t.GetTTL()))
			} else {
				mapped = append(mapped, t)
			}
		default:
			mapped = append(mapped, t)
		}
	}
	return mapped
}
