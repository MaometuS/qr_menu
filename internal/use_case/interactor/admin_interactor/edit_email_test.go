package admin_interactor

import (
	"context"
	"errors"
	"gitlab.com/maometusu/qr_menu/internal/entity/models"
	"gitlab.com/maometusu/qr_menu/internal/use_case/repository"
	"testing"
)

func TestAdminInteractor_EditEmail(t *testing.T) {
	type testCase struct {
		ctx              context.Context
		email            string
		verificationCode string
		id               int64
		rep              *repository.ProfileRepositoryMock
		err              string
	}

	cases := make([]testCase, 0)

	rep := &repository.ProfileRepositoryMock{}
	rep.On("GetOne", context.Background(), int64(1)).Return(&models.Profile{}, errors.New("error"))

	cases = append(cases, testCase{
		ctx:              context.Background(),
		email:            "",
		verificationCode: "",
		id:               1,
		rep:              rep,
		err:              "error",
	})

	profile := &models.Profile{
		VerificationCode: "123",
	}
	rep = &repository.ProfileRepositoryMock{}
	rep.On("GetOne", context.Background(), int64(1)).Return(profile, nil)

	cases = append(cases, testCase{
		ctx:              context.Background(),
		email:            "",
		verificationCode: "321",
		id:               1,
		rep:              rep,
		err:              "wrong verification code",
	})

	profile = &models.Profile{
		VerificationCode: "123",
	}
	rep = &repository.ProfileRepositoryMock{}
	rep.On("GetOne", context.Background(), int64(1)).Return(profile, nil)
	rep.On("Update", context.Background(), profile).Return(errors.New("error1"))

	cases = append(cases, testCase{
		ctx:              context.Background(),
		email:            "email",
		verificationCode: "123",
		id:               1,
		rep:              rep,
		err:              "error1",
	})

	profile = &models.Profile{
		VerificationCode: "123",
	}
	rep = &repository.ProfileRepositoryMock{}
	rep.On("GetOne", context.Background(), int64(1)).Return(profile, nil)
	rep.On("Update", context.Background(), profile).Return(nil)

	cases = append(cases, testCase{
		ctx:              context.Background(),
		email:            "email",
		verificationCode: "123",
		id:               1,
		rep:              rep,
		err:              "",
	})

	for i, el := range cases {
		interactor := NewAdminInteractor(nil, nil, el.rep, nil)
		err := interactor.EditEmail(el.ctx, el.email, el.verificationCode, el.id)
		if err != nil && err.Error() != el.err {
			t.Error("errors don't match: ", i, err)
		}

		if el.rep != nil {
			el.rep.AssertExpectations(t)
		}
	}
}
