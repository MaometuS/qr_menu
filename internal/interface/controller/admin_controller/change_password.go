package admin_controller

import (
	"context"
	"net/http"
)

func (c *controller) ChangePassword(w http.ResponseWriter, r *http.Request) {
	id, ok := r.Context().Value("id").(int64)
	if ok != true {
		http.Error(w, "missing profile id", http.StatusInternalServerError)
		return
	}

	password := r.URL.Query().Get("password")
	passRepeat := r.URL.Query().Get("repeat_password")

	err := c.interactor.ChangePassword(context.WithValue(r.Context(), "db", c.db), password, passRepeat, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
