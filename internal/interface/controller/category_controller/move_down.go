package category_controller

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
)

func (c *categoryController) MoveDown(w http.ResponseWriter, r *http.Request) {
	id, ok := r.Context().Value("id").(int64)
	if !ok {
		http.Error(w, "could not get id from context", http.StatusInternalServerError)
		return
	}

	categoryID, err := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusExpectationFailed)
		return
	}

	menuID, err := strconv.ParseInt(r.URL.Query().Get("menu_id"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusExpectationFailed)
		return
	}

	err = c.interactor.MoveDown(context.WithValue(r.Context(), "db", c.db), id, categoryID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/admin/categories?menu_id=%d", menuID), http.StatusSeeOther)
}
