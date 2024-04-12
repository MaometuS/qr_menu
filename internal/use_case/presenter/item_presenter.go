package presenter

import (
	"gitlab.com/maometusu/qr_menu/internal/entity/models"
	"io"
)

type ItemPresenter interface {
	PresentItems(w io.Writer, items []models.Item, categoryID, menuID int64) error
}
