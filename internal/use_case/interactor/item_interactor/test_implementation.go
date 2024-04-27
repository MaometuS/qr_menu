package item_interactor

import (
	"context"
	"github.com/stretchr/testify/mock"
	"gitlab.com/maometusu/qr_menu/internal/entity/models"
	"io"
	"mime/multipart"
)

type TestImplementation struct {
	mock.Mock
}

func (t *TestImplementation) GetItems(ctx context.Context, w io.Writer, profileID, categoryID int64) error {
	args := t.Called(ctx, w, profileID, categoryID)
	return args.Error(0)
}

func (t *TestImplementation) CreateItem(ctx context.Context, profileID int64, image multipart.File, item *models.Item) error {
	args := t.Called(ctx, profileID, item)
	return args.Error(0)
}

func (t *TestImplementation) MoveUp(ctx context.Context, profileID, itemID int64) error {
	args := t.Called(ctx, profileID, itemID)
	return args.Error(0)
}

func (t *TestImplementation) MoveDown(ctx context.Context, profileID, itemID int64) error {
	args := t.Called(ctx, profileID, itemID)
	return args.Error(0)
}

func (t *TestImplementation) EditItem(ctx context.Context, profileID int64, image multipart.File, item *models.Item) error {
	args := t.Called(ctx, profileID, item)
	return args.Error(0)
}

func (t *TestImplementation) DeleteItem(ctx context.Context, profileID, itemID int64) error {
	args := t.Called(ctx, profileID, itemID)
	return args.Error(0)
}

func NewTestImplementation() *TestImplementation {
	return &TestImplementation{}
}
