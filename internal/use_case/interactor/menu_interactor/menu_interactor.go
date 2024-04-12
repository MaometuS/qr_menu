package menu_interactor

import (
	"context"
	"gitlab.com/maometusu/qr_menu/internal/entity/models"
	"gitlab.com/maometusu/qr_menu/internal/use_case/presenter"
	"gitlab.com/maometusu/qr_menu/internal/use_case/repository"
	"io"
)

type MenuInteractor interface {
	GetMenus(ctx context.Context, w io.Writer, profileID, establishmentID int64) error
	CreateMenu(ctx context.Context, profileID int64, menu *models.Menu) error
	MoveUp(ctx context.Context, profileID, menuID int64) error
	MoveDown(ctx context.Context, profileID, menuID int64) error
	EditMenu(ctx context.Context, profileID int64, menu *models.Menu) error
	DeleteMenu(ctx context.Context, profileID, menuID int64) error
}

type interactor struct {
	presenter               presenter.MenuPresenter
	repository              repository.MenuRepository
	establishmentRepository repository.EstablishmentRepository
}

func NewMenuInteractor(
	presenter presenter.MenuPresenter,
	repository repository.MenuRepository,
	establishmentRepository repository.EstablishmentRepository,
) MenuInteractor {
	return &interactor{
		presenter:               presenter,
		repository:              repository,
		establishmentRepository: establishmentRepository,
	}
}
