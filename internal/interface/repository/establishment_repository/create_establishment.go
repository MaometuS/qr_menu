package establishment_repository

import (
	"context"
	"errors"
	"gitlab.com/maometusu/qr_menu/internal/entity"
	"gitlab.com/maometusu/qr_menu/internal/entity/models"
)

func (e *establishmentRepository) CreateEstablishment(ctx context.Context, establishment *models.Establishment) error {
	db, ok := ctx.Value("db").(entity.PgxIface)
	if !ok {
		return errors.New("db not found in context")
	}

	_, err := db.Exec(
		ctx,
		"insert into establishments(name, profile_id, currency_id, link) values ($1, $2, $3, $4)",
		establishment.Name,
		establishment.ProfileID,
		establishment.CurrencyID,
		establishment.Link,
	)
	if err != nil {
		return err
	}

	return nil
}
