package repository

import (
	"context"
	"github.com/stretchr/testify/mock"
	"gitlab.com/maometusu/qr_menu/internal/entity/models"
)

type CategoryRepository interface {
	GetAll(ctx context.Context, menuID int64) ([]models.Category, error)
	GetOne(ctx context.Context, categoryID int64) (*models.Category, error)
	CreateCategory(ctx context.Context, category *models.Category) error
	UpdateCategory(ctx context.Context, category *models.Category) error
	DeleteCategory(ctx context.Context, categoryID int64) error
	CheckBelongs(ctx context.Context, categoryID, profileID int64) (bool, error)
	MoveDown(ctx context.Context, categoryID int64) error
	MoveUp(ctx context.Context, categoryID int64) error
}

type CategoryRepositoryMock struct {
	mock.Mock
}

func (c *CategoryRepositoryMock) GetAll(ctx context.Context, menuID int64) ([]models.Category, error) {
	args := c.Called(ctx, menuID)
	return args.Get(0).([]models.Category), args.Error(1)
}

func (c *CategoryRepositoryMock) GetOne(ctx context.Context, categoryID int64) (*models.Category, error) {
	args := c.Called(ctx, categoryID)
	return args.Get(0).(*models.Category), args.Error(1)
}

func (c *CategoryRepositoryMock) CreateCategory(ctx context.Context, category *models.Category) error {
	args := c.Called(ctx, category)
	return args.Error(0)
}

func (c *CategoryRepositoryMock) UpdateCategory(ctx context.Context, category *models.Category) error {
	args := c.Called(ctx, category)
	return args.Error(0)
}

func (c *CategoryRepositoryMock) DeleteCategory(ctx context.Context, categoryID int64) error {
	args := c.Called(ctx, categoryID)
	return args.Error(0)
}

func (c *CategoryRepositoryMock) CheckBelongs(ctx context.Context, categoryID, profileID int64) (bool, error) {
	args := c.Called(ctx, categoryID, profileID)
	return args.Bool(0), args.Error(1)
}

func (c *CategoryRepositoryMock) MoveDown(ctx context.Context, categoryID int64) error {
	args := c.Called(ctx, categoryID)
	return args.Error(0)
}

func (c *CategoryRepositoryMock) MoveUp(ctx context.Context, categoryID int64) error {
	args := c.Called(ctx, categoryID)
	return args.Error(0)
}

func NewCategoryRepositoryMock() *CategoryRepositoryMock {
	return &CategoryRepositoryMock{}
}
