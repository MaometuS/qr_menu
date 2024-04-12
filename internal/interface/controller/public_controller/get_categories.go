package public_controller

import (
	"context"
	"net/http"
	"strconv"
)

func (p *publicController) GetCategories(w http.ResponseWriter, r *http.Request) {
	menuID, err := strconv.ParseInt(r.URL.Query().Get("menu_id"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusExpectationFailed)
		return
	}

	err = p.interactor.GetCategories(context.WithValue(r.Context(), "db", p.db), w, menuID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
