package item_interactor

import (
	"context"
	"gitlab.com/maometusu/qr_menu/internal/entity/models"
	"gitlab.com/maometusu/qr_menu/internal/use_case/presenter"
	"gitlab.com/maometusu/qr_menu/internal/use_case/repository"
	"io"
	"mime/multipart"
)

type ItemInteractor interface {
	GetItems(ctx context.Context, w io.Writer, profileID, categoryID int64) error
	CreateItem(ctx context.Context, profileID int64, image multipart.File, item *models.Item) error
	MoveUp(ctx context.Context, profileID, itemID int64) error
	MoveDown(ctx context.Context, profileID, itemID int64) error
	EditItem(ctx context.Context, profileID int64, image multipart.File, item *models.Item) error
	DeleteItem(ctx context.Context, profileID, itemID int64) error
}

type itemInteractor struct {
	presenter          presenter.ItemPresenter
	repository         repository.ItemRepository
	fileRepository     repository.FileRepository
	categoryRepository repository.CategoryRepository
}

func NewItemInteractor(
	presenter presenter.ItemPresenter,
	repository repository.ItemRepository,
	fileRepository repository.FileRepository,
	categoryRepository repository.CategoryRepository,
) ItemInteractor {
	return &itemInteractor{
		presenter:          presenter,
		repository:         repository,
		fileRepository:     fileRepository,
		categoryRepository: categoryRepository,
	}
}
