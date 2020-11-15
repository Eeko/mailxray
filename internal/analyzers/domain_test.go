// test cases for email date header tests

package analyzers

import (
	"encoding/json"
	"testing"

	"../finding"
	"../tools"
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

func TestDomainDoesNotexist(t *testing.T) {
	newDomain := tools.RandomDomain()
	want := finding.Finding{
		Message:  "No RDAP records found for domain.",
		Location: [2]int{0, len(newDomain)},
		Severity: 8,
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
		t.Errorf("Randomly generated \"domain\" %v should be identified as nonexistent. Got %v, wanted %v",
			string(newDomain),
			string(resultJSON),
			string(wantJSON),
		)
	}
}

func TestDotComDomainDoesNotexist(t *testing.T) {
	newDomain := tools.RandomDotComDomain()
	want := finding.Finding{
		Message:  "No RDAP records found for domain.",
		Location: [2]int{0, len(newDomain)},
		Severity: 8,
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
		t.Errorf("Randomly generated \"domain\" %v should be identified as nonexistent. Got %v, wanted %v",
			string(newDomain),
			string(resultJSON),
			string(wantJSON),
		)
	}
}

// we want to ensure we do not get the error for no domain records from a domain that actually does exist
func TestGoogleComDomainExists(t *testing.T) {
	newDomain := "google.com"
	notWant := finding.Finding{
		Message:  "No RDAP records found for domain.",
		Location: [2]int{0, len(newDomain)},
		Severity: 8,
	}

	testResult := AnalyzeDomain(newDomain)
	matches := true
	for _, item := range testResult {
		if item == notWant {
			matches = false
		}
	}
	resultJSON, _ := json.Marshal(testResult)
	notWantJSON, _ := json.Marshal(notWant)

	if matches == false {
		t.Errorf("Randomly generated \"domain\" %v should be identified as nonexistent. Got %v, which should not contain %v",
			string(newDomain),
			string(resultJSON),
			string(notWantJSON),
		)
	}
}
