package repository

import (
	"context"
	"github.com/stretchr/testify/mock"
	"gitlab.com/maometusu/qr_menu/internal/entity/models"
)

type CommonRepository interface {
	GetCurrencies(ctx context.Context) ([]models.Currency, error)
}

type CommonRepositoryMock struct {
	mock.Mock
}

func (c *CommonRepositoryMock) GetCurrencies(ctx context.Context) ([]models.Currency, error) {
	args := c.Called(ctx)
	return args.Get(0).([]models.Currency), args.Error(1)
}

func NewCommonRepositoryMock() *CommonRepositoryMock {
	return &CommonRepositoryMock{}
}
