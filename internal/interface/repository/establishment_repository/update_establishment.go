package establishment_repository

import (
	"context"
	"errors"
	"gitlab.com/maometusu/qr_menu/internal/entity"
	"gitlab.com/maometusu/qr_menu/internal/entity/models"
)

func (e *establishmentRepository) UpdateEstablishment(ctx context.Context, establishment *models.Establishment) error {
	db, ok := ctx.Value("db").(entity.PgxIface)
	if !ok {
		return errors.New("db not found in context")
	}

	_, err := db.Exec(
		ctx,
		`update establishments set 
				name = $1, 
				phone = $2, 
				wifi_password = $3, 
				can_make_orders = $4,
				country = $5,
				city = $6,
				address = $7,
				text = $8,
				currency_id = $9
			where id = $10`,
		establishment.Name,
		establishment.Phone,
		establishment.WifiPassword,
		establishment.CanMakeOrders,
		establishment.Country,
		establishment.City,
		establishment.Address,
		establishment.Text,
		establishment.CurrencyID,
		establishment.ID,
	)
	if err != nil {
		return err
	}

	if establishment.Logo != "" {
		_, err = db.Exec(ctx, "update establishments set logo = $1 where id = $2", establishment.Logo, establishment.ID)
		if err != nil {
			return err
		}
	}

	if establishment.Background != "" {
		_, err = db.Exec(ctx, "update establishments set background = $1 where id = $2", establishment.Background, establishment.ID)
		if err != nil {
			return err
		}
	}

	return nil
}
