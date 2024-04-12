package item_controller

import (
	"context"
	"fmt"
	"gitlab.com/maometusu/qr_menu/internal/entity/models"
	"net/http"
	"strconv"
)

func (i *itemController) EditItem(w http.ResponseWriter, r *http.Request) {
	profileID := r.Context().Value("id").(int64)

	err := r.ParseMultipartForm(0)
	if err != nil {
		http.Error(w, err.Error(), http.StatusExpectationFailed)
		return
	}

	id, err := strconv.ParseInt(r.PostFormValue("id"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusExpectationFailed)
		return
	}

	name := r.PostFormValue("name")
	weight := r.PostFormValue("weight")
	price, err := strconv.ParseFloat(r.PostFormValue("price"), 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusExpectationFailed)
		return
	}
	additionalInfo := r.PostFormValue("additional_info")
	isVisible := r.PostFormValue("is_visible") == "on"
	isAvailable := r.PostFormValue("is_available") == "on"

	image, _, _ := r.FormFile("image")

	categoryID, err := strconv.ParseInt(r.PostFormValue("category_id"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusExpectationFailed)
		return
	}

	order, err := strconv.ParseInt(r.PostFormValue("order_param"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusExpectationFailed)
		return
	}

	err = i.interactor.EditItem(context.WithValue(r.Context(), "db", i.db), profileID, image, &models.Item{
		ID:             id,
		Name:           name,
		Weight:         weight,
		Price:          price,
		AdditionalInfo: additionalInfo,
		IsVisible:      isVisible,
		IsAvailable:    isAvailable,
		CategoryID:     categoryID,
		OrderParam:     order,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/admin/items?category_id=%d", categoryID), http.StatusSeeOther)
}
