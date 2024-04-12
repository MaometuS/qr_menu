package menu_controller

import (
	"context"
	"fmt"
	"gitlab.com/maometusu/qr_menu/internal/entity/models"
	"net/http"
	"strconv"
)

func (m *menuController) EditMenu(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("id").(int64)

	menuID, err := strconv.ParseInt(r.PostFormValue("id"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusExpectationFailed)
		return
	}
	name := r.PostFormValue("name")
	isVisible := r.PostFormValue("is_visible") == "on"
	establishmentID, err := strconv.ParseInt(r.PostFormValue("establishment_id"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusExpectationFailed)
		return
	}

	order, err := strconv.ParseInt(r.PostFormValue("order"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusExpectationFailed)
		return
	}

	err = m.interactor.EditMenu(context.WithValue(r.Context(), "db", m.db), id, &models.Menu{
		ID:              menuID,
		Name:            name,
		IsVisible:       isVisible,
		EstablishmentID: establishmentID,
		OrderParam:      order,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/admin/menus?establishment_id=%d", establishmentID), http.StatusSeeOther)
}
