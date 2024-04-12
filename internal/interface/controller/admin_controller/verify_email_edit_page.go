package admin_controller

import (
	"context"
	"net/http"
)

func (c *controller) VerifyEmailEditPage(w http.ResponseWriter, r *http.Request) {
	id, ok := r.Context().Value("id").(int64)
	if ok != true {
		http.Error(w, "missing profile id", http.StatusInternalServerError)
		return
	}

	err := c.interactor.VerifyEmailEditPage(context.WithValue(r.Context(), "db", c.db), w, id, r.URL.Query().Get("email"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
