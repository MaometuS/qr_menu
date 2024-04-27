package establishment_interactor

import (
	"context"
	"github.com/stretchr/testify/mock"
	"gitlab.com/maometusu/qr_menu/internal/entity/models"
	"io"
	"mime/multipart"
)

type TestImplementation struct {
	mock.Mock
}

func (t *TestImplementation) GetEstablishments(context context.Context, w io.Writer, profileID int64) error {
	args := t.Called(context, w, profileID)
	return args.Error(0)
}

func (t *TestImplementation) CreateEstablishment(context context.Context, establishment *models.Establishment) error {
	args := t.Called(context, establishment)
	return args.Error(0)
}

func (t *TestImplementation) DeleteEstablishment(context context.Context, profileID, establishmentID int64) error {
	args := t.Called(context, profileID, establishmentID)
	return args.Error(0)
}

func (t *TestImplementation) EditEstablishmentPage(context context.Context, w io.Writer, profileID int64, link string) error {
	args := t.Called(context, w, profileID, link)
	return args.Error(0)
}

func (t *TestImplementation) EditEstablishment(ctx context.Context, logo, background multipart.File, establishment *models.Establishment) error {
	args := t.Called(ctx, establishment)
	return args.Error(0)
}

func (t *TestImplementation) EstablishmentInfo(context context.Context, w io.Writer, profileID, establishmentID int64) error {
	args := t.Called(context, w, profileID, establishmentID)
	return args.Error(0)
}

func (t *TestImplementation) GetQrCode(ctx context.Context, w io.Writer, host string, establishmentID int64) error {
	args := t.Called(ctx, w, host, establishmentID)
	return args.Error(0)
}

func (t *TestImplementation) DoesLinkExist(ctx context.Context, link string) (bool, error) {
	args := t.Called(ctx, link)
	return args.Bool(0), args.Error(1)
}

func NewTestImplementation() *TestImplementation {
	return &TestImplementation{}
}
