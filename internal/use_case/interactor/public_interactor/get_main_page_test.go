package public_interactor

import (
	"context"
	"errors"
	"gitlab.com/maometusu/qr_menu/internal/entity/models"
	"gitlab.com/maometusu/qr_menu/internal/use_case/presenter"
	"gitlab.com/maometusu/qr_menu/internal/use_case/repository"
	"testing"
)

func TestPublicInteractor_GetMainPage(t *testing.T) {
	type testCase struct {
		link          string
		selectedItems map[int64][]int64
		estRep        *repository.EstablishmentRepositoryMock
		menuRep       *repository.MenuRepositoryMock
		pres          *presenter.PublicPresenterMock
		err           string
	}

	cases := make([]testCase, 0)

	estRep := &repository.EstablishmentRepositoryMock{}
	estRep.On("GetByLink", context.Background(), "link").Return(&models.Establishment{}, errors.New("error"))

	cases = append(cases, testCase{
		link:          "link",
		selectedItems: nil,
		estRep:        estRep,
		menuRep:       nil,
		pres:          nil,
		err:           "error",
	})

	estRep = &repository.EstablishmentRepositoryMock{}
	estRep.On("GetByLink", context.Background(), "link").Return(&models.Establishment{ID: 1}, nil)
	menuRep := &repository.MenuRepositoryMock{}
	menuRep.On("GetAll", context.Background(), int64(1)).Return(make([]models.Menu, 0), errors.New("error1"))

	cases = append(cases, testCase{
		link:          "link",
		selectedItems: nil,
		estRep:        estRep,
		menuRep:       menuRep,
		pres:          nil,
		err:           "error1",
	})

	establishment := &models.Establishment{ID: 1}
	menus := []models.Menu{
		{ID: 3, IsVisible: true},
	}
	estRep = &repository.EstablishmentRepositoryMock{}
	estRep.On("GetByLink", context.Background(), "link").Return(establishment, nil)
	menuRep = &repository.MenuRepositoryMock{}
	menuRep.On("GetAll", context.Background(), int64(1)).Return(menus, nil)
	pres := &presenter.PublicPresenterMock{}
	pres.On("PresentMainPage", nil, establishment, menus, int64(3), true).Return(errors.New("error2"))

	cases = append(cases, testCase{
		link: "link",
		selectedItems: map[int64][]int64{
			1: {1, 2, 3},
		},
		estRep:  estRep,
		menuRep: menuRep,
		pres:    pres,
		err:     "error2",
	})

	establishment = &models.Establishment{ID: 1}
	menus = []models.Menu{
		{ID: 3, IsVisible: true},
	}
	estRep = &repository.EstablishmentRepositoryMock{}
	estRep.On("GetByLink", context.Background(), "link").Return(establishment, nil)
	menuRep = &repository.MenuRepositoryMock{}
	menuRep.On("GetAll", context.Background(), int64(1)).Return(menus, nil)
	pres = &presenter.PublicPresenterMock{}
	pres.On("PresentMainPage", nil, establishment, menus, int64(3), true).Return(nil)

	cases = append(cases, testCase{
		link: "link",
		selectedItems: map[int64][]int64{
			1: {1, 2, 3},
		},
		estRep:  estRep,
		menuRep: menuRep,
		pres:    pres,
		err:     "",
	})

	for i, el := range cases {
		interactor := NewPublicInteractor(el.pres, el.estRep, el.menuRep, nil, nil)
		err := interactor.GetMainPage(context.Background(), nil, el.link, el.selectedItems)

		if err != nil && err.Error() != el.err {
			t.Error("errors don't match: ", i)
		}

		if el.estRep != nil {
			el.estRep.AssertExpectations(t)
		}

		if el.menuRep != nil {
			el.menuRep.AssertExpectations(t)
		}

		if el.pres != nil {
			el.pres.AssertExpectations(t)
		}
	}
}
