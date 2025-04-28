package email

import (
	"fmt"
	"net/smtp"
	"notification/config"
)

type SMTPService struct {
	config *config.EmailConfig
}

func NewSMTPService(config *config.EmailConfig) *SMTPService {
	return &SMTPService{config: config}
}

func (s *SMTPService) SendMail(to []string, subject string, body string) error {
	auth := smtp.PlainAuth("", s.config.SMTPUser, s.config.SMTPPassword, s.config.SMTPHost)
	message := []byte(fmt.Sprintf("Subject: %s\r\n\r\n%s", subject, body))
	err := smtp.SendMail(
		fmt.Sprintf("%s:%d", s.config.SMTPHost, s.config.SMTPPort),
		auth,
		s.config.SMTPUser,
		to,
		message,
	)
	if err != nil {
		return err
	}

	return nil
}
