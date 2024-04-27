package admin_interactor

import (
	"context"
	"errors"
	"gitlab.com/maometusu/qr_menu/internal/entity/models"
	"gitlab.com/maometusu/qr_menu/internal/use_case/presenter"
	"gitlab.com/maometusu/qr_menu/internal/use_case/repository"
	"testing"
)

func TestAdminInteractor_ProfilePage(t *testing.T) {
	type testCase struct {
		id   int64
		rep  *repository.ProfileRepositoryMock
		pres *presenter.AdminPresenterMock
		err  string
	}

	cases := make([]testCase, 0)

	rep := &repository.ProfileRepositoryMock{}
	rep.On("GetOne", context.Background(), int64(1)).Return(&models.Profile{}, errors.New("error"))

	cases = append(cases, testCase{
		id:   1,
		rep:  rep,
		pres: nil,
		err:  "error",
	})

	profile := &models.Profile{}
	rep = &repository.ProfileRepositoryMock{}
	rep.On("GetOne", context.Background(), int64(1)).Return(profile, nil)
	pres := &presenter.AdminPresenterMock{}
	pres.On("ProfilePage", nil, profile).Return(errors.New("error1"))

	cases = append(cases, testCase{
		id:   1,
		rep:  rep,
		pres: pres,
		err:  "error1",
	})

	profile = &models.Profile{}
	rep = &repository.ProfileRepositoryMock{}
	rep.On("GetOne", context.Background(), int64(1)).Return(profile, nil)
	pres = &presenter.AdminPresenterMock{}
	pres.On("ProfilePage", nil, profile).Return(nil)

	cases = append(cases, testCase{
		id:   1,
		rep:  rep,
		pres: pres,
		err:  "",
	})

	for i, el := range cases {
		interactor := NewAdminInteractor(el.pres, nil, el.rep, nil)
		err := interactor.ProfilePage(context.Background(), nil, el.id)
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
