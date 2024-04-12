package file_repository

import (
	"context"
	"gitlab.com/maometusu/qr_menu/internal/use_case/repository"
	"mime/multipart"
)

type fileRepository struct {
}

func (f *fileRepository) SaveFile(ctx context.Context, file multipart.File) (string, error) {
	return uploadImage(file, "uploaded")
}

func NewFileRepository() repository.FileRepository {
	return &fileRepository{}
}
