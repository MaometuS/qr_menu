package public_interactor

import (
	"context"
	"io"
)

func (p *publicInteractor) GetSearch(ctx context.Context, w io.Writer, search string, establishmentID int64, menuID int64, selectedItems map[int64][]int64) error {
	items, err := p.itemsRepository.Search(ctx, search, establishmentID)
	if err != nil {
		return err
	}

	mp := make(map[int64]int64)
	for _, el := range selectedItems[establishmentID] {
		mp[el]++
	}

	for i := range items {
		items[i].SelectedAmount = mp[items[i].ID]
	}

	return p.presenter.PresentSearch(w, search, items, establishmentID, menuID)
}
