package admin_interactor

import (
	"context"
	"errors"
	"gitlab.com/maometusu/qr_menu/internal/entity/models"
	"gitlab.com/maometusu/qr_menu/internal/use_case/presenter"
	"gitlab.com/maometusu/qr_menu/internal/use_case/repository"
	"testing"
)

func TestAdminInteractor_VerifyEmailEditPage(t *testing.T) {
	type testCase struct {
		repo     *repository.ProfileRepositoryMock
		mailRepo *repository.EmailRepositoryMock
		pres     *presenter.AdminPresenterMock
		id       int64
		email    string
		err      string
	}

	cases := make([]testCase, 0)

	profile := &models.Profile{
		Email: "email",
	}
	repo := &repository.ProfileRepositoryMock{}
	repo.On("SetVerificationCode", context.Background(), int64(1)).Return(nil)
	repo.On("GetOne", context.Background(), int64(1)).Return(profile, nil)
	mailRepo := &repository.EmailRepositoryMock{}
	mailRepo.On("SendMail", "email", "Код для верификации").Return(errors.New("error"))

	cases = append(cases, testCase{
		repo:     repo,
		mailRepo: mailRepo,
		id:       1,
		email:    "",
		err:      "error",
	})

	profile = &models.Profile{}
	repo = &repository.ProfileRepositoryMock{}
	repo.On("SetVerificationCode", context.Background(), int64(1)).Return(nil)
	repo.On("GetOne", context.Background(), int64(1)).Return(profile, nil)
	mailRepo = &repository.EmailRepositoryMock{}
	mailRepo.On("SendMail", "email", "Код для верификации").Return(nil)
	pres := &presenter.AdminPresenterMock{}
	pres.On("VerifyEmailEditPage", nil, "email").Return(errors.New("error1"))

	cases = append(cases, testCase{
		repo:     repo,
		mailRepo: mailRepo,
		pres:     pres,
		id:       1,
		email:    "email",
		err:      "error1",
	})

	profile = &models.Profile{}
	repo = &repository.ProfileRepositoryMock{}
	repo.On("SetVerificationCode", context.Background(), int64(1)).Return(nil)
	repo.On("GetOne", context.Background(), int64(1)).Return(profile, nil)
	mailRepo = &repository.EmailRepositoryMock{}
	mailRepo.On("SendMail", "email", "Код для верификации").Return(nil)
	pres = &presenter.AdminPresenterMock{}
	pres.On("VerifyEmailEditPage", nil, "email").Return(nil)

	cases = append(cases, testCase{
		repo:     repo,
		mailRepo: mailRepo,
		pres:     pres,
		id:       1,
		email:    "email",
		err:      "",
	})

	for i, el := range cases {
		interactor := NewAdminInteractor(el.pres, el.mailRepo, el.repo, nil)
		err := interactor.VerifyEmailEditPage(context.Background(), nil, el.id, el.email)
		if err != nil && err.Error() != el.err {
			t.Error("errors don't match: ", i, err)
		}

		if el.repo != nil {
			el.repo.AssertExpectations(t)
		}

		if el.mailRepo != nil {
			el.mailRepo.AssertExpectations(t)
		}

		if el.pres != nil {
			el.pres.AssertExpectations(t)
		}
	}
}
