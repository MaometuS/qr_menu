package presenter

import (
	"github.com/stretchr/testify/mock"
	"gitlab.com/maometusu/qr_menu/internal/entity/models"
	"io"
)

type ItemPresenter interface {
	PresentItems(w io.Writer, items []models.Item, categoryID, menuID int64) error
}

type ItemPresenterMock struct {
	mock.Mock
}

func (i *ItemPresenterMock) PresentItems(w io.Writer, items []models.Item, categoryID, menuID int64) error {
	args := i.Called(w, items, categoryID, menuID)
	return args.Error(0)
}

func NewItemPresenterMock() *ItemPresenterMock {
	return &ItemPresenterMock{}
}
