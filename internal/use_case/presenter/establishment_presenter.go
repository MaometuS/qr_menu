package presenter

import (
	"gitlab.com/maometusu/qr_menu/internal/entity/models"
	"io"
)

type EstablishmentPresenter interface {
	PresentEstablishments(w io.Writer, establishments []models.Establishment, currencies []models.Currency) error
	PresentEstablishmentInfo(w io.Writer, establishment *models.Establishment, currencies []models.Currency) error
	PresentEditEstablishmentPage(w io.Writer, establishment *models.Establishment, topMenuID int64) error
	PresentQrCode(w io.Writer, link string) error
}
