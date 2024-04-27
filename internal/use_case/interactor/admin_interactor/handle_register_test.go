package admin_interactor

import (
	"context"
	"errors"
	"github.com/stretchr/testify/mock"
	"gitlab.com/maometusu/qr_menu/internal/entity/models"
	"gitlab.com/maometusu/qr_menu/internal/use_case/presenter"
	"gitlab.com/maometusu/qr_menu/internal/use_case/repository"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestAdminInteractor_HandleRegister(t *testing.T) {
	type testCase struct {
		ctx        context.Context
		name       string
		email      string
		password   string
		passRepeat string
		rep        *repository.ProfileRepositoryMock
		pres       *presenter.AdminPresenterMock
		userID     int64
		err        string
	}
	cases := make([]testCase, 0)

	rep := &repository.ProfileRepositoryMock{}
	rep.On("CheckExistence", context.Background(), "email").Return(false, errors.New("error"))
	pres := &presenter.AdminPresenterMock{}
	pres.On("RegisterPage", nil, false, false, true).Return(nil)

	cases = append(cases, testCase{
		ctx:        context.Background(),
		name:       "name",
		email:      "email",
		password:   "password",
		passRepeat: "password",
		rep:        rep,
		pres:       pres,
		userID:     0,
		err:        "error",
	})

	rep = &repository.ProfileRepositoryMock{}
	rep.On("CheckExistence", context.Background(), "email").Return(true, nil)
	pres = &presenter.AdminPresenterMock{}
	pres.On("RegisterPage", nil, true, false, false).Return(nil)

	cases = append(cases, testCase{
		ctx:        context.Background(),
		name:       "name",
		email:      "email",
		password:   "password",
		passRepeat: "password",
		rep:        rep,
		pres:       pres,
		userID:     0,
		err:        "email already exists",
	})

	rep = &repository.ProfileRepositoryMock{}
	rep.On("CheckExistence", context.Background(), "email").Return(false, nil)
	pres = &presenter.AdminPresenterMock{}
	pres.On("RegisterPage", nil, false, true, false).Return(nil)

	cases = append(cases, testCase{
		ctx:        context.Background(),
		name:       "name",
		email:      "email",
		password:   "password",
		passRepeat: "password1",
		rep:        rep,
		pres:       pres,
		userID:     0,
		err:        "passwords don't match",
	})

	rep = &repository.ProfileRepositoryMock{}
	rep.On("CheckExistence", context.Background(), "email").Return(false, nil)
	rep.On("Create", context.Background(), mock.MatchedBy(func(profile *models.Profile) bool {
		if profile == nil {
			return false
		}

		if profile.Name != "name" || profile.Email != "email" {
			return false
		}

		err := bcrypt.CompareHashAndPassword([]byte(profile.Password), []byte("password"))
		if err != nil {
			return false
		}

		return true
	})).Return(0, errors.New("error1"))
	pres = &presenter.AdminPresenterMock{}
	pres.On("RegisterPage", nil, false, false, true).Return(nil)

	cases = append(cases, testCase{
		ctx:        context.Background(),
		name:       "name",
		email:      "email",
		password:   "password",
		passRepeat: "password",
		rep:        rep,
		pres:       pres,
		userID:     0,
		err:        "error1",
	})

	rep = &repository.ProfileRepositoryMock{}
	rep.On("CheckExistence", context.Background(), "email").Return(false, nil)
	rep.On("Create", context.Background(), mock.MatchedBy(func(profile *models.Profile) bool {
		if profile == nil {
			return false
		}

		if profile.Name != "name" || profile.Email != "email" {
			return false
		}

		err := bcrypt.CompareHashAndPassword([]byte(profile.Password), []byte("password"))
		if err != nil {
			return false
		}

		return true
	})).Return(1, nil)
	pres = &presenter.AdminPresenterMock{}

	cases = append(cases, testCase{
		ctx:        context.Background(),
		name:       "name",
		email:      "email",
		password:   "password",
		passRepeat: "password",
		rep:        rep,
		pres:       pres,
		userID:     1,
		err:        "error1",
	})

	for i, el := range cases {
		interactor := NewAdminInteractor(el.pres, nil, el.rep, nil)
		res, err := interactor.HandleRegister(el.ctx, nil, el.name, el.email, el.password, el.passRepeat)
		if err != nil && err.Error() != el.err {
			t.Error("errors don't match: ", i, err)
		}

		if res != el.userID {
			t.Error("results don't match: ", i, "expected", el.userID, "got", res)
		}

		if el.rep != nil {
			el.rep.AssertExpectations(t)
		}

		if el.pres != nil {
			el.pres.AssertExpectations(t)
		}
	}
}
