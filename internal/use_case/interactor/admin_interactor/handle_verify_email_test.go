package admin_interactor

import (
	"context"
	"errors"
	"gitlab.com/maometusu/qr_menu/internal/entity"
	"gitlab.com/maometusu/qr_menu/internal/entity/models"
	"gitlab.com/maometusu/qr_menu/internal/use_case/presenter"
	"gitlab.com/maometusu/qr_menu/internal/use_case/repository"
	"testing"
)

func TestAdminInteractor_HandleVerifyEmail(t *testing.T) {
	type testCase struct {
		ctx              context.Context
		id               int64
		verificationCode string
		jwtSecret        string
		rep              *repository.ProfileRepositoryMock
		pres             *presenter.AdminPresenterMock
		res              string
		err              string
	}

	cases := make([]testCase, 0)

	rep := &repository.ProfileRepositoryMock{}
	rep.On("GetOne", context.Background(), int64(1)).Return(&models.Profile{}, errors.New("error"))
	pres := &presenter.AdminPresenterMock{}
	pres.On("VerifyPage", nil, int64(1), false, true).Return(nil)

	cases = append(cases, testCase{
		ctx:              context.Background(),
		id:               1,
		verificationCode: "",
		rep:              rep,
		pres:             pres,
		res:              "",
		err:              "error",
	})

	profile := &models.Profile{
		VerificationCode: "verification_code",
	}
	rep = &repository.ProfileRepositoryMock{}
	rep.On("GetOne", context.Background(), int64(1)).Return(profile, nil)
	pres = &presenter.AdminPresenterMock{}
	pres.On("VerifyPage", nil, int64(1), true, false).Return(nil)

	cases = append(cases, testCase{
		ctx:              context.Background(),
		id:               1,
		verificationCode: "code",
		rep:              rep,
		pres:             pres,
		res:              "",
		err:              "code does not match",
	})

	profile = &models.Profile{
		VerificationCode: "verification_code",
	}
	rep = &repository.ProfileRepositoryMock{}
	rep.On("GetOne", context.Background(), int64(1)).Return(profile, nil)
	rep.On("SetVerified", context.Background(), int64(1), true).Return(errors.New("error1"))
	pres = &presenter.AdminPresenterMock{}
	pres.On("VerifyPage", nil, int64(1), false, true).Return(nil)

	cases = append(cases, testCase{
		ctx:              context.Background(),
		id:               1,
		verificationCode: "verification_code",
		rep:              rep,
		pres:             pres,
		res:              "",
		err:              "error1",
	})

	profile = &models.Profile{
		VerificationCode: "verification_code",
	}
	rep = &repository.ProfileRepositoryMock{}
	rep.On("GetOne", context.Background(), int64(1)).Return(profile, nil)
	rep.On("SetVerified", context.Background(), int64(1), true).Return(nil)
	pres = &presenter.AdminPresenterMock{}

	cases = append(cases, testCase{
		ctx:              context.Background(),
		id:               1,
		verificationCode: "verification_code",
		rep:              rep,
		pres:             pres,
		res:              "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJxcm1lbnUudXoiLCJzdWIiOjF9.u002sjszD3945UZX3ncsxGQcNoOb24Kk6FmYv2qclLA",
		err:              "",
	})

	for i, el := range cases {
		interactor := NewAdminInteractor(el.pres, nil, el.rep, &entity.Config{JWTSignString: el.jwtSecret})
		res, err := interactor.HandleVerifyEmail(el.ctx, nil, el.id, el.verificationCode)
		if err != nil && err.Error() != el.err {
			t.Error("errors don't match: ", i, err)
		}

		if res != el.res {
			t.Error("results don't match: ", i, "expected", el.res, "got", res)
		}

		if el.rep != nil {
			el.rep.AssertExpectations(t)
		}

		if el.pres != nil {
			el.pres.AssertExpectations(t)
		}
	}
}
