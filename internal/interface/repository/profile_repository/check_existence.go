package profile_repository

import (
	"context"
	"errors"
	"gitlab.com/maometusu/qr_menu/internal/entity"
)

func (p *profileRepository) CheckExistence(context context.Context, email string) (bool, error) {
	db, ok := context.Value("db").(entity.PgxIface)
	if !ok {
		return false, errors.New("connection is missing from context")
	}

	var exists int64
	err := db.QueryRow(context, "select count(id) from profiles where email = $1 and verified = true", email).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists > 0, nil
}
