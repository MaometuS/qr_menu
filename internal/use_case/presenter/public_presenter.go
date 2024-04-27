package presenter

import (
	"github.com/stretchr/testify/mock"
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

type PublicPresenterMock struct {
	mock.Mock
}

func (p *PublicPresenterMock) PresentMainPage(w io.Writer, establishment *models.Establishment, menus []models.Menu, menuID int64, hasSelected bool) error {
	args := p.Called(w, establishment, menus, menuID, hasSelected)
	return args.Error(0)
}

func (p *PublicPresenterMock) PresentCategories(w io.Writer, menu string, categories []models.Category, establishmentID int64, menuID int64) error {
	args := p.Called(w, menu, categories, establishmentID, menuID)
	return args.Error(0)
}

func (p *PublicPresenterMock) PresentItems(w io.Writer, category string, items []models.Item, menuID int64) error {
	args := p.Called(w, category, items, menuID)
	return args.Error(0)
}

func (p *PublicPresenterMock) PresentItem(w io.Writer, item *models.Item, hasSelected bool) error {
	args := p.Called(w, item, hasSelected)
	return args.Error(0)
}

func (p *PublicPresenterMock) PresentSearch(w io.Writer, search string, items []models.Item, establishmentID int64, menuID int64) error {
	args := p.Called(w, search, items, establishmentID, menuID)
	return args.Error(0)
}

func (p *PublicPresenterMock) PresentOrder(w io.Writer, items []models.Item, total float64, menuID int64) error {
	args := p.Called(w, items, total, menuID)
	return args.Error(0)
}

func NewPublicPresenterMock() *PublicPresenterMock {
	return &PublicPresenterMock{}
}
