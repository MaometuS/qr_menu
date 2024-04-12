package category_controller

import (
	"gitlab.com/maometusu/qr_menu/internal/entity"
	"gitlab.com/maometusu/qr_menu/internal/use_case/interactor/category_interactor"
	"net/http"
)

type CategoryController interface {
	GetCategories(w http.ResponseWriter, r *http.Request)
	CreateCategory(w http.ResponseWriter, r *http.Request)
	MoveUp(w http.ResponseWriter, r *http.Request)
	MoveDown(w http.ResponseWriter, r *http.Request)
	EditCategory(w http.ResponseWriter, r *http.Request)
	DeleteCategory(w http.ResponseWriter, r *http.Request)
}

type categoryController struct {
	interactor category_interactor.CategoryInteractor
	db         entity.PgxIface
}

func NewCategoryController(
	interactor category_interactor.CategoryInteractor,
	db entity.PgxIface,
) CategoryController {
	return &categoryController{
		interactor: interactor,
		db:         db,
	}
}
