package item_controller

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
)

func (i *itemController) MoveDown(w http.ResponseWriter, r *http.Request) {
	profileID := r.Context().Value("id").(int64)

	id, err := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusExpectationFailed)
	}

	categoryID, err := strconv.ParseInt(r.URL.Query().Get("category_id"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusExpectationFailed)
	}

	err = i.interactor.MoveDown(context.WithValue(r.Context(), "db", i.db), profileID, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusExpectationFailed)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/admin/items?category_id=%d", categoryID), http.StatusSeeOther)
}
