package tools

import (
	"math/rand"
	"time"
)

// initialize the random generator only once
func init() {
	rand.Seed(time.Now().UnixNano()) // good enough random for our use
}

// RandomDomain generates strings that should pass as a legit domain name. E.g.
// RandomDomain()
// >> 7rv6jd7.fie
func RandomDomain() string {
	var chars = []rune("abcdefghijklmnopqrstuvwxyz0123456789")
	var letters = []rune("abcdefghijklmnopqrstuvwxyz")

	// we create domain names of length 5-12 characters with a .tld of 2-4 chars.
	bodyLen := 5 + rand.Intn(7)
	tldLen := 2 + rand.Intn(2)
	body := make([]rune, bodyLen)
	for i := range body {
		body[i] = chars[rand.Intn(len(chars))]
	}
	tld := make([]rune, tldLen)
	for i := range tld {
		tld[i] = letters[rand.Intn(len(letters))]
	}
	return string(body) + "." + string(tld)
}

func RandomDotComDomain() string {
	var chars = []rune("abcdefghijklmnopqrstuvwxyz0123456789")

	// we create domain names of length 5-12 characters with a .com tld
	bodyLen := 5 + rand.Intn(7)
	body := make([]rune, bodyLen)
	for i := range body {
		body[i] = chars[rand.Intn(len(chars))]
	}
	return string(body) + ".com"
}
