package public_presenter

import (
	"gitlab.com/maometusu/qr_menu/internal/entity/models"
	"io"
)

func (p *publicPresenter) PresentSearch(w io.Writer, search string, items []models.Item, establishmentID int64, menuID int64) error {
	return p.tmpl.ExecuteTemplate(w, "search", map[string]any{
		"Search":          search,
		"Items":           items,
		"EstablishmentID": establishmentID,
		"MenuID":          menuID,
	})
}
