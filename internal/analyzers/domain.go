// running tests for domain names

package analyzers

import (
	"time"

	"../finding"
	"../integrations"
	"github.com/openrdap/rdap"
)

// AnalyzeDomain runs various tests on a domain name. This is the main function which invokes a number of other tests as needed
func AnalyzeDomain(domain string) []finding.Finding {
	var findings []finding.Finding

	rdapDomainTests(domain, &findings)
	return findings
}

// Any tests based on RDAP query
func rdapDomainTests(domain string, findings *[]finding.Finding) {
	rdapResults := integrations.RdapDomain(domain)
	rdapDomainRegistrationDateTests(domain, rdapResults, findings)
	return
}

func rdapDomainRegistrationDateTests(domain string, rdapResults *rdap.Domain, findings *[]finding.Finding) {
	timeNow := time.Now()
	oneYearInPast := timeNow.Add(-time.Hour * 24 * 365)
	for _, event := range rdapResults.Events { // we have to find the registration date from the events array
		switch event.Action {
		case "registration":
			regDate, err := time.Parse(time.RFC3339, event.Date) // rfc3339 timestamp format defined in rfc7483 section 3

			if err != nil {
				*findings = append(*findings, finding.Finding{
					Message:  err.Error(),
					Location: [2]int{0, len(domain)},
					Severity: 7,
				})
			} else {
				if regDate.After(oneYearInPast) {
					*findings = append(*findings, finding.Finding{
						Message:  "Domain is registered less than 1 year ago.",
						Location: [2]int{0, len(domain)},
						Severity: 5,
					})
				}
			}
		}
	}
}
