package presenter

import (
	"gitlab.com/maometusu/qr_menu/internal/entity/models"
	"io"
)

type AdminPresenter interface {
	LoginPage(w io.Writer, noMatch, unexpected bool) error
	AdminPage(w io.Writer, profile *models.Profile) error
	ProfilePage(w io.Writer, profile *models.Profile) error
	RegisterPage(w io.Writer, emailExists, passwordsNotMatch, unexpected bool) error
	VerifyPage(w io.Writer, id int64) error
	VerifyEmailEditPage(w io.Writer, email string) error
}
