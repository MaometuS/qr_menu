package admin_interactor

import (
	"context"
	"gitlab.com/maometusu/qr_menu/internal/entity"
	"gitlab.com/maometusu/qr_menu/internal/use_case/presenter"
	"gitlab.com/maometusu/qr_menu/internal/use_case/repository"
	"io"
)

type AdminInteractor interface {
	Auth(context context.Context, tokenString string) (int64, error)
	AdminPage(context context.Context, w io.Writer, id int64) error
	ProfilePage(context context.Context, w io.Writer, id int64) error
	LoginPage(context context.Context, w io.Writer) error
	HandleLogin(context context.Context, w io.Writer, email, password string) (string, error)
	RegisterPage(context context.Context, w io.Writer) error
	HandleRegister(context context.Context, w io.Writer, name, email, password, passRepeat string) (int64, error)
	VerifyEmailPage(context context.Context, w io.Writer, id int64) error
	HandleVerifyEmail(context context.Context, w io.Writer, id int64, verificationCode string) (string, error)
	ChangePassword(ctx context.Context, password, passRepeat string, id int64) error
	EditName(ctx context.Context, email, name string, id int64) (bool, error)
	VerifyEmailEditPage(ctx context.Context, w io.Writer, id int64, email string) error
	EditEmail(ctx context.Context, email, verificationCode string, id int64) error
	RestorePasswordPage(ctx context.Context, w io.Writer) error
	RestorePassword(ctx context.Context, w io.Writer, email, pass, passRepeat string) (int64, error)
	VerifyRestorePasswordPage(ctx context.Context, w io.Writer, id int64) error
	VerifyRestorePassword(ctx context.Context, w io.Writer, id int64, verificationCode string) error
}

type adminInteractor struct {
	presenter         presenter.AdminPresenter
	emailRepository   repository.EmailRepository
	profileRepository repository.ProfileRepository
	config            *entity.Config
}

func NewAdminInteractor(
	presenter presenter.AdminPresenter,
	emailRepository repository.EmailRepository,
	repository repository.ProfileRepository,
	config *entity.Config,
) AdminInteractor {
	return &adminInteractor{
		presenter:         presenter,
		emailRepository:   emailRepository,
		profileRepository: repository,
		config:            config,
	}
}
