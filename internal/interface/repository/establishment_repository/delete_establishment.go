package establishment_repository

import (
	"context"
	"errors"
	"gitlab.com/maometusu/qr_menu/internal/entity"
)

func (e *establishmentRepository) DeleteEstablishment(ctx context.Context, profileID, establishmentID int64) error {
	db, ok := ctx.Value("db").(entity.PgxIface)
	if !ok {
		return errors.New("db not found in context")
	}

	_, err := db.Exec(ctx, "delete from establishments where id = $1 and profile_id = $2", establishmentID, profileID)
	if err != nil {
		return err
	}

	return nil
}
