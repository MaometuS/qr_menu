package public_interactor

import (
	"context"
	"io"
)

func (p *publicInteractor) GetCategories(ctx context.Context, w io.Writer, menuID int64) error {
	menu, err := p.menuRepository.GetOne(ctx, menuID)
	if err != nil {
		return err
	}

	categories, err := p.categoryRepository.GetAll(ctx, menuID)
	if err != nil {
		return err
	}

	return p.presenter.PresentCategories(w, menu.Name, categories, menu.EstablishmentID, menuID)
}
