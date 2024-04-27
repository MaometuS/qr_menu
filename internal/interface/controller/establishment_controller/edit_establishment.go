package establishment_controller

import (
	"context"
	"fmt"
	"gitlab.com/maometusu/qr_menu/internal/entity/models"
	"net/http"
	"strconv"
)

func (c *controller) EditEstablishment(w http.ResponseWriter, r *http.Request) {
	profileID, ok := r.Context().Value("id").(int64)
	if !ok {
		http.Error(w, "id not exist in context", http.StatusInternalServerError)
		return
	}

	err := r.ParseMultipartForm(0)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	name := r.PostFormValue("name")
	currencyID, err := strconv.ParseInt(r.PostFormValue("currency"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusExpectationFailed)
		return
	}

	id, err := strconv.ParseInt(r.PostFormValue("id"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusExpectationFailed)
		return
	}
	phone := r.PostFormValue("phone")
	logo, _, _ := r.FormFile("logo")
	backgorund, _, _ := r.FormFile("background")
	wifiPassword := r.PostFormValue("wifi_password")
	canMakeOrders := r.PostFormValue("can_make_orders") == "on"
	country := r.PostFormValue("country")
	city := r.PostFormValue("city")
	address := r.PostFormValue("address")
	additionalInfo := r.PostFormValue("additional_info")

	err = c.interactor.EditEstablishment(context.WithValue(r.Context(), "db", c.db), logo, backgorund, &models.Establishment{
		ID:            id,
		Name:          name,
		Phone:         phone,
		WifiPassword:  wifiPassword,
		CanMakeOrders: canMakeOrders,
		Country:       country,
		City:          city,
		Address:       address,
		Text:          additionalInfo,
		CurrencyID:    currencyID,
		ProfileID:     profileID,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/admin/establishment/info?id=%d", id), http.StatusSeeOther)
}
