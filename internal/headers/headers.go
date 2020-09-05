package headers

import (
	"fmt"
	"net/mail"
)

// ProcessHeaders will (eventually) send every header it knows how to evaluate 
// for their individual processors. Returns the header labels
// of any header that won't get processed.
func ProcessHeaders(headers mail.Header) {
	for header, value := range headers {
		fmt.Println(header, ":", value)
	}
}