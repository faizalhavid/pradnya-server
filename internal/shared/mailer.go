package shared

import (
	"fmt"
	"net/smtp"
)

type Mailer struct {
	Host     string
	Port     string
	Username string
	Password string
	FromName string
}

func NewMailer(host, port, username, password, fromName string) *Mailer {
	return &Mailer{
		Host:     host,
		Port:     port,
		Username: username,
		Password: password,
		FromName: fromName,
	}
}

func (m *Mailer) SendMail(to, subject, body string) error {
	auth := smtp.PlainAuth("", m.Username, m.Password, m.Host)
	from := fmt.Sprintf("%s <%s>", m.FromName, m.Username)
	msg := []byte(
		"From: " + from + "\r\n" +
			"To: " + to + "\r\n" +
			"Subject: " + subject + "\r\n" +
			"MIME-version: 1.0;\r\n" +
			"Content-Type: text/plain; charset=\"UTF-8\";\r\n\r\n" +
			body + "\r\n",
	)
	addr := fmt.Sprintf("%s:%s", m.Host, m.Port)
	err := smtp.SendMail(addr, auth, m.Username, []string{to}, msg)
	if err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	return nil
}
