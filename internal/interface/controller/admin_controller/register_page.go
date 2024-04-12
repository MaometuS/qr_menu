package admin_controller

import (
	"context"
	"net/http"
)

func (c *controller) RegisterPage(w http.ResponseWriter, r *http.Request) {
	err := c.interactor.RegisterPage(context.WithValue(r.Context(), "db", c.db), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
