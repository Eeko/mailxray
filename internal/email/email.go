package email

import (
	"net/mail"
	"io"
	"../headers"
)
// Function for processing email.
func ProcessEmail(message_reader io.Reader) (*mail.Message){
	message, err := mail.ReadMessage(message_reader)
	if err != nil {
		panic(err)
	}
	headers.ProcessHeaders(message.Header)
	return message
}