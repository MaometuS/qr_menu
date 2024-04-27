package item_controller

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
)

func (i *itemController) MoveUp(w http.ResponseWriter, r *http.Request) {
	profileID, ok := r.Context().Value("id").(int64)
	if !ok {
		http.Error(w, "id not present in context", http.StatusInternalServerError)
		return
	}

	id, err := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusExpectationFailed)
		return
	}

	categoryID, err := strconv.ParseInt(r.URL.Query().Get("category_id"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusExpectationFailed)
		return
	}

	err = i.interactor.MoveUp(context.WithValue(r.Context(), "db", i.db), profileID, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/admin/items?category_id=%d", categoryID), http.StatusSeeOther)
}
