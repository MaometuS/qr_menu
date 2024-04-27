package menu_interactor

import (
	"context"
	"errors"
)

func (i *interactor) MoveDown(ctx context.Context, profileID, menuID int64) error {
	belongs, err := i.repository.CheckBelongs(ctx, menuID, profileID)
	if err != nil {
		return err
	}

	if !belongs {
		return errors.New("not allowed")
	}

	err = i.repository.MoveDown(ctx, menuID)
	if err != nil {
		return err
	}

	return nil
}
