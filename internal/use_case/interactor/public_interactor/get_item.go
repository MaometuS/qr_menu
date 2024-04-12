package public_interactor

import (
	"context"
	"io"
)

func (p *publicInteractor) GetItem(ctx context.Context, w io.Writer, itemID int64, selectedItems map[int64][]int64) error {
	item, err := p.itemsRepository.GetOne(ctx, itemID)
	if err != nil {
		return err
	}

	establishmentID, err := p.itemsRepository.GetEstablishment(ctx, itemID)
	if err != nil {
		return err
	}

	for _, el := range selectedItems[establishmentID] {
		if el == itemID {
			item.SelectedAmount++
		}
	}

	return p.presenter.PresentItem(w, item, len(selectedItems[establishmentID]) > 0)
}
