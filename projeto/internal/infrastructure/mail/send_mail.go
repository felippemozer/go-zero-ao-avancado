package mail

import (
	"emailn/internal/domain/campaign"
	"os"

	"gopkg.in/gomail.v2"
)

func SendMail(c *campaign.Campaign) error {
	d := gomail.NewDialer(os.Getenv("GMAIL_SMTP"), 587, os.Getenv("EMAIL_USER"), os.Getenv("EMAIL_PASSWORD"))

	var emails []string
	for _, v := range c.Contacts {
		emails = append(emails, v.Email)
	}
	m := gomail.NewMessage()
	m.SetHeader("From", os.Getenv("EMAIL_USER"))
	m.SetHeader("To", emails...)
	m.SetHeader("Subject", c.Name)
	m.SetBody("text/plain", c.Content)

	return d.DialAndSend(m)
}
