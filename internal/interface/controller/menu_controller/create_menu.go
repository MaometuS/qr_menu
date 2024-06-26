package menu_controller

import (
	"context"
	"fmt"
	"gitlab.com/maometusu/qr_menu/internal/entity/models"
	"net/http"
	"strconv"
)

func (m *menuController) CreateMenu(w http.ResponseWriter, r *http.Request) {
	id, ok := r.Context().Value("id").(int64)
	if !ok {
		http.Error(w, "id not present in context", http.StatusInternalServerError)
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

	err = m.interactor.CreateMenu(context.WithValue(r.Context(), "db", m.db), id, &models.Menu{
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
