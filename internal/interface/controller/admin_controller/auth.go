package admin_controller

import (
	"context"
	"net/http"
	"strings"
)

func (c *controller) Auth(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cook, err := r.Cookie("auth_token")
		if err != nil {
			if r.URL.String() == "/admin/" || strings.Contains(r.URL.String(), "/admin/e") {
				http.Redirect(w, r, "/login/", http.StatusPermanentRedirect)
				return
			}
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		id, err := c.interactor.Auth(context.WithValue(r.Context(), "db", c.db), cook.Value)
		if err != nil {
			if r.URL.String() == "/admin/" || strings.HasPrefix(r.URL.String(), "/admin/e") {
				http.Redirect(w, r, "/login/", http.StatusPermanentRedirect)
				return
			}
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		handler.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), "id", id)))
	})
}
