package admin_interactor

import (
	"context"
	"errors"
	"gitlab.com/maometusu/qr_menu/internal/entity/models"
	"gitlab.com/maometusu/qr_menu/internal/use_case/repository"
	"testing"
)

func TestAdminInteractor_ChangePassword(t *testing.T) {
	type testCase struct {
		ctx        context.Context
		password   string
		passRepeat string
		id         int64
		rep        *repository.ProfileRepositoryMock
		err        string
	}

	cases := []testCase{
		{
			ctx:        context.Background(),
			password:   "123",
			passRepeat: "321",
			id:         0,
			rep:        nil,
			err:        "passwords don't match",
		},
	}

	rep := &repository.ProfileRepositoryMock{}
	rep.On("GetOne", context.Background(), int64(1)).Return(&models.Profile{}, errors.New("error"))

	cases = append(cases, testCase{
		ctx:        context.Background(),
		password:   "123",
		passRepeat: "123",
		id:         1,
		rep:        rep,
		err:        "error",
	})

	profile := &models.Profile{}
	rep = &repository.ProfileRepositoryMock{}
	rep.On("GetOne", context.Background(), int64(1)).Return(profile, nil)
	rep.On("Update", context.Background(), profile).Return(errors.New("error1"))

	cases = append(cases, testCase{
		ctx:        context.Background(),
		password:   "123",
		passRepeat: "123",
		id:         1,
		rep:        rep,
		err:        "error1",
	})

	profile = &models.Profile{}
	rep = &repository.ProfileRepositoryMock{}
	rep.On("GetOne", context.Background(), int64(1)).Return(profile, nil)
	rep.On("Update", context.Background(), profile).Return(nil)

	cases = append(cases, testCase{
		ctx:        context.Background(),
		password:   "123",
		passRepeat: "123",
		id:         1,
		rep:        rep,
		err:        "",
	})

	for i, el := range cases {
		interactor := NewAdminInteractor(nil, nil, el.rep, nil)
		err := interactor.ChangePassword(el.ctx, el.password, el.passRepeat, el.id)
		if err != nil && err.Error() != el.err {
			t.Error("errors don't match: ", i, err)
		}

		if el.rep != nil {
			el.rep.AssertExpectations(t)
		}
	}
}
