package category_interactor

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

func (t *TestImplementation) GetCategories(ctx context.Context, w io.Writer, profileID, menuID int64) error {
	args := t.Called(ctx, w, profileID, menuID)
	return args.Error(0)
}

func (t *TestImplementation) CreateCategory(ctx context.Context, profileID int64, background multipart.File, category *models.Category) error {
	args := t.Called(ctx, profileID, category)
	return args.Error(0)
}

func (t *TestImplementation) MoveUp(ctx context.Context, profileID, categoryID int64) error {
	args := t.Called(ctx, profileID, categoryID)
	return args.Error(0)
}

func (t *TestImplementation) MoveDown(ctx context.Context, profileID, categoryID int64) error {
	args := t.Called(ctx, profileID, categoryID)
	return args.Error(0)
}

func (t *TestImplementation) EditCategory(ctx context.Context, profileID int64, background multipart.File, category *models.Category) error {
	args := t.Called(ctx, profileID, category)
	return args.Error(0)
}

func (t *TestImplementation) DeleteCategory(ctx context.Context, profileID, categoryID int64) error {
	args := t.Called(ctx, profileID, categoryID)
	return args.Error(0)
}

func NewTestImplementation() *TestImplementation {
	return &TestImplementation{}
}
