package presenter

import (
	"gitlab.com/maometusu/qr_menu/internal/entity/models"
	"io"
)

type PublicPresenter interface {
	PresentMainPage(w io.Writer, establishment *models.Establishment, menus []models.Menu, menuID int64, hasSelected bool) error
	PresentCategories(w io.Writer, menu string, categories []models.Category, establishmentID int64, menuID int64) error
	PresentItems(w io.Writer, category string, items []models.Item, menuID int64) error
	PresentItem(w io.Writer, item *models.Item, hasSelected bool) error
	PresentSearch(w io.Writer, search string, items []models.Item, establishmentID int64, menuID int64) error
	PresentOrder(w io.Writer, items []models.Item, total float64, menuID int64) error
}
