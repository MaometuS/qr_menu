package public_interactor

import "context"

func (p *publicInteractor) AddItem(ctx context.Context, itemID int64, selectedItems map[int64][]int64) (map[int64][]int64, error) {
	establishmentID, err := p.itemsRepository.GetEstablishment(ctx, itemID)
	if err != nil {
		return nil, err
	}

	selectedItems[establishmentID] = append(selectedItems[establishmentID], itemID)
	return selectedItems, nil
}
