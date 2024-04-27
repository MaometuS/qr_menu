package admin_interactor

import (
	"context"
	"errors"
	"gitlab.com/maometusu/qr_menu/internal/entity/models"
	"gitlab.com/maometusu/qr_menu/internal/use_case/presenter"
	"gitlab.com/maometusu/qr_menu/internal/use_case/repository"
	"testing"
)

func TestAdminInteractor_AdminPage(t *testing.T) {
	type testCase struct {
		ctx        context.Context
		id         int64
		err        string
		presenter  *presenter.AdminPresenterMock
		repository *repository.ProfileRepositoryMock
	}

	cases := make([]testCase, 0)

	rep := &repository.ProfileRepositoryMock{}
	rep.On("GetOne", context.Background(), int64(1)).Return(&models.Profile{}, errors.New("error"))

	cases = append(cases, testCase{
		ctx:        context.Background(),
		id:         1,
		err:        "error",
		presenter:  nil,
		repository: rep,
	})

	prof := &models.Profile{
		ID: 1,
	}
	rep = &repository.ProfileRepositoryMock{}
	rep.On("GetOne", context.Background(), int64(1)).Return(prof, nil)
	pres := &presenter.AdminPresenterMock{}
	pres.On("AdminPage", nil, prof).Return(errors.New("error2"))

	cases = append(cases, testCase{
		ctx:        context.Background(),
		id:         1,
		err:        "error2",
		presenter:  pres,
		repository: rep,
	})

	prof = &models.Profile{
		ID: 1,
	}
	rep = &repository.ProfileRepositoryMock{}
	rep.On("GetOne", context.Background(), int64(1)).Return(prof, nil)
	pres = &presenter.AdminPresenterMock{}
	pres.On("AdminPage", nil, prof).Return(nil)

	cases = append(cases, testCase{
		ctx:        context.Background(),
		id:         1,
		err:        "",
		presenter:  pres,
		repository: rep,
	})

	for i, el := range cases {
		interactor := NewAdminInteractor(el.presenter, nil, el.repository, nil)
		err := interactor.AdminPage(el.ctx, nil, el.id)
		if err != nil && err.Error() != el.err {
			t.Error("errors don't match: ", i, err)
		}

		if el.repository != nil {
			el.repository.AssertExpectations(t)
		}

		if el.presenter != nil {
			el.presenter.AssertExpectations(t)
		}
	}
}
