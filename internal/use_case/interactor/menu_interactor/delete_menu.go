package menu_interactor

import (
	"context"
	"errors"
)

func (i *interactor) DeleteMenu(ctx context.Context, profileID, menuID int64) error {
	menu, err := i.repository.GetOne(ctx, menuID)
	if err != nil {
		return err
	}

	establishment, err := i.establishmentRepository.GetOne(ctx, menu.EstablishmentID)
	if err != nil {
		return err
	}

	if establishment.ProfileID != profileID {
		return errors.New("not allowed")
	}

	err = i.repository.DeleteMenu(ctx, menuID)
	if err != nil {
		return err
	}

	return nil
}
