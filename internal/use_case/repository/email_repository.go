package repository

import "github.com/stretchr/testify/mock"

type EmailRepository interface {
	SendMail(email, subject, body string) error
}

type EmailRepositoryMock struct {
	mock.Mock
}

func (e *EmailRepositoryMock) SendMail(email, subject, body string) error {
	args := e.Called(email, subject)
	return args.Error(0)
}

func NewEmailRepositoryMock() *EmailRepositoryMock {
	return &EmailRepositoryMock{}
}
