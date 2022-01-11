package util

import (
	"fmt"
	"net/smtp"
	"net/textproto"
	"os"

	"github.com/jordan-wright/email"
)

func SendEmail(to, subject, html string) {
	e := &email.Email{
		To:      []string{to},
		From:    fmt.Sprintf("JALBV <%v>", os.Getenv("EMAIL")),
		Subject: subject,
		HTML:    []byte(html),
		Headers: textproto.MIMEHeader{},
	}
	e.Send("smtp.gmail.com:587", smtp.PlainAuth("", os.Getenv("EMAIL"), os.Getenv("EMAIL_PASSWORD"), "smtp.gmail.com"))
}
