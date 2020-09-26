// tester function for whois analyzer
package analyzers

import (
	"strings"
	"testing"
)

func TestWhoisGivesResults(t *testing.T) {
	want := "Internet Assigned Numbers Authority"
	result := Whois("example.com")
	matches := false
	for _, item := range result {
		if strings.Contains(item, want) {
			matches = true
		}
	}
	if matches == false {
		t.Errorf("Whois response for example.com did not contain %v", want)
	}
}
