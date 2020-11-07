package analyzers

import (
	"fmt"

	"github.com/openrdap/rdap"
)

// RdapDomain is a  wrapper for openrdap (RFC7482) and returns a domain object.
// E.g. RdapDomain("example.com")
func RdapDomain(domain string) *rdap.Domain {
	req := &rdap.Request{
		Type:  rdap.DomainRequest,
		Query: domain,
	}
	client := &rdap.Client{}
	resp, _ := client.Do(req)
	rdapDomain, ok := resp.Object.(*rdap.Domain)
	if ok {
		fmt.Printf("Handle=%s Domain=%s\n", rdapDomain.Handle, rdapDomain.LDHName)
	}
	return rdapDomain
}
