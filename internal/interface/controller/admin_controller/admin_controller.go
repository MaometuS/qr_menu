package admin_controller

import (
	"gitlab.com/maometusu/qr_menu/internal/entity"
	"gitlab.com/maometusu/qr_menu/internal/use_case/interactor/admin_interactor"
	"net/http"
)

type AdminController interface {
	Auth(handler http.Handler) http.Handler
	AdminPage(w http.ResponseWriter, r *http.Request)
	ProfilePage(w http.ResponseWriter, r *http.Request)
	LoginPage(w http.ResponseWriter, r *http.Request)
	HandleLogin(w http.ResponseWriter, r *http.Request)
	RegisterPage(w http.ResponseWriter, r *http.Request)
	HandleRegister(w http.ResponseWriter, r *http.Request)
	VerifyEmailPage(w http.ResponseWriter, r *http.Request)
	HandleVerifyEmail(w http.ResponseWriter, r *http.Request)
	Logout(w http.ResponseWriter, r *http.Request)
	ChangePassword(w http.ResponseWriter, r *http.Request)
	VerifyEmailEditPage(w http.ResponseWriter, r *http.Request)
	EditEmail(w http.ResponseWriter, r *http.Request)
	EditName(w http.ResponseWriter, r *http.Request)
}

type controller struct {
	db         entity.PgxIface
	interactor admin_interactor.AdminInteractor
}

func NewAdminController(interactor admin_interactor.AdminInteractor, db entity.PgxIface) AdminController {
	return &controller{interactor: interactor, db: db}
}
