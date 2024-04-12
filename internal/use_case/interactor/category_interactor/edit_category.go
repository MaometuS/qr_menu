package category_interactor

import (
	"context"
	"errors"
	"gitlab.com/maometusu/qr_menu/internal/entity/models"
	"mime/multipart"
)

func (c *categoryInteractor) EditCategory(ctx context.Context, profileID int64, background multipart.File, category *models.Category) error {
	menuBelongs, err := c.menuRepository.CheckBelongs(ctx, category.MenuID, profileID)
	if err != nil {
		return err
	}

	if !menuBelongs {
		return errors.New("menu does not belong to current profile")
	}

	if background != nil {
		category.Background, err = c.fileRepository.SaveFile(ctx, background)
		if err != nil {
			return err
		}
	}

	err = c.repository.UpdateCategory(ctx, category)
	if err != nil {
		return err
	}

	return nil
}
