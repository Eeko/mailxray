// test cases for email date header tests

package analyzers

import (
	"encoding/json"
	"testing"

	"../finding"
)

func TestDomainLessThanYearOld(t *testing.T) {
	newDomain := "1da0.com" // manually updated variable. New domains can be found from https://dnpedia.com/tlds/daily.php

	want := finding.Finding{
		Message:  "Domain is registered less than 1 year ago.",
		Location: [2]int{0, len(newDomain)},
		Severity: 5,
	}

	testResult := AnalyzeDomain(newDomain)
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

func TestDomainNothingWrong(t *testing.T) {
	newDomain := "google.com" // if google.com is untrusted, we have to rethink our lives and priorities

	testResult := AnalyzeDomain(newDomain)
	matches := false

	if testResult == nil {
		matches = true
	}

	resultJSON, _ := json.Marshal(testResult)
	if matches == false {
		t.Errorf("Got: %v, Wanted the slice to be empty.", string(resultJSON))
	}
}
