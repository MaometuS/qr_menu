package public_interactor

import (
	"context"
	"errors"
	"gitlab.com/maometusu/qr_menu/internal/entity/models"
	"gitlab.com/maometusu/qr_menu/internal/use_case/presenter"
	"gitlab.com/maometusu/qr_menu/internal/use_case/repository"
	"testing"
)

func TestPublicInteractor_GetSearch(t *testing.T) {
	type testCase struct {
		search          string
		establishmentID int64
		menuID          int64
		selectedItems   map[int64][]int64
		itemRep         *repository.ItemRepositoryMock
		pres            *presenter.PublicPresenterMock
		err             string
	}

	cases := make([]testCase, 0)

	itemRep := &repository.ItemRepositoryMock{}
	itemRep.On("Search", context.Background(), "search", int64(1)).Return(make([]models.Item, 0), errors.New("error"))

	cases = append(cases, testCase{
		search:          "search",
		establishmentID: 1,
		menuID:          0,
		selectedItems:   nil,
		itemRep:         itemRep,
		pres:            nil,
		err:             "error",
	})

	items := []models.Item{
		{ID: 2, SelectedAmount: 3},
	}
	itemRep = &repository.ItemRepositoryMock{}
	itemRep.On("Search", context.Background(), "search", int64(1)).Return(items, nil)
	pres := &presenter.PublicPresenterMock{}
	pres.On("PresentSearch", nil, "search", items, int64(1), int64(5)).Return(errors.New("error1"))

	cases = append(cases, testCase{
		search:          "search",
		establishmentID: 1,
		menuID:          5,
		selectedItems: map[int64][]int64{
			1: {2, 2, 2},
		},
		itemRep: itemRep,
		pres:    pres,
		err:     "error1",
	})

	items = []models.Item{
		{ID: 2, SelectedAmount: 3},
	}
	itemRep = &repository.ItemRepositoryMock{}
	itemRep.On("Search", context.Background(), "search", int64(1)).Return(items, nil)
	pres = &presenter.PublicPresenterMock{}
	pres.On("PresentSearch", nil, "search", items, int64(1), int64(5)).Return(nil)

	cases = append(cases, testCase{
		search:          "search",
		establishmentID: 1,
		menuID:          5,
		selectedItems: map[int64][]int64{
			1: {2, 2, 2},
		},
		itemRep: itemRep,
		pres:    pres,
		err:     "",
	})

	for i, el := range cases {
		interactor := NewPublicInteractor(el.pres, nil, nil, nil, el.itemRep)
		err := interactor.GetSearch(context.Background(), nil, el.search, el.establishmentID, el.menuID, el.selectedItems)

		if err != nil && err.Error() != el.err {
			t.Error("errors don't match: ", i)
		}

		if el.itemRep != nil {
			el.itemRep.AssertExpectations(t)
		}

		if el.pres != nil {
			el.pres.AssertExpectations(t)
		}
	}
}
