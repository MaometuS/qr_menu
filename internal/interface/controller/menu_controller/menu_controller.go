package menu_controller

import (
	"gitlab.com/maometusu/qr_menu/internal/entity"
	"gitlab.com/maometusu/qr_menu/internal/use_case/interactor/menu_interactor"
	"net/http"
)

type MenuController interface {
	GetMenus(w http.ResponseWriter, r *http.Request)
	CreateMenu(w http.ResponseWriter, r *http.Request)
	MoveUp(w http.ResponseWriter, r *http.Request)
	MoveDown(w http.ResponseWriter, r *http.Request)
	EditMenu(w http.ResponseWriter, r *http.Request)
	DeleteMenu(w http.ResponseWriter, r *http.Request)
}

type menuController struct {
	interactor menu_interactor.MenuInteractor
	db         entity.PgxIface
}

func NewMenuController(
	interactor menu_interactor.MenuInteractor,
	db entity.PgxIface,
) MenuController {
	return &menuController{
		interactor, db,
	}
}
