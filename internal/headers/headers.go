package headers

import (
	//"fmt"
	"net/mail"
	"strings"

	//"./date"
	"../finding"
)

// ProcessHeaders will (eventually) send every header it knows how to evaluate
// for their individual processors. Returns the header labels
// of any header that won't get processed.
func ProcessHeaders(headers mail.Header) []finding.HeaderFindings {
	var headerFindings []finding.HeaderFindings
	for header, value := range headers {
		var findings []finding.Finding
		switch strings.ToLower(header) {
		case "date":
			findings = ProcessDate(header, value[0])
			headerFindings = append(headerFindings, finding.HeaderFindings{
				HeaderName: "date",
				Findings:   findings,
			})
			//fmt.Println("Date (processed) :", findings)
		default:
			//fmt.Println(header, "(not processed) :", value)
		}

	}
	return headerFindings
}
