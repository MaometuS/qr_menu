package public_interactor

import (
	"context"
	"gitlab.com/maometusu/qr_menu/internal/use_case/presenter"
	"gitlab.com/maometusu/qr_menu/internal/use_case/repository"
	"io"
)

type PublicInteractor interface {
	GetMainPage(ctx context.Context, w io.Writer, link string, selectedItems map[int64][]int64) error
	GetCategories(ctx context.Context, w io.Writer, menuID int64) error
	GetItems(ctx context.Context, w io.Writer, categoryID int64, selectedItems map[int64][]int64) error
	GetSearch(ctx context.Context, w io.Writer, search string, establishmentID int64, menuID int64, selectedItems map[int64][]int64) error
	GetItem(ctx context.Context, w io.Writer, itemID int64, selectedItems map[int64][]int64) error
	AddItem(ctx context.Context, itemID int64, selectedItems map[int64][]int64) (map[int64][]int64, error)
	RemoveItem(ctx context.Context, itemID int64, selectedItems map[int64][]int64) (map[int64][]int64, error)
	GetOrder(ctx context.Context, w io.Writer, menuID int64, selectedItems map[int64][]int64) error
}

type publicInteractor struct {
	presenter               presenter.PublicPresenter
	establishmentRepository repository.EstablishmentRepository
	menuRepository          repository.MenuRepository
	categoryRepository      repository.CategoryRepository
	itemsRepository         repository.ItemRepository
}

func NewPublicInteractor(
	presenter presenter.PublicPresenter,
	establishmentRepository repository.EstablishmentRepository,
	menuRepository repository.MenuRepository,
	categoryRepository repository.CategoryRepository,
	itemsRepository repository.ItemRepository,
) PublicInteractor {
	return &publicInteractor{
		presenter:               presenter,
		establishmentRepository: establishmentRepository,
		menuRepository:          menuRepository,
		categoryRepository:      categoryRepository,
		itemsRepository:         itemsRepository,
	}
}
