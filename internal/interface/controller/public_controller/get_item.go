package public_controller

import (
	"context"
	"net/http"
	"strconv"
)

func (p *publicController) GetItem(w http.ResponseWriter, r *http.Request) {
	item, err := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusExpectationFailed)
		return
	}

	selectedItems, err := getSelectedItems(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusExpectationFailed)
		return
	}

	err = p.interactor.GetItem(context.WithValue(r.Context(), "db", p.db), w, item, selectedItems)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
