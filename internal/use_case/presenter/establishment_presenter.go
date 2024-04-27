package presenter

import (
	"github.com/stretchr/testify/mock"
	"gitlab.com/maometusu/qr_menu/internal/entity/models"
	"io"
)

type EstablishmentPresenter interface {
	PresentEstablishments(w io.Writer, establishments []models.Establishment, currencies []models.Currency) error
	PresentEstablishmentInfo(w io.Writer, establishment *models.Establishment, currencies []models.Currency) error
	PresentEditEstablishmentPage(w io.Writer, establishment *models.Establishment, topMenuID int64) error
	PresentQrCode(w io.Writer, link string) error
}

type EstablishmentPresenterMock struct {
	mock.Mock
}

func (e *EstablishmentPresenterMock) PresentEstablishments(w io.Writer, establishments []models.Establishment, currencies []models.Currency) error {
	args := e.Called(w, establishments, currencies)
	return args.Error(0)
}

func (e *EstablishmentPresenterMock) PresentEstablishmentInfo(w io.Writer, establishment *models.Establishment, currencies []models.Currency) error {
	args := e.Called(w, establishment, currencies)
	return args.Error(0)
}

func (e *EstablishmentPresenterMock) PresentEditEstablishmentPage(w io.Writer, establishment *models.Establishment, topMenuID int64) error {
	args := e.Called(w, establishment, topMenuID)
	return args.Error(0)
}

func (e *EstablishmentPresenterMock) PresentQrCode(w io.Writer, link string) error {
	args := e.Called(w, link)
	return args.Error(0)
}

func NewEstablishmentPresenterMock() *EstablishmentPresenterMock {
	return &EstablishmentPresenterMock{}
}
