package category_interactor

import (
	"context"
	"errors"
)

func (c *categoryInteractor) MoveDown(ctx context.Context, profileID, categoryID int64) error {
	categoryBelongs, err := c.repository.CheckBelongs(ctx, categoryID, profileID)
	if err != nil {
		return err
	}

	if !categoryBelongs {
		return errors.New("category does not belong to current profile")
	}

	err = c.repository.MoveDown(ctx, categoryID)
	if err != nil {
		return err
	}

	return nil
}
