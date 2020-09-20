package email

import (
	"net/mail"
	"io"
	"../headers"
	"../finding"
)
// Function for processing email.
func ProcessEmail(message_reader io.Reader) (finding.EmailFindings) {
	var email_findings finding.EmailFindings
	var body_findings = []finding.Finding{
		finding.Finding{
			Message: "TODO: Body findings not yet implemented", 
			Location: [2]int{0, 1}, 
			Severity: 0,
		},
	}
	message, err := mail.ReadMessage(message_reader)
	if err != nil {
		panic(err)
	}
	header_findings := headers.ProcessHeaders(message.Header)
	email_findings.HeaderFindings = header_findings
	email_findings.BodyFindings = body_findings
	return email_findings
}