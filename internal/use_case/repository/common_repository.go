package repository

import (
	"context"
	"gitlab.com/maometusu/qr_menu/internal/entity/models"
)

type CommonRepository interface {
	GetCurrencies(ctx context.Context) ([]models.Currency, error)
}
