package provisioner

import (
	"fmt"
	"net/smtp"
	"os"
)

// Mailer provides an interface for sending emails
type Mailer interface {
	SendInvitationEmail(to string, inviterName string, orgName string, role string) error
}

type SMTPMailer struct {
	host     string
	port     string
	username string
	password string
	from     string
}

func NewSMTPMailer() *SMTPMailer {
	return &SMTPMailer{
		host:     os.Getenv("SMTP_HOST"),
		port:     os.Getenv("SMTP_PORT"),
		username: os.Getenv("SMTP_USER"),
		password: os.Getenv("SMTP_PASS"),
		from:     os.Getenv("SMTP_FROM"),
	}
}

func (m *SMTPMailer) SendInvitationEmail(to string, inviterName string, orgName string, role string) error {
	// If SMTP is not fully configured, just skip or log (good for local dev)
	if m.host == "" || m.port == "" {
		fmt.Printf("SMTP not configured. Would send invitation to %s for org %s\n", to, orgName)
		return nil
	}

	auth := smtp.PlainAuth("", m.username, m.password, m.host)

	subject := fmt.Sprintf("Subject: Invitation to join %s on SupaDash\r\n", orgName)
	body := fmt.Sprintf("You have been invited by %s to join the organization '%s' as a %s.\r\n\r\nLogin to SupaDash to accept.", inviterName, orgName, role)
	
	msg := []byte("To: " + to + "\r\n" + subject + "\r\n" + body)

	addr := fmt.Sprintf("%s:%s", m.host, m.port)
	return smtp.SendMail(addr, auth, m.from, []string{to}, msg)
}
