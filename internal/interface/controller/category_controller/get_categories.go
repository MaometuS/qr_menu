package category_controller

import (
	"context"
	"net/http"
	"strconv"
)

func (c *categoryController) GetCategories(w http.ResponseWriter, r *http.Request) {
	id, ok := r.Context().Value("id").(int64)
	if !ok {
		http.Error(w, "id is not in context", http.StatusInternalServerError)
		return
	}

	menuID, err := strconv.ParseInt(r.URL.Query().Get("menu_id"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusExpectationFailed)
		return
	}

	err = c.interactor.GetCategories(context.WithValue(r.Context(), "db", c.db), w, id, menuID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
