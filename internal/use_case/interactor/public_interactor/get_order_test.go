package public_interactor

import (
	"context"
	"errors"
	"gitlab.com/maometusu/qr_menu/internal/entity/models"
	"gitlab.com/maometusu/qr_menu/internal/use_case/presenter"
	"gitlab.com/maometusu/qr_menu/internal/use_case/repository"
	"testing"
)

func TestPublicInteractor_GetOrder(t *testing.T) {
	type testCase struct {
		menuID        int64
		selectedItems map[int64][]int64
		menuRep       *repository.MenuRepositoryMock
		itemRep       *repository.ItemRepositoryMock
		pres          *presenter.PublicPresenterMock
		err           string
	}

	cases := make([]testCase, 0)

	menuRep := &repository.MenuRepositoryMock{}
	menuRep.On("GetOne", context.Background(), int64(1)).Return(&models.Menu{}, errors.New("error"))

	cases = append(cases, testCase{
		menuID:        1,
		selectedItems: nil,
		menuRep:       menuRep,
		itemRep:       nil,
		pres:          nil,
		err:           "error",
	})

	menuRep = &repository.MenuRepositoryMock{}
	menuRep.On("GetOne", context.Background(), int64(1)).Return(&models.Menu{EstablishmentID: 2}, nil)
	itemRep := &repository.ItemRepositoryMock{}
	itemRep.On("GetMultiple", context.Background(), []int64{3}).Return(make([]models.Item, 0), errors.New("error1"))

	cases = append(cases, testCase{
		menuID: 1,
		selectedItems: map[int64][]int64{
			2: {3, 3, 3, 3},
		},
		menuRep: menuRep,
		itemRep: itemRep,
		pres:    nil,
		err:     "error1",
	})

	items := []models.Item{
		{ID: 3, SelectedAmount: 4, Price: 100},
	}
	menuRep = &repository.MenuRepositoryMock{}
	menuRep.On("GetOne", context.Background(), int64(1)).Return(&models.Menu{EstablishmentID: 2}, nil)
	itemRep = &repository.ItemRepositoryMock{}
	itemRep.On("GetMultiple", context.Background(), []int64{3}).Return(items, nil)
	pres := &presenter.PublicPresenterMock{}
	pres.On("PresentOrder", nil, items, float64(400), int64(1)).Return(errors.New("error2"))

	cases = append(cases, testCase{
		menuID: 1,
		selectedItems: map[int64][]int64{
			2: {3, 3, 3, 3},
		},
		menuRep: menuRep,
		itemRep: itemRep,
		pres:    pres,
		err:     "error2",
	})

	items = []models.Item{
		{ID: 3, SelectedAmount: 4, Price: 100},
	}
	menuRep = &repository.MenuRepositoryMock{}
	menuRep.On("GetOne", context.Background(), int64(1)).Return(&models.Menu{EstablishmentID: 2}, nil)
	itemRep = &repository.ItemRepositoryMock{}
	itemRep.On("GetMultiple", context.Background(), []int64{3}).Return(items, nil)
	pres = &presenter.PublicPresenterMock{}
	pres.On("PresentOrder", nil, items, float64(400), int64(1)).Return(nil)

	cases = append(cases, testCase{
		menuID: 1,
		selectedItems: map[int64][]int64{
			2: {3, 3, 3, 3},
		},
		menuRep: menuRep,
		itemRep: itemRep,
		pres:    pres,
		err:     "",
	})

	for i, el := range cases {
		interactor := NewPublicInteractor(el.pres, nil, el.menuRep, nil, el.itemRep)
		err := interactor.GetOrder(context.Background(), nil, el.menuID, el.selectedItems)

		if err != nil && err.Error() != el.err {
			t.Error("errors don't match: ", i)
		}

		if el.itemRep != nil {
			el.itemRep.AssertExpectations(t)
		}

		if el.menuRep != nil {
			el.menuRep.AssertExpectations(t)
		}

		if el.pres != nil {
			el.pres.AssertExpectations(t)
		}
	}
}
