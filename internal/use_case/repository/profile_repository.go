package repository

import (
	"context"
	"github.com/stretchr/testify/mock"
	"gitlab.com/maometusu/qr_menu/internal/entity/models"
)

type ProfileRepository interface {
	GetOneByEmail(context context.Context, email string) (*models.Profile, error)
	GetOne(context context.Context, id int64) (*models.Profile, error)
	CheckExistence(context context.Context, email string) (bool, error)
	Create(context context.Context, profile *models.Profile) (int64, error)
	Update(context context.Context, profile *models.Profile) error
	SetVerificationCode(context context.Context, id int64, verificationCode string) error
	SetVerified(context context.Context, id int64, verified bool) error
}

type ProfileRepositoryMock struct {
	mock.Mock
}

func (p *ProfileRepositoryMock) GetOneByEmail(context context.Context, email string) (*models.Profile, error) {
	args := p.Called(context, email)
	return args.Get(0).(*models.Profile), args.Error(1)
}

func (p *ProfileRepositoryMock) GetOne(context context.Context, id int64) (*models.Profile, error) {
	args := p.Called(context, id)
	return args.Get(0).(*models.Profile), args.Error(1)
}

func (p *ProfileRepositoryMock) CheckExistence(context context.Context, email string) (bool, error) {
	args := p.Called(context, email)
	return args.Bool(0), args.Error(1)
}

func (p *ProfileRepositoryMock) Create(context context.Context, profile *models.Profile) (int64, error) {
	args := p.Called(context, profile)
	return int64(args.Int(0)), args.Error(1)
}

func (p *ProfileRepositoryMock) Update(context context.Context, profile *models.Profile) error {
	args := p.Called(context, profile)
	return args.Error(0)
}

func (p *ProfileRepositoryMock) SetVerificationCode(context context.Context, id int64, verificationCode string) error {
	args := p.Called(context, id)
	return args.Error(0)
}

func (p *ProfileRepositoryMock) SetVerified(context context.Context, id int64, verified bool) error {
	args := p.Called(context, id, verified)
	return args.Error(0)
}

func NewProfileRepositoryMock() *ProfileRepositoryMock {
	return &ProfileRepositoryMock{}
}
