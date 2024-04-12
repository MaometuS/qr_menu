package profile_repository

import (
	"context"
	"errors"
	"gitlab.com/maometusu/qr_menu/internal/entity"
	"gitlab.com/maometusu/qr_menu/internal/entity/models"
)

func (p *profileRepository) Create(context context.Context, profile *models.Profile) (int64, error) {
	db, ok := context.Value("db").(entity.PgxIface)
	if !ok {
		return 0, errors.New("connection is missing from context")
	}

	var id int64
	err := db.QueryRow(context, "insert into profiles(name, email, password) values ($1, $2, $3) returning id", profile.Name, profile.Email, profile.Password).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}
