package headers

import (
	"encoding/json"
	"testing"

	"../finding"
	"../tools"
)

// we should get domain alerts based on the from field
func DomainAlertsFromAddress(t *testing.T) {
	fakeDomain := tools.RandomDomain()
	fromField := "From: Hacker <hacker@" + fakeDomain + ">"

	want := finding.Finding{
		Message:  "No RDAP records found for domain.",
		Location: [2]int{21, 21 + len(fakeDomain)},
		Severity: 8,
	}

	testResult := ProcessFrom(fromField)
	matches := false
	for _, item := range testResult {
		if item == want {
			matches = true
		}
	}
	resultJSON, _ := json.Marshal(testResult)
	wantJSON, _ := json.Marshal(want)
	if matches == false {
		t.Errorf("Got: %v, Wanted %v", string(resultJSON), string(wantJSON))
	}
}
