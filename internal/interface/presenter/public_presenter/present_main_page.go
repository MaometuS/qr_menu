package public_presenter

import (
	"gitlab.com/maometusu/qr_menu/internal/entity/models"
	"io"
)

func (p *publicPresenter) PresentMainPage(w io.Writer, establishment *models.Establishment, menus []models.Menu, menuID int64, hasSelected bool) error {
	return p.tmpl.ExecuteTemplate(w, "main_page", map[string]any{
		"Establishment": establishment,
		"Menus":         menus,
		"MenuID":        menuID,
		"HasSelected":   hasSelected,
	})
}
