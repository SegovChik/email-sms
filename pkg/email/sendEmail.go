package email

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)



func SendEmail(sess *session.Session, template string, sender string, reciver string, data string) {

	svc := ses.New(sess)
	input := &ses.SendTemplatedEmailInput{
		Source:   &sender,
		Template: &template,
		Destination: &ses.Destination{
			ToAddresses: []*string{&reciver},
		},
		TemplateData: &data,
	}
	svc.SendTemplatedEmail(input)
}
