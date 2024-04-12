package profile_repository

import (
	"context"
	"errors"
	"gitlab.com/maometusu/qr_menu/internal/entity"
)

func (p *profileRepository) SetVerificationCode(context context.Context, id int64, verificationCode string) error {
	db, ok := context.Value("db").(entity.PgxIface)
	if !ok {
		return errors.New("connection is missing from context")
	}

	_, err := db.Exec(context, "update profiles set verification_code = $1 where id = $2", verificationCode, id)
	if err != nil {
		return err
	}

	return nil
}
