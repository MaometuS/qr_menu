package repository

import (
	"context"
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
