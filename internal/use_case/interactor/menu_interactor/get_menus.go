package menu_interactor

import (
	"context"
	"errors"
	"io"
)

func (i *interactor) GetMenus(ctx context.Context, w io.Writer, profileID, establishmentID int64) error {
	establishment, err := i.establishmentRepository.GetOne(ctx, establishmentID)
	if err != nil {
		return err
	}

	if establishment.ProfileID != profileID {
		return errors.New("not allowed")
	}

	menus, err := i.repository.GetAll(ctx, establishmentID)
	if err != nil {
		return err
	}

	err = i.presenter.PresentMenus(w, menus, establishmentID)
	if err != nil {
		return err
	}

	return nil
}
