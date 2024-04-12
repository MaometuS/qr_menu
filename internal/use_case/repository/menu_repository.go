package repository

import (
	"context"
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
