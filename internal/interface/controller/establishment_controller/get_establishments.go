package establishment_controller

import (
	"context"
	"net/http"
	"strconv"
)

func (c *controller) GetEstablishments(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusExpectationFailed)
		return
	}

	err = c.interactor.GetEstablishments(context.WithValue(r.Context(), "db", c.db), w, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
