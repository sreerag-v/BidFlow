package smtp

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"net/smtp"
	"strconv"

	"github.com/sreerag_v/BidFlow/pkg/config"
)


func getRandNum() (string, error) {
	nBig, e := rand.Int(rand.Reader, big.NewInt(8999))
	if e != nil {
		return "", e
	}
	return strconv.FormatInt(nBig.Int64()+1000, 10), nil
}

func sendMail(email string, otp string) {

	fmt.Println("Email : ", email, " otp :", otp)
	// Sender data.
	from := config.GetSmtp().Email
	password := config.GetSmtp().Password

	// Receiver email address.
	to := []string{
		email,
	}

	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Message.

	//  Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, []byte(otp+" is your otp"))
	if err != nil {
		fmt.Println(err)
		return
	}
}

func VerifyOTP(email string) string {
	Otp, err := getRandNum()
	if err != nil {
		panic(err)
	}

	sendMail(email, Otp)
	return Otp
}
