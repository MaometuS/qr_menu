package menu_interactor

import (
	"context"
	"errors"
	"gitlab.com/maometusu/qr_menu/internal/entity/models"
	"gitlab.com/maometusu/qr_menu/internal/use_case/repository"
	"testing"
)

func TestInteractor_EditMenu(t *testing.T) {
	type testCase struct {
		profileID int64
		menu      *models.Menu
		estRep    *repository.EstablishmentRepositoryMock
		rep       *repository.MenuRepositoryMock
		err       string
	}

	cases := make([]testCase, 0)

	estRep := &repository.EstablishmentRepositoryMock{}
	estRep.On("GetOne", context.Background(), int64(1)).Return(&models.Establishment{}, errors.New("error"))

	cases = append(cases, testCase{
		profileID: 0,
		menu:      &models.Menu{EstablishmentID: 1},
		estRep:    estRep,
		rep:       nil,
		err:       "error",
	})

	establishment := &models.Establishment{
		ProfileID: 2,
	}
	estRep = &repository.EstablishmentRepositoryMock{}
	estRep.On("GetOne", context.Background(), int64(1)).Return(establishment, nil)

	cases = append(cases, testCase{
		profileID: 0,
		menu:      &models.Menu{EstablishmentID: 1},
		estRep:    estRep,
		rep:       nil,
		err:       "not allowed",
	})

	establishment = &models.Establishment{
		ProfileID: 2,
	}
	menu := &models.Menu{
		EstablishmentID: 1,
	}
	estRep = &repository.EstablishmentRepositoryMock{}
	estRep.On("GetOne", context.Background(), int64(1)).Return(establishment, nil)
	rep := &repository.MenuRepositoryMock{}
	rep.On("UpdateMenu", context.Background(), menu).Return(errors.New("error1"))

	cases = append(cases, testCase{
		profileID: 2,
		menu:      menu,
		estRep:    estRep,
		rep:       rep,
		err:       "error1",
	})

	establishment = &models.Establishment{
		ProfileID: 2,
	}
	menu = &models.Menu{
		EstablishmentID: 1,
	}
	estRep = &repository.EstablishmentRepositoryMock{}
	estRep.On("GetOne", context.Background(), int64(1)).Return(establishment, nil)
	rep = &repository.MenuRepositoryMock{}
	rep.On("UpdateMenu", context.Background(), menu).Return(nil)

	cases = append(cases, testCase{
		profileID: 2,
		menu:      menu,
		estRep:    estRep,
		rep:       rep,
		err:       "",
	})

	for i, el := range cases {
		interactor := NewMenuInteractor(nil, el.rep, el.estRep)
		err := interactor.EditMenu(context.Background(), el.profileID, el.menu)
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
