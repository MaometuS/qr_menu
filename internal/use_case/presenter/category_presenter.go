package presenter

import (
	"gitlab.com/maometusu/qr_menu/internal/entity/models"
	"io"
)

type CategoryPresenter interface {
	PresentCategories(w io.Writer, categories []models.Category, menuID int64) error
}
