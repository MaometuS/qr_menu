package repository

import (
	"context"
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
