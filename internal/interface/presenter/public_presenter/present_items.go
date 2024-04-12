package public_presenter

import (
	"gitlab.com/maometusu/qr_menu/internal/entity/models"
	"io"
)

func (p *publicPresenter) PresentItems(w io.Writer, category string, items []models.Item, menuID int64) error {
	return p.tmpl.ExecuteTemplate(w, "items", map[string]any{
		"Category": category,
		"Items":    items,
		"MenuID":   menuID,
	})
}
