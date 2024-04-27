package establishment_controller

import (
	"context"
	"fmt"
	"gitlab.com/maometusu/qr_menu/internal/entity/models"
	"net/http"
	"strconv"
)

func (c *controller) CreateEstablishment(w http.ResponseWriter, r *http.Request) {
	id, ok := r.Context().Value("id").(int64)
	if !ok {
		http.Error(w, "no id in context", http.StatusInternalServerError)
		return
	}

	name := r.PostFormValue("name")
	link := r.PostFormValue("link")
	currencyID, err := strconv.ParseInt(r.PostFormValue("currency"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusExpectationFailed)
		return
	}

	err = c.interactor.CreateEstablishment(context.WithValue(r.Context(), "db", c.db), &models.Establishment{
		Name:       name,
		ProfileID:  id,
		CurrencyID: currencyID,
		Link:       link,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/admin/establishments?id=%d", id), http.StatusSeeOther)
}
