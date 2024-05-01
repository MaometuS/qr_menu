package establishment_interactor

import (
	"context"
	"errors"
	"gitlab.com/maometusu/qr_menu/internal/entity/models"
	"mime/multipart"
	"strings"
)

func (e *establishmentInteractor) EditEstablishment(ctx context.Context, logo, background multipart.File, establishment *models.Establishment) error {
	prevEst, err := e.repository.GetOne(ctx, establishment.ID)
	if err != nil {
		return err
	}

	if prevEst.ProfileID != establishment.ProfileID {
		return errors.New("not allowed for current profile")
	}

	if logo != nil {
		establishment.Logo, err = e.fileRepository.SaveFile(ctx, logo)
		if err != nil {
			return err
		}
	}

	if background != nil {
		establishment.Background, err = e.fileRepository.SaveFile(ctx, background)
		if err != nil {
			return err
		}
	}

	establishment.Logo = strings.ToLower(establishment.Link)

	err = e.repository.UpdateEstablishment(ctx, establishment)
	if err != nil {
		return err
	}

	return nil
}
