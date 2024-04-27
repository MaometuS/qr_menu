package menu_interactor

import (
	"context"
	"errors"
	"gitlab.com/maometusu/qr_menu/internal/entity/models"
	"gitlab.com/maometusu/qr_menu/internal/use_case/repository"
	"testing"
)

func TestInteractor_DeleteMenu(t *testing.T) {
	type testCase struct {
		profileID int64
		menuID    int64
		estRep    *repository.EstablishmentRepositoryMock
		rep       *repository.MenuRepositoryMock
		err       string
	}

	cases := make([]testCase, 0)

	rep := &repository.MenuRepositoryMock{}
	rep.On("GetOne", context.Background(), int64(1)).Return(&models.Menu{}, errors.New("error"))

	cases = append(cases, testCase{
		profileID: 2,
		menuID:    1,
		estRep:    nil,
		rep:       rep,
		err:       "error",
	})

	rep = &repository.MenuRepositoryMock{}
	rep.On("GetOne", context.Background(), int64(1)).Return(&models.Menu{EstablishmentID: 3}, nil)
	estRep := &repository.EstablishmentRepositoryMock{}
	estRep.On("GetOne", context.Background(), int64(3)).Return(&models.Establishment{}, errors.New("error1"))

	cases = append(cases, testCase{
		profileID: 2,
		menuID:    1,
		estRep:    estRep,
		rep:       rep,
		err:       "error1",
	})

	rep = &repository.MenuRepositoryMock{}
	rep.On("GetOne", context.Background(), int64(1)).Return(&models.Menu{EstablishmentID: 3}, nil)
	estRep = &repository.EstablishmentRepositoryMock{}
	estRep.On("GetOne", context.Background(), int64(3)).Return(&models.Establishment{ProfileID: 2}, nil)
	rep.On("DeleteMenu", context.Background(), int64(1)).Return(errors.New("error2"))

	cases = append(cases, testCase{
		profileID: 2,
		menuID:    1,
		estRep:    estRep,
		rep:       rep,
		err:       "error2",
	})

	rep = &repository.MenuRepositoryMock{}
	rep.On("GetOne", context.Background(), int64(1)).Return(&models.Menu{EstablishmentID: 3}, nil)
	estRep = &repository.EstablishmentRepositoryMock{}
	estRep.On("GetOne", context.Background(), int64(3)).Return(&models.Establishment{ProfileID: 2}, nil)
	rep.On("DeleteMenu", context.Background(), int64(1)).Return(nil)

	cases = append(cases, testCase{
		profileID: 2,
		menuID:    1,
		estRep:    estRep,
		rep:       rep,
		err:       "",
	})

	for i, el := range cases {
		interactor := NewMenuInteractor(nil, el.rep, el.estRep)
		err := interactor.DeleteMenu(context.Background(), el.profileID, el.menuID)
		if err != nil && err.Error() != el.err {
			t.Error("errors don't match: ", i, err)
		}

		if el.rep != nil {
			el.rep.AssertExpectations(t)
		}

		if el.estRep != nil {
			el.estRep.AssertExpectations(t)
		}
	}
}
