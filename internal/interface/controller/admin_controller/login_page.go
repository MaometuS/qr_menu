package admin_controller

import (
	"context"
	"net/http"
)

func (c *controller) LoginPage(w http.ResponseWriter, r *http.Request) {
	err := c.interactor.LoginPage(context.WithValue(r.Context(), "db", c.db), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
