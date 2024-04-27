package menu_controller

import (
	"context"
	"net/http"
	"strconv"
)

func (m *menuController) GetMenus(w http.ResponseWriter, r *http.Request) {
	id, ok := r.Context().Value("id").(int64)
	if !ok {
		http.Error(w, "id not present in context", http.StatusInternalServerError)
		return
	}

	establishmentID, err := strconv.ParseInt(r.URL.Query().Get("establishment_id"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusExpectationFailed)
		return
	}

	err = m.interactor.GetMenus(context.WithValue(r.Context(), "db", m.db), w, id, establishmentID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
