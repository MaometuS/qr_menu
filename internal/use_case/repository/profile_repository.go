package repository

import (
	"context"
	"gitlab.com/maometusu/qr_menu/internal/entity/models"
)

type ProfileRepository interface {
	GetOneByEmail(context context.Context, email string) (*models.Profile, error)
	GetOne(context context.Context, id int64) (*models.Profile, error)
	CheckExistence(context context.Context, email string) (bool, error)
	Create(context context.Context, profile *models.Profile) (int64, error)
	Update(context context.Context, profile *models.Profile) error
	SetVerificationCode(context context.Context, id int64, verificationCode string) error
	SetVerified(context context.Context, id int64, verified bool) error
}
