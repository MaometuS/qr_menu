package admin_controller

import (
	"net/http"
	"strconv"
)

func (c *controller) VerifyRestorePassword(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PostFormValue("id"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusExpectationFailed)
		return
	}

	err = c.interactor.VerifyRestorePassword(r.Context(), w, id, r.PostFormValue("verification_code"))
	if err != nil {
		return
	}

	http.Redirect(w, r, "/login/", http.StatusSeeOther)
}
