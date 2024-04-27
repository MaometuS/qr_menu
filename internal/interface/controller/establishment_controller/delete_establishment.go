package establishment_controller

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
)

func (c *controller) DeleteEstablishment(w http.ResponseWriter, r *http.Request) {
	id, ok := r.Context().Value("id").(int64)
	if !ok {
		http.Error(w, "id not present in context", http.StatusInternalServerError)
		return
	}

	establishmentID, err := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusExpectationFailed)
		return
	}

	err = c.interactor.DeleteEstablishment(context.WithValue(r.Context(), "db", c.db), id, establishmentID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/admin/establishments?id=%d", id), http.StatusSeeOther)
}
