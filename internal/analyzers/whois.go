package analyzers

import (
	"bytes"
	"io"
	"log"
	"net"
	"strconv"
	"strings"
	//"fmt"
	"regexp"
)

const (
	StopCharacter = "\r\n"
	WhoisPort     = 43
)

// performs a whois query as described in RFC3912 and returns the lines as an array
func Whois(domain string) []string {
	//whois_server := "whois.iana.org" // TODO: Add support for multiple servers.
	lines := WhoisQuery(domain, "whois.iana.org")
	return lines
}

func WhoisQuery(domain, whoisserver string) []string {
	whois_server := whoisserver
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
	lines = ReadWhoisResponse(domain, lines) // recursion!
	return lines
}

func ReadWhoisResponse(domain string, lines []string) []string {
	// check each line for "refer" line
	for i := 0; i < len(lines); i++ {
		// TODO: add more ways to recurse as refer is not standard and most whois servers use something else.
		// BSD WHOIS(1) uses #define WHOIS_SERVER_ID	"Whois Server: "
		// if we find "refer" line, we do another whois query
		matched, _ := regexp.MatchString(`^refer:\s+`, lines[i])
		if matched {
			re := regexp.MustCompile(`^refer:\s+(.*)`)
			whoisserver := re.FindStringSubmatch(lines[i])[1]
			refer_lines := WhoisQuery(domain, whoisserver) 
			lines = append(lines, refer_lines...)
			//fmt.Println(domain)
		}
	}
	return lines

}