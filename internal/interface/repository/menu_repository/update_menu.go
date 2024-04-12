package menu_repository

import (
	"context"
	"errors"
	"gitlab.com/maometusu/qr_menu/internal/entity"
	"gitlab.com/maometusu/qr_menu/internal/entity/models"
)

func (m *menuRepository) UpdateMenu(ctx context.Context, menu *models.Menu) error {
	db, ok := ctx.Value("db").(entity.PgxIface)
	if !ok {
		return errors.New("db not found in context")
	}

	_, err := db.Exec(
		ctx,
		"update menus set name = $1, is_visible = $2, establishment_id = $3 where id = $4",
		menu.Name, menu.IsVisible, menu.EstablishmentID, menu.ID)
	if err != nil {
		return err
	}

	return nil
}
