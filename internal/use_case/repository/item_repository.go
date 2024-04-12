package repository

import (
	"context"
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
