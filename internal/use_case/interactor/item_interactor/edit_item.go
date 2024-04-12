package item_interactor

import (
	"context"
	"errors"
	"gitlab.com/maometusu/qr_menu/internal/entity/models"
	"mime/multipart"
)

func (i *itemInteractor) EditItem(ctx context.Context, profileID int64, image multipart.File, item *models.Item) error {
	categoryBelongs, err := i.categoryRepository.CheckBelongs(ctx, item.CategoryID, profileID)
	if err != nil {
		return err
	}

	if !categoryBelongs {
		return errors.New("category does not belong to current profile")
	}

	if image != nil {
		item.Image, err = i.fileRepository.SaveFile(ctx, image)
		if err != nil {
			return err
		}
	}

	err = i.repository.UpdateItem(ctx, item)
	if err != nil {
		return err
	}

	return nil
}
