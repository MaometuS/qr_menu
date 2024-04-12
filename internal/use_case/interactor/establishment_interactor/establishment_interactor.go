package establishment_interactor

import (
	"context"
	"gitlab.com/maometusu/qr_menu/internal/entity/models"
	"gitlab.com/maometusu/qr_menu/internal/use_case/presenter"
	"gitlab.com/maometusu/qr_menu/internal/use_case/repository"
	"io"
	"mime/multipart"
)

type EstablishmentInteractor interface {
	GetEstablishments(context context.Context, w io.Writer, profileID int64) error
	CreateEstablishment(context context.Context, establishment *models.Establishment) error
	DeleteEstablishment(context context.Context, profileID, establishmentID int64) error
	EditEstablishmentPage(context context.Context, w io.Writer, profileID int64, link string) error
	EditEstablishment(ctx context.Context, logo, background multipart.File, establishment *models.Establishment) error
	EstablishmentInfo(context context.Context, w io.Writer, profileID, establishmentID int64) error
	GetQrCode(ctx context.Context, w io.Writer, host string, establishmentID int64) error
	DoesLinkExist(ctx context.Context, link string) (bool, error)
}

type establishmentInteractor struct {
	presenter        presenter.EstablishmentPresenter
	repository       repository.EstablishmentRepository
	menuRepository   repository.MenuRepository
	fileRepository   repository.FileRepository
	commonRepository repository.CommonRepository
}

func NewEstablishmentInteractor(
	presenter presenter.EstablishmentPresenter,
	repository repository.EstablishmentRepository,
	menuRepository repository.MenuRepository,
	fileRepository repository.FileRepository,
	commonRepository repository.CommonRepository,
) EstablishmentInteractor {
	return &establishmentInteractor{
		presenter:        presenter,
		repository:       repository,
		menuRepository:   menuRepository,
		fileRepository:   fileRepository,
		commonRepository: commonRepository,
	}
}
