package admin_controller

import (
	"context"
	"net/http"
)

func (c *controller) EditName(w http.ResponseWriter, r *http.Request) {
	id, ok := r.Context().Value("id").(int64)
	if ok != true {
		http.Error(w, "missing profile id", http.StatusInternalServerError)
		return
	}

	isEmailChanged, err := c.interactor.EditName(context.WithValue(r.Context(), "db", c.db), r.PostFormValue("email"), r.PostFormValue("name"), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if isEmailChanged {
		w.Write([]byte("true"))
		return
	} else {
		w.Write([]byte("false"))
		return
	}
}
