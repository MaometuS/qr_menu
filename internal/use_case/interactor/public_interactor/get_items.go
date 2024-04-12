package public_interactor

import (
	"context"
	"io"
)

func (p *publicInteractor) GetItems(ctx context.Context, w io.Writer, categoryID int64, selectedItems map[int64][]int64) error {
	category, err := p.categoryRepository.GetOne(ctx, categoryID)
	if err != nil {
		return err
	}

	items, err := p.itemsRepository.GetAll(ctx, categoryID)
	if err != nil {
		return err
	}

	var establishmentID int64
	if len(items) > 0 {
		establishmentID, err = p.itemsRepository.GetEstablishment(ctx, items[0].ID)
		if err != nil {
			return err
		}
	}

	mp := make(map[int64]int64)
	for _, el := range selectedItems[establishmentID] {
		mp[el]++
	}

	for i := range items {
		items[i].SelectedAmount = mp[items[i].ID]
	}

	return p.presenter.PresentItems(w, category.Name, items, category.MenuID)
}
