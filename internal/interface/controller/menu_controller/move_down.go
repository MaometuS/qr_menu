package menu_controller

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
)

func (m *menuController) MoveDown(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("id").(int64)

	menuID, err := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusExpectationFailed)
		return
	}

	establishmentID, err := strconv.ParseInt(r.URL.Query().Get("establishment_id"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusExpectationFailed)
		return
	}

	err = m.interactor.MoveDown(context.WithValue(r.Context(), "db", m.db), id, menuID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/admin/menus?establishment_id=%d", establishmentID), http.StatusSeeOther)
}
