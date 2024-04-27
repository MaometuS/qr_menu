package menu_interactor

import (
	"context"
	"errors"
	"gitlab.com/maometusu/qr_menu/internal/entity/models"
	"gitlab.com/maometusu/qr_menu/internal/use_case/presenter"
	"gitlab.com/maometusu/qr_menu/internal/use_case/repository"
	"testing"
)

func TestInteractor_GetMenus(t *testing.T) {
	type testCase struct {
		profileID       int64
		establishmentID int64
		estRep          *repository.EstablishmentRepositoryMock
		rep             *repository.MenuRepositoryMock
		pres            *presenter.MenuPresenterMock
		err             string
	}

	cases := make([]testCase, 0)

	estRep := &repository.EstablishmentRepositoryMock{}
	estRep.On("GetOne", context.Background(), int64(1)).Return(&models.Establishment{}, errors.New("error"))

	cases = append(cases, testCase{
		profileID:       2,
		establishmentID: 1,
		estRep:          estRep,
		rep:             nil,
		pres:            nil,
		err:             "error",
	})

	estRep = &repository.EstablishmentRepositoryMock{}
	estRep.On("GetOne", context.Background(), int64(1)).Return(&models.Establishment{ProfileID: 3}, nil)

	cases = append(cases, testCase{
		profileID:       2,
		establishmentID: 1,
		estRep:          estRep,
		rep:             nil,
		pres:            nil,
		err:             "not allowed",
	})

	estRep = &repository.EstablishmentRepositoryMock{}
	estRep.On("GetOne", context.Background(), int64(1)).Return(&models.Establishment{ProfileID: 2}, nil)
	rep := &repository.MenuRepositoryMock{}
	rep.On("GetAll", context.Background(), int64(1)).Return(make([]models.Menu, 0), errors.New("error1"))

	cases = append(cases, testCase{
		profileID:       2,
		establishmentID: 1,
		estRep:          estRep,
		rep:             rep,
		pres:            nil,
		err:             "error1",
	})

	menus := make([]models.Menu, 0)
	estRep = &repository.EstablishmentRepositoryMock{}
	estRep.On("GetOne", context.Background(), int64(1)).Return(&models.Establishment{ProfileID: 2}, nil)
	rep = &repository.MenuRepositoryMock{}
	rep.On("GetAll", context.Background(), int64(1)).Return(menus, nil)
	pres := &presenter.MenuPresenterMock{}
	pres.On("PresentMenus", nil, menus, int64(1)).Return(errors.New("error2"))

	cases = append(cases, testCase{
		profileID:       2,
		establishmentID: 1,
		estRep:          estRep,
		rep:             rep,
		pres:            pres,
		err:             "error2",
	})

	menus = make([]models.Menu, 0)
	estRep = &repository.EstablishmentRepositoryMock{}
	estRep.On("GetOne", context.Background(), int64(1)).Return(&models.Establishment{ProfileID: 2}, nil)
	rep = &repository.MenuRepositoryMock{}
	rep.On("GetAll", context.Background(), int64(1)).Return(menus, nil)
	pres = &presenter.MenuPresenterMock{}
	pres.On("PresentMenus", nil, menus, int64(1)).Return(nil)

	cases = append(cases, testCase{
		profileID:       2,
		establishmentID: 1,
		estRep:          estRep,
		rep:             rep,
		pres:            pres,
		err:             "",
	})

	for i, el := range cases {
		interactor := NewMenuInteractor(el.pres, el.rep, el.estRep)
		err := interactor.GetMenus(context.Background(), nil, el.profileID, el.establishmentID)
		if err != nil && err.Error() != el.err {
			t.Error("errors don't match: ", i, err)
		}

		if el.rep != nil {
			el.rep.AssertExpectations(t)
		}

		if el.estRep != nil {
			el.estRep.AssertExpectations(t)
		}

		if el.pres != nil {
			el.pres.AssertExpectations(t)
		}
	}
}
