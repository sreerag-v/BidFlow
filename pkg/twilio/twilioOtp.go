package twilio

import (
	"github.com/sreerag_v/BidFlow/pkg/config"
	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/verify/v2"
)

func TwillioOtpSent(phoneNumber string) (string, error) {
	// create a twilio clint
	user := config.GetTilio().AccountSid
	pass := config.GetTilio().AuthToken
	serviceSid := config.GetTilio().ServiceSid

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: user,
		Password: pass,
	})

	params := &twilioApi.CreateVerificationParams{}
	params.SetTo(phoneNumber)
	params.SetChannel("sms")

	resp, err := client.VerifyV2.CreateVerification(serviceSid, params)

	if err != nil {
		return "", err
	}

	return *resp.Sid, nil
}

func TwilioVerifyOTP(phoneNumber string, code string) error {

	//create a twilio client with twilio details

	user := config.GetTilio().AccountSid
	pass := config.GetTilio().AuthToken
	serviceSid := config.GetTilio().ServiceSid

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Password: pass,
		Username: user,
	})

	params := &twilioApi.CreateVerificationCheckParams{}
	params.SetTo(phoneNumber)
	params.SetCode(code)

	resp, err := client.VerifyV2.CreateVerificationCheck(serviceSid, params)
	if err != nil {
		return err
	} else if *resp.Status == "approved" {
		return nil
	}

	return nil
}
