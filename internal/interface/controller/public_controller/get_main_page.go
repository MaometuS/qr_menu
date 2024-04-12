package public_controller

import (
	"context"
	"net/http"
)

func (p *publicController) GetMainPage(w http.ResponseWriter, r *http.Request) {
	selectedItems, err := getSelectedItems(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusExpectationFailed)
		return
	}

	err = p.interactor.GetMainPage(context.WithValue(r.Context(), "db", p.db), w, r.PathValue("link"), selectedItems)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
