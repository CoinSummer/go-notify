package email

import (
	"encoding/json"
	"errors"
	"net/smtp"
	"strings"
)

type Options struct {
	ToEmail  string `json:"to_email"`
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
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
	if "" == c.opt.ToEmail {
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

	user := c.opt.User
	password := c.opt.Password
	host := c.opt.Host

	to := []string{c.opt.ToEmail}
	cc := []string{}
	bcc := []string{}

	mailType := "text"
	replyToAddress := c.opt.User

	body := content

	if err := SendToMail(user, password, host, subject, body, mailType, replyToAddress, to, cc, bcc); err != nil {
		return errors.New("send email error: " + err.Error())
	} else {
		return nil
	}
}

func MergeSlice(s1 []string, s2 []string) []string {
	slice := make([]string, len(s1)+len(s2))
	copy(slice, s1)
	copy(slice[len(s1):], s2)
	return slice
}

func SendToMail(user, password, host, subject, body, mailtype, replyToAddress string, to, cc, bcc []string) error {
	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, password, hp[0])
	var content_type string
	if mailtype == "html" {
		content_type = "Content-Type: text/" + mailtype + "; charset=UTF-8"
	} else {
		content_type = "Content-Type: text/plain" + "; charset=UTF-8"
	}

	cc_address := strings.Join(cc, ";")
	bcc_address := strings.Join(bcc, ";")
	to_address := strings.Join(to, ";")
	msg := []byte("To: " + to_address + "\r\nFrom: " + user + "\r\nSubject: " + subject + "\r\nReply-To: " + replyToAddress + "\r\nCc: " + cc_address + "\r\nBcc: " + bcc_address + "\r\n" + content_type + "\r\n\r\n" + body)

	send_to := MergeSlice(to, cc)
	send_to = MergeSlice(send_to, bcc)
	err := smtp.SendMail(host, auth, user, send_to, msg)
	return err
}
