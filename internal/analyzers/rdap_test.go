// tester function for rdap analyzer
package analyzers

import (
	"strings"
	"testing"
)

// Tests for a domain I operate
func TestRdapCyberHorse(t *testing.T) {
	rdapDomain := RdapDomain("cyber.horse")

	t.Run("cyber.horse-domain", func(t *testing.T) {
		wantDomain := "cyber.horse"
		gotDomain := rdapDomain.LDHName
		if wantDomain != gotDomain {
			t.Errorf("Rdap.LDHName \"%s\" does not match wanted value \"%s\"", gotDomain, wantDomain)
		}
	})

	t.Run("cyber.horse-redactions", func(t *testing.T) {
		wantRedactions := 5
		gotRedactions := 0
		for _, entity := range rdapDomain.Entities {
			for _, remark := range entity.Remarks {
				if strings.Contains(remark.Title, "REDACTED") {
					gotRedactions = gotRedactions + 1
				}
			}
		}
		if gotRedactions < wantRedactions {
			t.Errorf("Found only %d redactions. Wanted at least %d", gotRedactions, wantRedactions)
		}
	})

	t.Run("cyber.horse-regdate", func(t *testing.T) {
		wantDate := "2015-03-17T08:55:03Z"
		gotDate := ""
		for _, event := range rdapDomain.Events { // we have to find the registration date from the events array
			switch event.Action {
			case "registration":
				gotDate = event.Date
				// fmt.Printf("Registration Date=%s\n", event.Date)
			}
		}
		if gotDate != wantDate {
			t.Errorf("Date detected \"%s\" does not match wanted date. Wanted %s", gotDate, wantDate)
		}
	})
}
