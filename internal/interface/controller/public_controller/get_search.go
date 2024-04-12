package public_controller

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
)

func (p *publicController) GetSearch(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	search := query.Get("search")
	establishmentID, err := strconv.ParseInt(query.Get("establishment"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusExpectationFailed)
		return
	}

	menuID, err := strconv.ParseInt(query.Get("menu"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusExpectationFailed)
		return
	}

	if search == "" {
		http.Redirect(w, r, fmt.Sprintf("/categories?menu_id=%d", menuID), http.StatusSeeOther)
		return
	}

	selectedItems, err := getSelectedItems(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusExpectationFailed)
		return
	}

	err = p.interactor.GetSearch(context.WithValue(r.Context(), "db", p.db), w, search, establishmentID, menuID, selectedItems)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
