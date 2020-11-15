// test cases for random domain generator tool

package tools

import (
	"regexp"
	"testing"
)

// Generate 20 domains and check that they match the regexp we have set to match the domain syntax
func Test20DomainsMatchRegexp(t *testing.T) {
	var validDomain = regexp.MustCompile(`^[a-z0-9]{5,12}\.[a-z]{2,4}$`)
	for a := 0; a < 20; a++ {
		randomDomain := RandomDomain()
		matchesRule := validDomain.MatchString(randomDomain)
		if !matchesRule {
			t.Errorf("%v does not match the regexp %v", randomDomain, validDomain.String())
		}
	}
}

// Generate 20 domains and check that they match the regexp we have set to match the domain syntax
func Test20DotComDomainsMatchRegexp(t *testing.T) {
	var validDomain = regexp.MustCompile(`^[a-z0-9]{5,12}\.com$`)
	for a := 0; a < 20; a++ {
		randomDomain := RandomDotComDomain()
		matchesRule := validDomain.MatchString(randomDomain)
		if !matchesRule {
			t.Errorf("%v does not match the regexp %v", randomDomain, validDomain.String())
		}
	}
}
