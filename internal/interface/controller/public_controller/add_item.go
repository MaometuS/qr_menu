package public_controller

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
)

func (p *publicController) AddItem(w http.ResponseWriter, r *http.Request) {
	item, err := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusExpectationFailed)
		return
	}

	selectedItems, err := getSelectedItems(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusExpectationFailed)
		return
	}

	selectedItems, err = p.interactor.AddItem(context.WithValue(r.Context(), "db", p.db), item, selectedItems)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	cookie, err := setSelectedItems(selectedItems)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, cookie)
	http.Redirect(w, r, fmt.Sprintf("/item?id=%d", item), http.StatusSeeOther)
}
