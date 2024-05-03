package admin_interactor

import (
	"context"
	"errors"
	"gitlab.com/maometusu/qr_menu/internal/entity/models"
	"gitlab.com/maometusu/qr_menu/internal/use_case/presenter"
	"gitlab.com/maometusu/qr_menu/internal/use_case/repository"
	"testing"
)

func TestAdminInteractor_RestorePassword(t *testing.T) {
	type testCase struct {
		email      string
		pass       string
		passRepeat string
		pres       *presenter.AdminPresenterMock
		rep        *repository.ProfileRepositoryMock
		res        int64
		err        string
	}

	cases := make([]testCase, 0)

	rep := repository.NewProfileRepositoryMock()
	rep.On("GetOneByEmail", context.Background(), "email").Return(&models.Profile{}, errors.New("error"))
	pres := presenter.NewAdminPresenterMock()
	pres.On("RestorePasswordPage", nil, true, false, false).Return(nil)

	cases = append(cases, testCase{
		email:      "email",
		pass:       "",
		passRepeat: "",
		pres:       pres,
		rep:        rep,
		res:        0,
		err:        "error",
	})

	rep = repository.NewProfileRepositoryMock()
	rep.On("GetOneByEmail", context.Background(), "email").Return(&models.Profile{}, nil)
	pres = presenter.NewAdminPresenterMock()
	pres.On("RestorePasswordPage", nil, false, true, false).Return(nil)

	cases = append(cases, testCase{
		email:      "email",
		pass:       "pass",
		passRepeat: "ssap",
		pres:       pres,
		rep:        rep,
		res:        0,
		err:        "passwords don't match",
	})

	prof := &models.Profile{}
	rep = repository.NewProfileRepositoryMock()
	rep.On("GetOneByEmail", context.Background(), "email").Return(prof, nil)
	rep.On("Update", context.Background(), prof).Return(errors.New("error1"))
	pres = presenter.NewAdminPresenterMock()
	pres.On("RestorePasswordPage", nil, false, false, true).Return(nil)

	cases = append(cases, testCase{
		email:      "email",
		pass:       "pass",
		passRepeat: "pass",
		pres:       pres,
		rep:        rep,
		res:        0,
		err:        "error1",
	})

	prof = &models.Profile{ID: 1}
	rep = repository.NewProfileRepositoryMock()
	rep.On("GetOneByEmail", context.Background(), "email").Return(prof, nil)
	rep.On("Update", context.Background(), prof).Return(nil)

	cases = append(cases, testCase{
		email:      "email",
		pass:       "pass",
		passRepeat: "pass",
		pres:       nil,
		rep:        rep,
		res:        1,
		err:        "",
	})

	for i, el := range cases {
		interactor := NewAdminInteractor(el.pres, nil, el.rep, nil)
		res, err := interactor.RestorePassword(context.Background(), nil, el.email, el.pass, el.passRepeat)
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
