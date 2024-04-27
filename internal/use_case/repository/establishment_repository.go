package repository

import (
	"context"
	"github.com/stretchr/testify/mock"
	"gitlab.com/maometusu/qr_menu/internal/entity/models"
)

type EstablishmentRepository interface {
	GetEstablishments(ctx context.Context, profileID int64) ([]models.Establishment, error)
	CreateEstablishment(ctx context.Context, establishment *models.Establishment) error
	DeleteEstablishment(ctx context.Context, profileID, establishmentID int64) error
	UpdateEstablishment(ctx context.Context, establishment *models.Establishment) error
	GetByLink(ctx context.Context, link string) (*models.Establishment, error)
	DoesLinkExist(ctx context.Context, link string) (bool, error)
	GetOne(ctx context.Context, id int64) (*models.Establishment, error)
}

type EstablishmentRepositoryMock struct {
	mock.Mock
}

func (e *EstablishmentRepositoryMock) GetEstablishments(ctx context.Context, profileID int64) ([]models.Establishment, error) {
	args := e.Called(ctx, profileID)
	return args.Get(0).([]models.Establishment), args.Error(1)
}

func (e *EstablishmentRepositoryMock) CreateEstablishment(ctx context.Context, establishment *models.Establishment) error {
	args := e.Called(ctx, establishment)
	return args.Error(0)
}

func (e *EstablishmentRepositoryMock) DeleteEstablishment(ctx context.Context, profileID, establishmentID int64) error {
	args := e.Called(ctx, profileID, establishmentID)
	return args.Error(0)
}

func (e *EstablishmentRepositoryMock) UpdateEstablishment(ctx context.Context, establishment *models.Establishment) error {
	args := e.Called(ctx, establishment)
	return args.Error(0)
}

func (e *EstablishmentRepositoryMock) GetByLink(ctx context.Context, link string) (*models.Establishment, error) {
	args := e.Called(ctx, link)
	return args.Get(0).(*models.Establishment), args.Error(1)
}

func (e *EstablishmentRepositoryMock) DoesLinkExist(ctx context.Context, link string) (bool, error) {
	args := e.Called(ctx, link)
	return args.Bool(0), args.Error(1)
}

func (e *EstablishmentRepositoryMock) GetOne(ctx context.Context, id int64) (*models.Establishment, error) {
	args := e.Called(ctx, id)
	return args.Get(0).(*models.Establishment), args.Error(1)
}

func NewEstablishmentRepositoryMock() *EstablishmentRepositoryMock {
	return &EstablishmentRepositoryMock{}
}
