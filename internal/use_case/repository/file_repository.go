package repository

import (
	"context"
	"mime/multipart"
)

type FileRepository interface {
	SaveFile(ctx context.Context, file multipart.File) (string, error)
}
