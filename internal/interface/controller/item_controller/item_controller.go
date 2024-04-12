package item_controller

import (
	"gitlab.com/maometusu/qr_menu/internal/entity"
	"gitlab.com/maometusu/qr_menu/internal/use_case/interactor/item_interactor"
	"net/http"
)

type ItemController interface {
	GetItems(w http.ResponseWriter, r *http.Request)
	CreateItem(w http.ResponseWriter, r *http.Request)
	MoveUp(w http.ResponseWriter, r *http.Request)
	MoveDown(w http.ResponseWriter, r *http.Request)
	EditItem(w http.ResponseWriter, r *http.Request)
	DeleteItem(w http.ResponseWriter, r *http.Request)
}

type itemController struct {
	interactor item_interactor.ItemInteractor
	db         entity.PgxIface
}

func NewItemController(
	interactor item_interactor.ItemInteractor,
	db entity.PgxIface,
) ItemController {
	return &itemController{
		interactor: interactor,
		db:         db,
	}
}
