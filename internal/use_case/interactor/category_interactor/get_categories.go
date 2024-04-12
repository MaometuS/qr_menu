package category_interactor

import (
	"context"
	"io"
)

func (c *categoryInteractor) GetCategories(ctx context.Context, w io.Writer, profileID, menuID int64) error {
	categories, err := c.repository.GetAll(ctx, menuID)
	if err != nil {
		return err
	}

	err = c.presenter.PresentCategories(w, categories, menuID)
	if err != nil {
		return err
	}

	return nil
}
