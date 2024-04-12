package category_interactor

import (
	"context"
	"gitlab.com/maometusu/qr_menu/internal/entity/models"
	"gitlab.com/maometusu/qr_menu/internal/use_case/presenter"
	"gitlab.com/maometusu/qr_menu/internal/use_case/repository"
	"io"
	"mime/multipart"
)

type CategoryInteractor interface {
	GetCategories(ctx context.Context, w io.Writer, profileID, menuID int64) error
	CreateCategory(ctx context.Context, profileID int64, background multipart.File, category *models.Category) error
	MoveUp(ctx context.Context, profileID, categoryID int64) error
	MoveDown(ctx context.Context, profileID, categoryID int64) error
	EditCategory(ctx context.Context, profileID int64, background multipart.File, category *models.Category) error
	DeleteCategory(ctx context.Context, profileID, categoryID int64) error
}

type categoryInteractor struct {
	presenter      presenter.CategoryPresenter
	repository     repository.CategoryRepository
	menuRepository repository.MenuRepository
	fileRepository repository.FileRepository
}

func NewCategoryInteractor(
	presenter presenter.CategoryPresenter,
	repository repository.CategoryRepository,
	menuRepository repository.MenuRepository,
	fileRepository repository.FileRepository,
) CategoryInteractor {
	return &categoryInteractor{
		presenter:      presenter,
		repository:     repository,
		menuRepository: menuRepository,
		fileRepository: fileRepository,
	}
}
