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
	rdapDomainExistsTests(domain, rdapResults, findings)
	rdapDomainRegistrationDateTests(domain, rdapResults, findings)
	return
}

// run tests on domain registration date
func rdapDomainRegistrationDateTests(domain string, rdapResults *rdap.Domain, findings *[]finding.Finding) {
	// handle errors in case registration records are not available. Can happen e.g. when domain does not exist at all.
	defer func() {
		if err := recover(); err != nil {
			*findings = append(*findings, finding.Finding{
				Message:  "Error evaluating RDAP registration records.",
				Location: [2]int{0, len(domain)},
				Severity: 7,
			})
		}
	}()

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

func rdapDomainExistsTests(domain string, rdapResults *rdap.Domain, findings *[]finding.Finding) {
	defer func() {
		if err := recover(); err != nil {
			*findings = append(*findings, finding.Finding{
				Message:  "Error evaluating RDAP records existence.",
				Location: [2]int{0, len(domain)},
				Severity: 7,
			})
		}
	}()
	if rdapResults == nil {
		*findings = append(*findings, finding.Finding{
			Message:  "No RDAP records found for domain.",
			Location: [2]int{0, len(domain)},
			Severity: 8,
		})
	}

}
