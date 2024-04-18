package admin_controller

import (
	"context"
	"net/http"
	"time"
)

func (c *controller) HandleLogin(w http.ResponseWriter, r *http.Request) {
	email := r.PostFormValue("email")
	password := r.PostFormValue("password")

	tokenString, err := c.interactor.HandleLogin(context.WithValue(r.Context(), "db", c.db), w, email, password)
	if err != nil {
		//http.Error(w, err.Error(), http.StatusInternalServerError)
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
