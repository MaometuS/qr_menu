package admin_controller

import (
	"net/http"
	"strconv"
)

func (c *controller) VerifyRestorePasswordPage(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusExpectationFailed)
		return
	}

	err = c.interactor.VerifyRestorePasswordPage(r.Context(), w, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusExpectationFailed)
		return
	}
}
