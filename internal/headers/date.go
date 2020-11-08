package headers

import (
	"net/mail"
	"time"

	"../finding"
)

// ProcessDate contains tests on the date header. Returns an array of findings.
func ProcessDate(header, value string) []finding.Finding {
	currentTime := time.Now()
	var findings []finding.Finding
	parsedDate, err := mail.ParseDate(value)
	if err != nil {
		findings = append(findings,
			finding.Finding{
				Message:  err.Error(),
				Location: [2]int{0, len(value)},
				Severity: 0,
			})
		return findings
	}
	if parsedDate.After(currentTime) {
		findings = append(findings,
			finding.Finding{
				Message:  "Parsed Date is after current date",
				Location: [2]int{0, len(value)},
				Severity: 4,
			})
	}

	if parsedDate.Before(currentTime) {
		findings = append(findings,
			finding.Finding{
				Message:  "Parsed Date is before current date. This is fine and purely informational for testing purposes.",
				Location: [2]int{0, len(value)},
				Severity: 0,
			})
	}

	return findings
}
