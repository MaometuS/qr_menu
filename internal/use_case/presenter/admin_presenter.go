package presenter

import (
	"github.com/stretchr/testify/mock"
	"gitlab.com/maometusu/qr_menu/internal/entity/models"
	"io"
)

type AdminPresenter interface {
	LoginPage(w io.Writer, noMatch, unexpected bool) error
	AdminPage(w io.Writer, profile *models.Profile) error
	ProfilePage(w io.Writer, profile *models.Profile) error
	RegisterPage(w io.Writer, emailExists, passwordsNotMatch, unexpected bool) error
	VerifyPage(w io.Writer, id int64, mismatch, hasError bool) error
	VerifyEmailEditPage(w io.Writer, email string) error
}

type AdminPresenterMock struct {
	mock.Mock
}

func (t *AdminPresenterMock) LoginPage(w io.Writer, noMatch, unexpected bool) error {
	args := t.Called(w, noMatch, unexpected)
	return args.Error(0)
}

func (t *AdminPresenterMock) AdminPage(w io.Writer, profile *models.Profile) error {
	args := t.Called(w, profile)
	return args.Error(0)
}

func (t *AdminPresenterMock) ProfilePage(w io.Writer, profile *models.Profile) error {
	args := t.Called(w, profile)
	return args.Error(0)
}

func (t *AdminPresenterMock) RegisterPage(w io.Writer, emailExists, passwordsNotMatch, unexpected bool) error {
	args := t.Called(w, emailExists, passwordsNotMatch, unexpected)
	return args.Error(0)
}

func (t *AdminPresenterMock) VerifyPage(w io.Writer, id int64, mismatch, hasError bool) error {
	args := t.Called(w, id, mismatch, hasError)
	return args.Error(0)
}

func (t *AdminPresenterMock) VerifyEmailEditPage(w io.Writer, email string) error {
	args := t.Called(w, email)
	return args.Error(0)
}

func NewAdminPresenterMock() *AdminPresenterMock {
	return &AdminPresenterMock{}
}
