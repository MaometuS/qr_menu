package public_interactor

import (
	"context"
	"io"
)

func (p *publicInteractor) GetMainPage(ctx context.Context, w io.Writer, link string, selectedItems map[int64][]int64) error {
	establishment, err := p.establishmentRepository.GetByLink(ctx, link)
	if err != nil {
		return err
	}

	menus, err := p.menuRepository.GetAll(ctx, establishment.ID)
	if err != nil {
		return err
	}

	var topMenuID int64
	for _, el := range menus {
		if el.IsVisible {
			topMenuID = el.ID
			break
		}
	}

	return p.presenter.PresentMainPage(w, establishment, menus, topMenuID, len(selectedItems[establishment.ID]) > 0)
}
