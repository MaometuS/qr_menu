package repository

import (
	"context"
	"github.com/stretchr/testify/mock"
	"gitlab.com/maometusu/qr_menu/internal/entity/models"
)

type ItemRepository interface {
	GetAll(ctx context.Context, categoryID int64) ([]models.Item, error)
	GetOne(ctx context.Context, itemID int64) (*models.Item, error)
	Search(ctx context.Context, search string, establishmentID int64) ([]models.Item, error)
	CreateItem(ctx context.Context, item *models.Item) error
	UpdateItem(ctx context.Context, item *models.Item) error
	DeleteItem(ctx context.Context, itemID int64) error
	CheckBelongs(ctx context.Context, itemID, profileID int64) (bool, error)
	MoveDown(ctx context.Context, itemID int64) error
	MoveUp(ctx context.Context, itemID int64) error
	GetEstablishment(ctx context.Context, itemID int64) (int64, error)
	GetMultiple(ctx context.Context, ids []int64) ([]models.Item, error)
}

type ItemRepositoryMock struct {
	mock.Mock
}

func (i *ItemRepositoryMock) GetAll(ctx context.Context, categoryID int64) ([]models.Item, error) {
	args := i.Called(ctx, categoryID)
	return args.Get(0).([]models.Item), args.Error(1)
}

func (i *ItemRepositoryMock) GetOne(ctx context.Context, itemID int64) (*models.Item, error) {
	args := i.Called(ctx, itemID)
	return args.Get(0).(*models.Item), args.Error(1)
}

func (i *ItemRepositoryMock) Search(ctx context.Context, search string, establishmentID int64) ([]models.Item, error) {
	args := i.Called(ctx, search, establishmentID)
	return args.Get(0).([]models.Item), args.Error(1)
}

func (i *ItemRepositoryMock) CreateItem(ctx context.Context, item *models.Item) error {
	args := i.Called(ctx, item)
	return args.Error(0)
}

func (i *ItemRepositoryMock) UpdateItem(ctx context.Context, item *models.Item) error {
	args := i.Called(ctx, item)
	return args.Error(0)
}

func (i *ItemRepositoryMock) DeleteItem(ctx context.Context, itemID int64) error {
	args := i.Called(ctx, itemID)
	return args.Error(0)
}

func (i *ItemRepositoryMock) CheckBelongs(ctx context.Context, itemID, profileID int64) (bool, error) {
	args := i.Called(ctx, itemID, profileID)
	return args.Bool(0), args.Error(1)
}

func (i *ItemRepositoryMock) MoveDown(ctx context.Context, itemID int64) error {
	args := i.Called(ctx, itemID)
	return args.Error(0)
}

func (i *ItemRepositoryMock) MoveUp(ctx context.Context, itemID int64) error {
	args := i.Called(ctx, itemID)
	return args.Error(0)
}

func (i *ItemRepositoryMock) GetEstablishment(ctx context.Context, itemID int64) (int64, error) {
	args := i.Called(ctx, itemID)
	return int64(args.Int(0)), args.Error(1)
}

func (i *ItemRepositoryMock) GetMultiple(ctx context.Context, ids []int64) ([]models.Item, error) {
	args := i.Called(ctx, ids)
	return args.Get(0).([]models.Item), args.Error(1)
}

func NewItemRepositoryMock() ItemRepository {
	return &ItemRepositoryMock{}
}
