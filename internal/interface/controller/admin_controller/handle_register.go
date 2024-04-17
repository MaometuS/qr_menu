package admin_controller

import (
	"context"
	"net/http"
)

func (c *controller) HandleRegister(w http.ResponseWriter, r *http.Request) {
	name := r.PostFormValue("name")
	email := r.PostFormValue("email")
	password := r.PostFormValue("password")
	passRepeat := r.PostFormValue("pass_repeat")

	err := c.interactor.HandleRegister(context.WithValue(r.Context(), "db", c.db), w, name, email, password, passRepeat)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
