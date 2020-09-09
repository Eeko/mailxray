package headers

import (
	"fmt"
	"net/mail"
	"strings"
	//"./date"
	"../finding"
)

// ProcessHeaders will (eventually) send every header it knows how to evaluate 
// for their individual processors. Returns the header labels
// of any header that won't get processed.
func ProcessHeaders(headers mail.Header) {
	
	for header, value := range headers {
		var findings []finding.Finding
		switch strings.ToLower(header) {
		case "date":
			findings = ProcessDate(header, value[0])
			fmt.Println("Date (processed) :", findings)
		default:
			fmt.Println(header, "(not processed) :", value)
		}
		
	}
	
}