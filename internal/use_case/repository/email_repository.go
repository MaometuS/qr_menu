package repository

type EmailRepository interface {
	SendMail(email, subject, body string) error
}
