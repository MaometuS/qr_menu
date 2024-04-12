package public_interactor

import (
	"context"
	"io"
)

func (p *publicInteractor) GetOrder(ctx context.Context, w io.Writer, menuID int64, selectedItems map[int64][]int64) error {
	menu, err := p.menuRepository.GetOne(ctx, menuID)
	if err != nil {
		return err
	}

	items := selectedItems[menu.EstablishmentID]
	cnts := make(map[int64]int64)
	for _, el := range items {
		cnts[el]++
	}

	ids := make([]int64, 0)
	for key, _ := range cnts {
		ids = append(ids, key)
	}

	itemsObj, err := p.itemsRepository.GetMultiple(ctx, ids)
	if err != nil {
		return err
	}

	var total float64
	for i := range itemsObj {
		total += itemsObj[i].Price * float64(cnts[itemsObj[i].ID])
		itemsObj[i].SelectedAmount = cnts[itemsObj[i].ID]
	}

	return p.presenter.PresentOrder(w, itemsObj, total, menuID)
}
