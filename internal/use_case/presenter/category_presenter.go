package presenter

import (
	"github.com/stretchr/testify/mock"
	"gitlab.com/maometusu/qr_menu/internal/entity/models"
	"io"
)

type CategoryPresenter interface {
	PresentCategories(w io.Writer, categories []models.Category, menuID int64) error
}

type CategoryPresenterMock struct {
	mock.Mock
}

func (c *CategoryPresenterMock) PresentCategories(w io.Writer, categories []models.Category, menuID int64) error {
	args := c.Called(w, categories, menuID)
	return args.Error(0)
}

func NewCategoryPresenterMock() *CategoryPresenterMock {
	return &CategoryPresenterMock{}
}
