package profile_repository

import (
	"context"
	"errors"
	"gitlab.com/maometusu/qr_menu/internal/entity"
	"gitlab.com/maometusu/qr_menu/internal/entity/models"
)

func (p *profileRepository) Update(context context.Context, profile *models.Profile) error {
	db, ok := context.Value("db").(entity.PgxIface)
	if !ok {
		return errors.New("connection is missing from context")
	}

	_, err := db.Exec(context, "update profiles set name = $1, email = $2, password = $3", profile.Name, profile.Email, profile.Password)
	if err != nil {
		return err
	}

	return nil
}
