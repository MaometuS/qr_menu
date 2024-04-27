package admin_interactor

import (
	"context"
	"errors"
	"gitlab.com/maometusu/qr_menu/internal/entity/models"
	"gitlab.com/maometusu/qr_menu/internal/use_case/repository"
	"testing"
)

func TestAdminInteractor_EditName(t *testing.T) {
	type testCase struct {
		ctx   context.Context
		email string
		name  string
		id    int64
		rep   *repository.ProfileRepositoryMock
		res   bool
		err   string
	}
	cases := make([]testCase, 0)

	rep := &repository.ProfileRepositoryMock{}
	rep.On("GetOne", context.Background(), int64(1)).Return(&models.Profile{}, errors.New("error"))

	cases = append(cases, testCase{
		ctx:   context.Background(),
		email: "",
		name:  "",
		id:    1,
		rep:   rep,
		res:   false,
		err:   "error",
	})

	profile := &models.Profile{}
	rep = &repository.ProfileRepositoryMock{}
	rep.On("GetOne", context.Background(), int64(1)).Return(profile, nil)
	rep.On("Update", context.Background(), profile).Return(errors.New("error1"))

	cases = append(cases, testCase{
		ctx:   context.Background(),
		email: "email",
		name:  "name",
		id:    1,
		rep:   rep,
		res:   false,
		err:   "error1",
	})

	profile = &models.Profile{
		Email: "email",
	}
	rep = &repository.ProfileRepositoryMock{}
	rep.On("GetOne", context.Background(), int64(1)).Return(profile, nil)
	rep.On("Update", context.Background(), profile).Return(nil)

	cases = append(cases, testCase{
		ctx:   context.Background(),
		email: "email",
		name:  "name",
		id:    1,
		rep:   rep,
		res:   false,
		err:   "",
	})

	profile = &models.Profile{
		Email: "email1",
	}
	rep = &repository.ProfileRepositoryMock{}
	rep.On("GetOne", context.Background(), int64(1)).Return(profile, nil)
	rep.On("Update", context.Background(), profile).Return(nil)

	cases = append(cases, testCase{
		ctx:   context.Background(),
		email: "email",
		name:  "name",
		id:    1,
		rep:   rep,
		res:   true,
		err:   "",
	})

	for i, el := range cases {
		interactor := NewAdminInteractor(nil, nil, el.rep, nil)
		res, err := interactor.EditName(el.ctx, el.email, el.name, el.id)
		if err != nil && err.Error() != el.err {
			t.Error("errors don't match: ", i, err)
		}

		if res != el.res {
			t.Error("results don't match: ", i, "expected", el.res, "got", res)
		}

		if el.rep != nil {
			el.rep.AssertExpectations(t)
		}
	}
}
