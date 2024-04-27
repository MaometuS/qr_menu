package presenter

import (
	"github.com/stretchr/testify/mock"
	"gitlab.com/maometusu/qr_menu/internal/entity/models"
	"io"
)

type MenuPresenter interface {
	PresentMenus(w io.Writer, menus []models.Menu, establishmentID int64) error
}

type MenuPresenterMock struct {
	mock.Mock
}

func (m *MenuPresenterMock) PresentMenus(w io.Writer, menus []models.Menu, establishmentID int64) error {
	args := m.Called(w, menus, establishmentID)
	return args.Error(0)
}

func NewMenuPresenterMock() *MenuPresenterMock {
	return &MenuPresenterMock{}
}
