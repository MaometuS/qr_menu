package public_presenter

import (
	"gitlab.com/maometusu/qr_menu/internal/entity/models"
	"io"
)

func (p *publicPresenter) PresentItem(w io.Writer, item *models.Item, hasSelected bool) error {
	return p.tmpl.ExecuteTemplate(w, "item", map[string]any{
		"I":           item,
		"HasSelected": hasSelected,
	})
}
