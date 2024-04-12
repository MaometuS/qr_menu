package category_controller

import (
	"context"
	"fmt"
	"gitlab.com/maometusu/qr_menu/internal/entity/models"
	"net/http"
	"strconv"
)

func (c *categoryController) EditCategory(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("id").(int64)

	err := r.ParseMultipartForm(0)
	if err != nil {
		http.Error(w, err.Error(), http.StatusExpectationFailed)
		return
	}

	categoryID, err := strconv.ParseInt(r.PostFormValue("id"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusExpectationFailed)
		return
	}

	name := r.PostFormValue("name")
	background, _, _ := r.FormFile("background")

	menuID, err := strconv.ParseInt(r.PostFormValue("menu_id"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusExpectationFailed)
		return
	}

	isVisible := r.PostFormValue("is_visible") == "on"

	order, err := strconv.ParseInt(r.PostFormValue("order_param"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusExpectationFailed)
		return
	}

	err = c.interactor.EditCategory(context.WithValue(r.Context(), "db", c.db), id, background, &models.Category{
		ID:         categoryID,
		Name:       name,
		MenuID:     menuID,
		IsVisible:  isVisible,
		OrderParam: order,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/admin/categories?menu_id=%d", menuID), http.StatusSeeOther)
}
