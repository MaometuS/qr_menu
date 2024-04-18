package admin_controller

import (
	"context"
	"fmt"
	"net/http"
)

func (c *controller) HandleRegister(w http.ResponseWriter, r *http.Request) {
	name := r.PostFormValue("name")
	email := r.PostFormValue("email")
	password := r.PostFormValue("password")
	passRepeat := r.PostFormValue("pass_repeat")

	id, err := c.interactor.HandleRegister(context.WithValue(r.Context(), "db", c.db), w, name, email, password, passRepeat)
	if err != nil {
		//http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/verify_email?id=%d", id), http.StatusSeeOther)
}
