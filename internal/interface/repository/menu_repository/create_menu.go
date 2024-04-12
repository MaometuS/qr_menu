package menu_repository

import (
	"context"
	"errors"
	"gitlab.com/maometusu/qr_menu/internal/entity"
	"gitlab.com/maometusu/qr_menu/internal/entity/models"
)

func (m *menuRepository) CreateMenu(ctx context.Context, menu *models.Menu) error {
	db, ok := ctx.Value("db").(entity.PgxIface)
	if !ok {
		return errors.New("db not found in context")
	}

	_, err := db.Exec(ctx, "update menus set order_param=order_param+1 where order_param >= $1 and establishment_id = $2", menu.OrderParam, menu.EstablishmentID)
	if err != nil {
		return err
	}

	_, err = db.Exec(
		ctx,
		"insert into menus(name, is_visible, establishment_id, order_param) values ($1, $2, $3, $4)",
		menu.Name, menu.IsVisible, menu.EstablishmentID, menu.OrderParam)
	if err != nil {
		return err
	}

	return nil
}
