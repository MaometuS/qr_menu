package admin_interactor

import (
	"context"
	"errors"
	"gitlab.com/maometusu/qr_menu/internal/entity/models"
	"gitlab.com/maometusu/qr_menu/internal/use_case/repository"
	"testing"
)

func TestSendVerificationCode(t *testing.T) {
	type testCase struct {
		repo     *repository.ProfileRepositoryMock
		mailRepo *repository.EmailRepositoryMock
		id       int64
		email    string
		err      string
	}

	cases := make([]testCase, 0)

	repo := &repository.ProfileRepositoryMock{}
	repo.On("SetVerificationCode", context.Background(), int64(1)).Return(errors.New("error"))

	cases = append(cases, testCase{
		repo:     repo,
		mailRepo: nil,
		id:       1,
		email:    "",
		err:      "error",
	})

	repo = &repository.ProfileRepositoryMock{}
	repo.On("SetVerificationCode", context.Background(), int64(1)).Return(nil)
	repo.On("GetOne", context.Background(), int64(1)).Return(&models.Profile{}, errors.New("error1"))

	cases = append(cases, testCase{
		repo:     repo,
		mailRepo: nil,
		id:       1,
		email:    "",
		err:      "error1",
	})

	profile := &models.Profile{
		Email: "email",
	}
	repo = &repository.ProfileRepositoryMock{}
	repo.On("SetVerificationCode", context.Background(), int64(1)).Return(nil)
	repo.On("GetOne", context.Background(), int64(1)).Return(profile, nil)
	mailRepo := &repository.EmailRepositoryMock{}
	mailRepo.On("SendMail", "email", "Код для верификации").Return(errors.New("error2"))

	cases = append(cases, testCase{
		repo:     repo,
		mailRepo: mailRepo,
		id:       1,
		email:    "",
		err:      "error2",
	})

	profile = &models.Profile{
		Email: "email",
	}
	repo = &repository.ProfileRepositoryMock{}
	repo.On("SetVerificationCode", context.Background(), int64(1)).Return(nil)
	repo.On("GetOne", context.Background(), int64(1)).Return(profile, nil)
	mailRepo = &repository.EmailRepositoryMock{}
	mailRepo.On("SendMail", "email", "Код для верификации").Return(nil)

	cases = append(cases, testCase{
		repo:     repo,
		mailRepo: mailRepo,
		id:       1,
		email:    "",
		err:      "",
	})

	for i, el := range cases {
		err := sendVerificationCode(context.Background(), el.repo, el.mailRepo, el.id, el.email)
		if err != nil && err.Error() != el.err {
			t.Error("errors don't match: ", i, err)
		}

		if el.repo != nil {
			el.repo.AssertExpectations(t)
		}

		if el.mailRepo != nil {
			el.mailRepo.AssertExpectations(t)
		}
	}
}
