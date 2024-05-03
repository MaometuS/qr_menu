package admin_interactor

import (
	"context"
	"errors"
	"gitlab.com/maometusu/qr_menu/internal/entity/models"
	"gitlab.com/maometusu/qr_menu/internal/use_case/presenter"
	"gitlab.com/maometusu/qr_menu/internal/use_case/repository"
	"testing"
)

func TestAdminInteractor_VerifyRestorePassword(t *testing.T) {
	type testCase struct {
		id               int64
		verificationCode string
		rep              *repository.ProfileRepositoryMock
		pres             *presenter.AdminPresenterMock
		err              string
	}

	cases := make([]testCase, 0)

	rep := repository.NewProfileRepositoryMock()
	rep.On("GetOne", context.Background(), int64(1)).Return(&models.Profile{}, errors.New("error"))
	pres := presenter.NewAdminPresenterMock()
	pres.On("VerifyRestorePasswordPage", nil, int64(1), false, true).Return(nil)

	cases = append(cases, testCase{
		id:               1,
		verificationCode: "",
		rep:              rep,
		pres:             pres,
		err:              "error",
	})

	prof := &models.Profile{VerificationCode: "code"}
	rep = repository.NewProfileRepositoryMock()
	rep.On("GetOne", context.Background(), int64(1)).Return(prof, nil)
	pres = presenter.NewAdminPresenterMock()
	pres.On("VerifyRestorePasswordPage", nil, int64(1), true, false).Return(nil)

	cases = append(cases, testCase{
		id:               1,
		verificationCode: "code1",
		rep:              rep,
		pres:             pres,
		err:              "codes don't match",
	})

	prof = &models.Profile{VerificationCode: "code"}
	rep = repository.NewProfileRepositoryMock()
	rep.On("GetOne", context.Background(), int64(1)).Return(prof, nil)
	rep.On("Update", context.Background(), prof).Return(errors.New("error1"))
	pres = presenter.NewAdminPresenterMock()
	pres.On("VerifyRestorePasswordPage", nil, int64(1), false, true).Return(nil)

	cases = append(cases, testCase{
		id:               1,
		verificationCode: "code",
		rep:              rep,
		pres:             pres,
		err:              "error1",
	})

	prof = &models.Profile{VerificationCode: "code"}
	rep = repository.NewProfileRepositoryMock()
	rep.On("GetOne", context.Background(), int64(1)).Return(prof, nil)
	rep.On("Update", context.Background(), prof).Return(nil)

	cases = append(cases, testCase{
		id:               1,
		verificationCode: "code",
		rep:              rep,
		pres:             pres,
		err:              "",
	})

	for i, el := range cases {
		interactor := NewAdminInteractor(el.pres, nil, el.rep, nil)
		err := interactor.VerifyRestorePassword(context.Background(), nil, el.id, el.verificationCode)
		if err != nil && err.Error() != el.err {
			t.Error("errors don't match: ", i, err)
		}

		if el.rep != nil {
			el.rep.AssertExpectations(t)
		}

		if el.pres != nil {
			el.pres.AssertExpectations(t)
		}
	}
}
