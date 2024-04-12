package public_presenter

import (
	"gitlab.com/maometusu/qr_menu/internal/entity/models"
	"io"
)

func (p *publicPresenter) PresentCategories(w io.Writer, menu string, categories []models.Category, establishmentID int64, menuID int64) error {
	return p.tmpl.ExecuteTemplate(w, "categories", map[string]any{
		"Menu":            menu,
		"Categories":      categories,
		"EstablishmentID": establishmentID,
		"MenuID":          menuID,
	})
}
