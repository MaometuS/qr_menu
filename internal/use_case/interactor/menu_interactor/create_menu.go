package menu_interactor

import (
	"context"
	"errors"
	"gitlab.com/maometusu/qr_menu/internal/entity/models"
)

func (i *interactor) CreateMenu(ctx context.Context, profileID int64, menu *models.Menu) error {
	establishment, err := i.establishmentRepository.GetOne(ctx, menu.EstablishmentID)
	if err != nil {
		return err
	}

	if establishment.ProfileID != profileID {
		return errors.New("not allowed")
	}

	err = i.repository.CreateMenu(ctx, menu)
	if err != nil {
		return err
	}

	return nil
}
