package main

import (
	"github.com/SegovChik/email-sms/pkg/awsconfig"
	"github.com/SegovChik/email-sms/pkg/email"
	"github.com/SegovChik/email-sms/pkg/sms"
	"log"
)


func main() {
	// Create a session instance.
	ses, err := awsconfig.New(awsconfig.Config{
		Region:  "eu-central-1",
		Profile: "default",
		ID:      "secret",
		Secret:  "secret",
	})
	if err != nil {
		log.Fatalln(err)
	}
	firstName := "John"
	lastName := "Doe"
	data := "{ \"first_name\":\"" + firstName + "\", \"last_name\": \"" + lastName + "\"}"

	email.SendEmail(ses, "SampleTemplate", "segovchik@gmail.com", "seh@sk.ua" , data)

	sms.SendSMS(ses, "text", "+380665065480")
}
