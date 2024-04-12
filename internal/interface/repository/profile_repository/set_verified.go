package profile_repository

import (
	"context"
	"errors"
	"gitlab.com/maometusu/qr_menu/internal/entity"
)

func (p *profileRepository) SetVerified(context context.Context, id int64, verified bool) error {
	db, ok := context.Value("db").(entity.PgxIface)
	if !ok {
		return errors.New("connection is missing from context")
	}

	_, err := db.Exec(context, "update profiles set verified = $2 where id = $1", id, verified)
	if err != nil {
		return err
	}

	return nil
}
