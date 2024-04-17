package file_repository

import (
	"context"
	"gitlab.com/maometusu/qr_menu/internal/entity"
	"gitlab.com/maometusu/qr_menu/internal/use_case/repository"
	"mime/multipart"
)

type fileRepository struct {
	config *entity.Config
}

func (f *fileRepository) SaveFile(ctx context.Context, file multipart.File) (string, error) {
	return uploadImage(file, f.config.UploadFolder)
}

func NewFileRepository(config *entity.Config) repository.FileRepository {
	return &fileRepository{config}
}
