package admin_controller

import (
	"fmt"
	"net/http"
)

func (c *controller) RestorePassword(w http.ResponseWriter, r *http.Request) {
	id, err := c.interactor.RestorePassword(
		r.Context(),
		w,
		r.PostFormValue("email"),
		r.PostFormValue("password"),
		r.PostFormValue("password_repeat"),
	)
	if err != nil {
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/verify_restore_password?id=%d", id), http.StatusSeeOther)
}
