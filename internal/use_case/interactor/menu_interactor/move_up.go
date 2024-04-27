package menu_interactor

import (
	"context"
	"errors"
)

func (i *interactor) MoveUp(ctx context.Context, profileID, menuID int64) error {
	belongs, err := i.repository.CheckBelongs(ctx, menuID, profileID)
	if err != nil {
		return err
	}

	if !belongs {
		return errors.New("not allowed")
	}

	err = i.repository.MoveUp(ctx, menuID)
	if err != nil {
		return err
	}

	return nil
}
