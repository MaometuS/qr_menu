package admin_controller

import "net/http"

func (c *controller) RestorePasswordPage(w http.ResponseWriter, r *http.Request) {
	err := c.interactor.RestorePasswordPage(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
