package repository

import (
	"context"
	"github.com/stretchr/testify/mock"
	"mime/multipart"
)

type FileRepository interface {
	SaveFile(ctx context.Context, file multipart.File) (string, error)
}

type FileRepositoryMock struct {
	mock.Mock
}

func (f *FileRepositoryMock) SaveFile(ctx context.Context, file multipart.File) (string, error) {
	args := f.Called(ctx, file)
	return args.String(0), args.Error(1)
}

func NewFileRepositoryMock() *FileRepositoryMock {
	return &FileRepositoryMock{}
}
