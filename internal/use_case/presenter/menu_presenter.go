package presenter

import (
	"gitlab.com/maometusu/qr_menu/internal/entity/models"
	"io"
)

type MenuPresenter interface {
	PresentMenus(w io.Writer, menus []models.Menu, establishmentID int64) error
}
