package admin_interactor

import (
	"context"
	"errors"
	"gitlab.com/maometusu/qr_menu/internal/entity/models"
	"gitlab.com/maometusu/qr_menu/internal/use_case/presenter"
	"gitlab.com/maometusu/qr_menu/internal/use_case/repository"
	"testing"
)

func TestAdminInteractor_VerifyRestorePasswordPage(t *testing.T) {
	profile := &models.Profile{
		Email: "email",
	}
	repo := &repository.ProfileRepositoryMock{}
	repo.On("SetVerificationCode", context.Background(), int64(1)).Return(nil)
	repo.On("GetOne", context.Background(), int64(1)).Return(profile, nil)
	mailRepo := &repository.EmailRepositoryMock{}
	mailRepo.On("SendMail", "email", "Код для верификации").Return(nil)
	pres := presenter.NewAdminPresenterMock()
	pres.On("VerifyRestorePasswordPage", nil, int64(1), false, false).Return(errors.New("error"))

	inter := NewAdminInteractor(pres, mailRepo, repo, nil)
	err := inter.VerifyRestorePasswordPage(context.Background(), nil, 1)
	if err == nil {
		t.Error("must have error")
	}
	if err != nil && err.Error() != "error" {
		t.Error("error mismatch")
	}

	repo.AssertExpectations(t)
	mailRepo.AssertExpectations(t)
	pres.AssertExpectations(t)

	repo = &repository.ProfileRepositoryMock{}
	repo.On("SetVerificationCode", context.Background(), int64(1)).Return(nil)
	repo.On("GetOne", context.Background(), int64(1)).Return(profile, nil)
	mailRepo = &repository.EmailRepositoryMock{}
	mailRepo.On("SendMail", "email", "Код для верификации").Return(nil)
	pres = presenter.NewAdminPresenterMock()
	pres.On("VerifyRestorePasswordPage", nil, int64(1), false, false).Return(nil)

	inter = NewAdminInteractor(pres, mailRepo, repo, nil)
	err = inter.VerifyRestorePasswordPage(context.Background(), nil, 1)
	if err != nil {
		t.Error("has error")
	}

	repo.AssertExpectations(t)
	mailRepo.AssertExpectations(t)
	pres.AssertExpectations(t)
}
