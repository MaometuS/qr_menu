package menu_interactor

import (
	"context"
	"github.com/stretchr/testify/mock"
	"gitlab.com/maometusu/qr_menu/internal/entity/models"
	"io"
)

type TestImplementation struct {
	mock.Mock
}

func (t *TestImplementation) GetMenus(ctx context.Context, w io.Writer, profileID, establishmentID int64) error {
	args := t.Called(ctx, w, profileID, establishmentID)
	return args.Error(0)
}

func (t *TestImplementation) CreateMenu(ctx context.Context, profileID int64, menu *models.Menu) error {
	args := t.Called(ctx, profileID, menu)
	return args.Error(0)
}

func (t *TestImplementation) MoveUp(ctx context.Context, profileID, menuID int64) error {
	args := t.Called(ctx, profileID, menuID)
	return args.Error(0)
}

func (t *TestImplementation) MoveDown(ctx context.Context, profileID, menuID int64) error {
	args := t.Called(ctx, profileID, menuID)
	return args.Error(0)
}

func (t *TestImplementation) EditMenu(ctx context.Context, profileID int64, menu *models.Menu) error {
	args := t.Called(ctx, profileID, menu)
	return args.Error(0)
}

func (t *TestImplementation) DeleteMenu(ctx context.Context, profileID, menuID int64) error {
	args := t.Called(ctx, profileID, menuID)
	return args.Error(0)
}

func NewTestImplementation() *TestImplementation {
	return &TestImplementation{}
}
