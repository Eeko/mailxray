package analyzers

import (
	"../finding"
)

// PrintableASCIIOnly will implement a function which can check whether a given string is composed only of Printable ASCII characters.
func PrintableASCIIOnly(placeholder string) []finding.Finding {
	var findings []finding.Finding

	findings = append(findings, finding.Finding{
		Message:  "To be implemented.",
		Location: [2]int{0, len(placeholder)},
		Severity: 0,
	})
	return findings
}
