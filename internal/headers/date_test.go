// test cases for email date header tests

package headers

import (
	"encoding/json"
	"testing"
	"time"

	"../finding"
)

// tests whether we get the desired error if the message date is set to future
func TestDateInFuture(t *testing.T) {
	timeNow := time.Now()
	oneWeekInFuture := timeNow.Add(time.Hour * 24 * 7)
	formattedDate := oneWeekInFuture.Format(time.RFC1123Z)

	want := finding.Finding{
		Message:  "Parsed Date is after current date",
		Location: [2]int{0, len(formattedDate)},
		Severity: 4,
	}

	testResult := ProcessDate("Date", formattedDate)
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

// tests whether we get the desired error if the message date is set far back enough in the past
func TestDateInDistantPast(t *testing.T) {
	timeNow := time.Now()
	tenYearsInPast := timeNow.Add(-time.Hour * 24 * 365 * 15)
	formattedDate := tenYearsInPast.Format(time.RFC1123Z)

	want := finding.Finding{
		Message:  "Parsed Date is over 10 years in the past",
		Location: [2]int{0, len(formattedDate)},
		Severity: 2,
	}

	testResult := ProcessDate("Date", formattedDate)
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
