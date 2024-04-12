package establishment_controller

import (
	"context"
	"net/http"
)

func (c *controller) EditEstablishmentPage(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("id").(int64)
	link := r.PathValue("establishment")

	err := c.interactor.EditEstablishmentPage(context.WithValue(r.Context(), "db", c.db), w, id, link)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
