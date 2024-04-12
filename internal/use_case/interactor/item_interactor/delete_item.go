package item_interactor

import (
	"context"
	"errors"
)

func (i *itemInteractor) DeleteItem(ctx context.Context, profileID, itemID int64) error {
	itemBelongs, err := i.repository.CheckBelongs(ctx, itemID, profileID)
	if err != nil {
		return err
	}

	if !itemBelongs {
		return errors.New("item does not belong to current profile")
	}

	err = i.repository.DeleteItem(ctx, itemID)
	if err != nil {
		return err
	}

	return nil
}
