package public_presenter

import (
	"gitlab.com/maometusu/qr_menu/internal/entity/models"
	"io"
)

func (p *publicPresenter) PresentOrder(w io.Writer, items []models.Item, total float64, menuID int64) error {
	return p.tmpl.ExecuteTemplate(w, "order", map[string]any{
		"Items":       items,
		"Total":       total,
		"MenuID":      menuID,
		"HasSelected": len(items) > 0,
	})
}
