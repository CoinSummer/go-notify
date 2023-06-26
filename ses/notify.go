package ses

import (
	"encoding/json"
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

type Options struct {
	To     string `json:"to"`
	Key    string `json:"key"`
	Secret string `json:"secret"`
	Area   string `json:"host"`
	Sender string `json:"sender"`
}

type Info struct {
	Subject string `json:"subject"`
	Content string `json:"content"`
}

type client struct {
	opt Options
}

func New(opt Options) *client {
	return &client{opt: opt}
}

func (c *client) Send(message string) error {
	if "" == c.opt.To {
		return errors.New("missing email address")
	}

	if "" == message {
		return errors.New("missing message")
	}

	var subject string
	var content string

	var info Info
	err := json.Unmarshal([]byte(message), &info)
	if err == nil {
		subject = info.Subject
		content = info.Content
	} else {
		subject = message
		content = message
	}

	key := c.opt.Key
	secret := c.opt.Secret
	area := c.opt.Area
	to := []*string{
		aws.String(c.opt.To),
	}
	sender := c.opt.Sender
	body := content

	if err := SendToMail(key, secret, area, sender, subject, body, to); err != nil {
		return errors.New("send email error: " + err.Error())
	} else {
		return nil
	}
}

func SendToMail(key string, secret string, area string, sender string, subject string, body string, to []*string) error {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(area),
		Credentials: credentials.NewStaticCredentials(key, secret, ""),
	})

	if err != nil {
		return err
	}

	svc := ses.New(sess)

	input := &ses.SendEmailInput{
		Destination: &ses.Destination{
			CcAddresses: []*string{},
			ToAddresses: to,
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Html: &ses.Content{
					Charset: aws.String("UTF-8"),
					Data:    aws.String(body),
				},
			},
			Subject: &ses.Content{
				Charset: aws.String("UTF-8"),
				Data:    aws.String(subject),
			},
		},
		Source: aws.String(sender),
	}

	_, err_send_email := svc.SendEmail(input)
	if err_send_email != nil {
		return err
	}

	return nil
}
