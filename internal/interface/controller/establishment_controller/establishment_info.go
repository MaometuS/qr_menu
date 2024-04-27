package establishment_controller

import (
	"context"
	"net/http"
	"strconv"
)

func (c *controller) EstablishmentInfo(w http.ResponseWriter, r *http.Request) {
	profileID, ok := r.Context().Value("id").(int64)
	if !ok {
		http.Error(w, "id is not in context", http.StatusInternalServerError)
		return
	}

	establishmentID, err := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusExpectationFailed)
		return
	}

	err = c.interactor.EstablishmentInfo(context.WithValue(r.Context(), "db", c.db), w, profileID, establishmentID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
