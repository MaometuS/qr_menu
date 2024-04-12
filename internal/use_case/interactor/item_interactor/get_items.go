package item_interactor

import (
	"context"
	"io"
)

func (i *itemInteractor) GetItems(ctx context.Context, w io.Writer, profileID, categoryID int64) error {
	category, err := i.categoryRepository.GetOne(ctx, categoryID)
	if err != nil {
		return err
	}

	items, err := i.repository.GetAll(ctx, categoryID)
	if err != nil {
		return err
	}

	err = i.presenter.PresentItems(w, items, categoryID, category.MenuID)
	if err != nil {
		return err
	}

	return nil
}
