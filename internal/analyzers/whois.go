package analyzers

import (
	"bytes"
	"io"
	"log"
	"net"
	"strconv"
	"strings"
)

const (
	StopCharacter = "\r\n"
	WhoisPort     = 43
)

// performs a whois query as described in RFC3912 and returns the lines as an array
func Whois(domain string) []string {
	whois_server := "whois.iana.org" // TODO: Add support for multiple servers.
	whois_address := strings.Join([]string{whois_server, strconv.Itoa(WhoisPort)}, ":")
	conn, err := net.Dial("tcp", whois_address)

	if err != nil {
		log.Println(err)
		return []string{"No WHOIS Records found"}
	}

	defer conn.Close() // closes the connection once we return from the Whois() block
	conn.Write([]byte(domain))
	conn.Write([]byte(StopCharacter))
	// copy response to a buffer
	var buf bytes.Buffer
	var lines []string
	io.Copy(&buf, conn)
	lines = strings.Split(buf.String(), "\n")
	return lines

}
