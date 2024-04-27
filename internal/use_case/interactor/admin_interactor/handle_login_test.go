package admin_interactor

import (
	"context"
	"errors"
	"gitlab.com/maometusu/qr_menu/internal/entity"
	"gitlab.com/maometusu/qr_menu/internal/entity/models"
	"gitlab.com/maometusu/qr_menu/internal/use_case/presenter"
	"gitlab.com/maometusu/qr_menu/internal/use_case/repository"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestAdminInteractor_HandleLogin(t *testing.T) {
	type testCase struct {
		ctx       context.Context
		email     string
		password  string
		rep       *repository.ProfileRepositoryMock
		pres      *presenter.AdminPresenterMock
		res       string
		jwtSecret string
		err       string
	}
	cases := make([]testCase, 0)

	rep := &repository.ProfileRepositoryMock{}
	rep.On("GetOneByEmail", context.Background(), "email").Return(&models.Profile{}, errors.New("error"))
	pres := &presenter.AdminPresenterMock{}
	pres.On("LoginPage", nil, true, false).Return(nil)

	cases = append(cases, testCase{
		ctx:      context.Background(),
		email:    "email",
		password: "password",
		rep:      rep,
		pres:     pres,
		res:      "",
		err:      "error",
	})

	profile := &models.Profile{
		Password: "$2a$06$mAPJjRvk89Iz2KH7yRUXx.z7QSYEnHf8ONWHuPBE2uJGMJKlIrCXu",
	}
	rep = &repository.ProfileRepositoryMock{}
	rep.On("GetOneByEmail", context.Background(), "email").Return(profile, nil)
	pres = &presenter.AdminPresenterMock{}
	pres.On("LoginPage", nil, true, false).Return(nil)

	cases = append(cases, testCase{
		ctx:      context.Background(),
		email:    "email",
		password: "password",
		rep:      rep,
		pres:     pres,
		res:      "",
		err:      "crypto/bcrypt: hashedPassword is not the hash of the given password",
	})

	pass, err := bcrypt.GenerateFromPassword([]byte("password"), 6)
	if err != nil {
		t.Error(err)
	}

	profile = &models.Profile{
		ID:       1,
		Password: string(pass),
	}
	rep = &repository.ProfileRepositoryMock{}
	rep.On("GetOneByEmail", context.Background(), "email").Return(profile, nil)
	pres = &presenter.AdminPresenterMock{}

	cases = append(cases, testCase{
		ctx:       context.Background(),
		email:     "email",
		password:  "password",
		rep:       rep,
		pres:      pres,
		res:       "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJxcm1lbnUudXoiLCJzdWIiOjF9.t6jTULrdQ16MaRFRXS_yWcBFDUZ4kWJ2AJkMuChm4lU",
		jwtSecret: "secret",
		err:       "some error",
	})

	for i, el := range cases {
		interactor := NewAdminInteractor(el.pres, nil, el.rep, &entity.Config{JWTSignString: el.jwtSecret})
		res, err := interactor.HandleLogin(el.ctx, nil, el.email, el.password)
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
