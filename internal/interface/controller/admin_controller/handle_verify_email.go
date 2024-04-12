package admin_controller

import (
	"context"
	"net/http"
	"strconv"
	"time"
)

func (c *controller) HandleVerifyEmail(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PostFormValue("id"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusExpectationFailed)
		return
	}

	code := r.PostFormValue("code")

	tokenString, err := c.interactor.HandleVerifyEmail(context.WithValue(r.Context(), "db", c.db), id, code)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	cookie := &http.Cookie{
		Name:     "auth_token",
		Value:    tokenString,
		Path:     "/",
		Expires:  time.Now().AddDate(1, 0, 0),
		Secure:   false,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	}
	http.SetCookie(w, cookie)
	http.Redirect(w, r, "/admin/", http.StatusPermanentRedirect)
}
