package public_controller

import (
	"context"
	"net/http"
	"strconv"
)

func (p *publicController) GetItems(w http.ResponseWriter, r *http.Request) {
	categoryID, err := strconv.ParseInt(r.URL.Query().Get("category_id"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusExpectationFailed)
		return
	}

	selectedItems, err := getSelectedItems(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusExpectationFailed)
		return
	}

	err = p.interactor.GetItems(context.WithValue(r.Context(), "db", p.db), w, categoryID, selectedItems)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
