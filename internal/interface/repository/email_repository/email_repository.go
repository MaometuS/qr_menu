package email_repository

import (
	"crypto/tls"
	"gitlab.com/maometusu/qr_menu/internal/entity"
	"gitlab.com/maometusu/qr_menu/internal/use_case/repository"
	"gopkg.in/gomail.v2"
)

type emailRepository struct {
	config *entity.Config
}

func (e *emailRepository) SendMail(email, subject, body string) error {
	m := gomail.NewMessage()

	m.SetHeader("From", e.config.MailFrom)
	m.SetHeader("To", email)
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", body)

	d := gomail.NewDialer(e.config.MailAddress, int(e.config.MailPort), e.config.MailFrom, e.config.MailPassword)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}

func NewEmailRepository(config *entity.Config) repository.EmailRepository {
	return &emailRepository{config}
}
