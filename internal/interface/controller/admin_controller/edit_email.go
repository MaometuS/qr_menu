package admin_controller

import (
	"context"
	"net/http"
)

func (c *controller) EditEmail(w http.ResponseWriter, r *http.Request) {
	id, ok := r.Context().Value("id").(int64)
	if ok != true {
		http.Error(w, "missing profile id", http.StatusInternalServerError)
		return
	}

	err := c.interactor.EditEmail(
		context.WithValue(r.Context(), "db", c.db),
		r.PostFormValue("email"),
		r.PostFormValue("code"),
		id,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/admin/", http.StatusSeeOther)
}
