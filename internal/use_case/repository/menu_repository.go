package repository

import (
	"context"
	"github.com/stretchr/testify/mock"
	"gitlab.com/maometusu/qr_menu/internal/entity/models"
)

type MenuRepository interface {
	GetAll(ctx context.Context, establishmentID int64) ([]models.Menu, error)
	GetOne(ctx context.Context, menuID int64) (*models.Menu, error)
	GetTopID(ctx context.Context, establishmentID int64) (int64, error)
	CreateMenu(ctx context.Context, menu *models.Menu) error
	UpdateMenu(ctx context.Context, menu *models.Menu) error
	DeleteMenu(ctx context.Context, menuID int64) error
	CheckBelongs(ctx context.Context, menuID, profileID int64) (bool, error)
	MoveDown(ctx context.Context, menuID int64) error
	MoveUp(ctx context.Context, menuID int64) error
}

type MenuRepositoryMock struct {
	mock.Mock
}

func (m *MenuRepositoryMock) GetAll(ctx context.Context, establishmentID int64) ([]models.Menu, error) {
	args := m.Called(ctx, establishmentID)
	return args.Get(0).([]models.Menu), args.Error(1)
}

func (m *MenuRepositoryMock) GetOne(ctx context.Context, menuID int64) (*models.Menu, error) {
	args := m.Called(ctx, menuID)
	return args.Get(0).(*models.Menu), args.Error(1)
}

func (m *MenuRepositoryMock) GetTopID(ctx context.Context, establishmentID int64) (int64, error) {
	args := m.Called(ctx, establishmentID)
	return int64(args.Int(0)), args.Error(1)
}

func (m *MenuRepositoryMock) CreateMenu(ctx context.Context, menu *models.Menu) error {
	args := m.Called(ctx, menu)
	return args.Error(0)
}

func (m *MenuRepositoryMock) UpdateMenu(ctx context.Context, menu *models.Menu) error {
	args := m.Called(ctx, menu)
	return args.Error(0)
}

func (m *MenuRepositoryMock) DeleteMenu(ctx context.Context, menuID int64) error {
	args := m.Called(ctx, menuID)
	return args.Error(0)
}

func (m *MenuRepositoryMock) CheckBelongs(ctx context.Context, menuID, profileID int64) (bool, error) {
	args := m.Called(ctx, menuID, profileID)
	return args.Bool(0), args.Error(1)
}

func (m *MenuRepositoryMock) MoveDown(ctx context.Context, menuID int64) error {
	args := m.Called(ctx, menuID)
	return args.Error(0)
}

func (m *MenuRepositoryMock) MoveUp(ctx context.Context, menuID int64) error {
	args := m.Called(ctx, menuID)
	return args.Error(0)
}

func NewMenuRepositoryMock() *MenuRepositoryMock {
	return &MenuRepositoryMock{}
}
