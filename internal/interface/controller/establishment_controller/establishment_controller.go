package establishment_controller

import (
	"gitlab.com/maometusu/qr_menu/internal/entity"
	"gitlab.com/maometusu/qr_menu/internal/use_case/interactor/establishment_interactor"
	"net/http"
)

type EstablishmentController interface {
	GetEstablishments(w http.ResponseWriter, r *http.Request)
	CreateEstablishment(w http.ResponseWriter, r *http.Request)
	DeleteEstablishment(w http.ResponseWriter, r *http.Request)
	EditEstablishmentPage(w http.ResponseWriter, r *http.Request)
	EditEstablishment(w http.ResponseWriter, r *http.Request)
	EstablishmentInfo(w http.ResponseWriter, r *http.Request)
	GetQrCode(w http.ResponseWriter, r *http.Request)
	DoesLinkExist(w http.ResponseWriter, r *http.Request)
}

type controller struct {
	interactor establishment_interactor.EstablishmentInteractor
	db         entity.PgxIface
}

func NewEstablishmentController(interactor establishment_interactor.EstablishmentInteractor, db entity.PgxIface) EstablishmentController {
	return &controller{
		interactor: interactor,
		db:         db,
	}
}
