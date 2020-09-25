// test cases for email date header tests

package headers

import (
	"testing"
	"time"
	"encoding/json"
	"../finding"
)

// tests whether we get the desired error if the message date is set to future
func TestDateInFuture(t *testing.T) {
	time_now := time.Now()
	one_week_in_future := time_now.Add(time.Hour * 24 * 7)
	formatted_date := one_week_in_future.Format(time.RFC1123Z)
	
	want := finding.Finding{
		Message: "Parsed Date is after current date", 
		Location: [2]int{0, len(formatted_date)}, 
		Severity: 4,
	}

	test_result := ProcessDate("Date", formatted_date)
	matches := false
	for _, item := range test_result {
		if item == want {
			matches = true
		}
	}
	result_json, _ := json.Marshal(test_result)
	want_json, _ := json.Marshal(want)
	if matches == false {
		t.Errorf("Got: %v, Wanted %v", string(result_json), string(want_json) )
	}
}

// tests whether we get the desired error if the message date is set far back enough in the past
func TestDateInDistantPast(t *testing.T) {
	time_now := time.Now()
	ten_years_in_past := time_now.Add(-time.Hour * 24 * 365 * 15)
	formatted_date := ten_years_in_past.Format(time.RFC1123Z)
	
	want := finding.Finding{
		Message: "Parsed Date is over 10 years in the past", 
		Location: [2]int{0, len(formatted_date)}, 
		Severity: 2,
	}

	test_result := ProcessDate("Date", formatted_date)
	matches := false
	for _, item := range test_result {
		if item == want {
			matches = true
		}
	}
	result_json, _ := json.Marshal(test_result)
	want_json, _ := json.Marshal(want)
	if matches == false {
		t.Errorf("Got: %v, Wanted %v", string(result_json), string(want_json) )
	}
}