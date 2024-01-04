package mailer

import (
	"encoding/json"
	"fmt"
	"log"
	"net/smtp"

	"github.com/sreerag_v/BidFlow/pkg/config"
)

type message struct {
	To      string `json:"to,omitempty"`
	Subject string `json:"subject,omitempty"`
	Body    string `json:"body,omitempty"`
}

func SendMail(jsonBody []byte) {

	fmt.Println(jsonBody)
	var msg message

	if err := json.Unmarshal(jsonBody, &msg); err != nil {
		log.Fatal(err)
	}

	// Sender data.
	from := config.GetSmtp().Email
	password := config.GetSmtp().Password

	// Receiver email address.
	to := []string{
		msg.To,
	}

	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Message.

	//  Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, []byte(msg.Body))
	if err != nil {
		fmt.Println(err)
		return
	}
}
