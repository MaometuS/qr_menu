package establishment_controller

import (
	"context"
	"net/http"
	"strconv"
)

func (c *controller) GetQrCode(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusExpectationFailed)
		return
	}

	w.Header().Set("Content-Type", "image/png")
	err = c.interactor.GetQrCode(context.WithValue(r.Context(), "db", c.db), w, "https://"+r.Host, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
