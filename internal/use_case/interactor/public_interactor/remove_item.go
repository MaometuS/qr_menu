package public_interactor

import (
	"context"
	"slices"
)

func (p *publicInteractor) RemoveItem(ctx context.Context, itemID int64, selectedItems map[int64][]int64) (map[int64][]int64, error) {
	establishmentID, err := p.itemsRepository.GetEstablishment(ctx, itemID)
	if err != nil {
		return nil, err
	}

	ind := slices.Index(selectedItems[establishmentID], itemID)
	if ind > -1 {
		selectedItems[establishmentID] = slices.Delete(selectedItems[establishmentID], ind, ind+1)
	}

	return selectedItems, nil
}
