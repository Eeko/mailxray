package finding

// Defines a generic Finding type
type Finding struct {
	Message string // should tell what has happened
	Location [2]int // the bytes in the content (value) where the error was detected
	Severity int // numeric representation of how suspicious we deem this finding
}
type HeaderFindings struct {
	HeaderName string // tells where the finding was made in headers
	Findings []Finding // Array of findings in a given header
}
type EmailFindings struct {
	HeaderFindings []HeaderFindings
	BodyFindings []Finding
}