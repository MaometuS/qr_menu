package public_controller

import (
	"gitlab.com/maometusu/qr_menu/internal/entity"
	"gitlab.com/maometusu/qr_menu/internal/use_case/interactor/public_interactor"
	"net/http"
)

type PublicController interface {
	GetMainPage(w http.ResponseWriter, r *http.Request)
	GetCategories(w http.ResponseWriter, r *http.Request)
	GetItems(w http.ResponseWriter, r *http.Request)
	GetSearch(w http.ResponseWriter, r *http.Request)
	GetItem(w http.ResponseWriter, r *http.Request)
	AddItem(w http.ResponseWriter, r *http.Request)
	RemoveItem(w http.ResponseWriter, r *http.Request)
	GetOrder(w http.ResponseWriter, r *http.Request)
	AddOrder(w http.ResponseWriter, r *http.Request)
	RemoveOrder(w http.ResponseWriter, r *http.Request)
}

type publicController struct {
	interactor public_interactor.PublicInteractor
	db         entity.PgxIface
}

func NewPublicController(
	interactor public_interactor.PublicInteractor,
	db entity.PgxIface,
) PublicController {
	return &publicController{
		interactor: interactor,
		db:         db,
	}
}
