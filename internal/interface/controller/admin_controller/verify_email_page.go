package admin_controller

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
)

func (c *controller) VerifyEmailPage(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusExpectationFailed)
		return
	}

	err = c.interactor.VerifyEmailPage(context.WithValue(r.Context(), "db", c.db), w, id)
	if err != nil {
		http.Redirect(w, r, fmt.Sprintf("/verify_email?id=%d", id), http.StatusSeeOther)
		return
	}
}
