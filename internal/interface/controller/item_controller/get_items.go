package item_controller

import (
	"context"
	"net/http"
	"strconv"
)

func (i *itemController) GetItems(w http.ResponseWriter, r *http.Request) {
	profileID := r.Context().Value("id").(int64)

	categoryID, err := strconv.ParseInt(r.URL.Query().Get("category_id"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusExpectationFailed)
		return
	}

	err = i.interactor.GetItems(context.WithValue(r.Context(), "db", i.db), w, profileID, categoryID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
