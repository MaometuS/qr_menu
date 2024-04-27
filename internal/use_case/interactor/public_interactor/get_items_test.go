package public_interactor

import (
	"context"
	"errors"
	"github.com/stretchr/testify/mock"
	"gitlab.com/maometusu/qr_menu/internal/entity/models"
	"gitlab.com/maometusu/qr_menu/internal/use_case/presenter"
	"gitlab.com/maometusu/qr_menu/internal/use_case/repository"
	"testing"
)

func TestPublicInteractor_GetItems(t *testing.T) {
	type testCase struct {
		categoryID    int64
		selectedItems map[int64][]int64
		catRep        *repository.CategoryRepositoryMock
		itemRep       *repository.ItemRepositoryMock
		pres          *presenter.PublicPresenterMock
		err           string
	}

	cases := make([]testCase, 0)

	catRep := &repository.CategoryRepositoryMock{}
	catRep.On("GetOne", context.Background(), int64(1)).Return(&models.Category{}, errors.New("error"))

	cases = append(cases, testCase{
		categoryID:    1,
		selectedItems: nil,
		catRep:        catRep,
		itemRep:       nil,
		pres:          nil,
		err:           "error",
	})

	catRep = &repository.CategoryRepositoryMock{}
	catRep.On("GetOne", context.Background(), int64(1)).Return(&models.Category{}, nil)
	itemRep := &repository.ItemRepositoryMock{}
	itemRep.On("GetAll", context.Background(), int64(1)).Return(make([]models.Item, 0), errors.New("error1"))

	cases = append(cases, testCase{
		categoryID:    1,
		selectedItems: nil,
		catRep:        catRep,
		itemRep:       itemRep,
		pres:          nil,
		err:           "error1",
	})

	category := &models.Category{
		Name:   "name",
		MenuID: 5,
	}
	items := []models.Item{
		{ID: 2},
	}
	catRep = &repository.CategoryRepositoryMock{}
	catRep.On("GetOne", context.Background(), int64(1)).Return(category, nil)
	itemRep = &repository.ItemRepositoryMock{}
	itemRep.On("GetAll", context.Background(), int64(1)).Return(items, nil)
	itemRep.On("GetEstablishment", context.Background(), int64(2)).Return(3, nil)
	pres := &presenter.PublicPresenterMock{}
	pres.On("PresentItems", nil, "name", mock.MatchedBy(func(i []models.Item) bool {
		return len(i) == 1 && i[0].ID == 2 && i[0].SelectedAmount == 4
	}), int64(5)).Return(errors.New("error2"))

	cases = append(cases, testCase{
		categoryID: 1,
		selectedItems: map[int64][]int64{
			3: {2, 2, 2, 2},
		},
		catRep:  catRep,
		itemRep: itemRep,
		pres:    pres,
		err:     "error2",
	})

	category = &models.Category{
		Name:   "name",
		MenuID: 5,
	}
	items = []models.Item{
		{ID: 2},
	}
	catRep = &repository.CategoryRepositoryMock{}
	catRep.On("GetOne", context.Background(), int64(1)).Return(category, nil)
	itemRep = &repository.ItemRepositoryMock{}
	itemRep.On("GetAll", context.Background(), int64(1)).Return(items, nil)
	itemRep.On("GetEstablishment", context.Background(), int64(2)).Return(3, nil)
	pres = &presenter.PublicPresenterMock{}
	pres.On("PresentItems", nil, "name", mock.MatchedBy(func(i []models.Item) bool {
		return len(i) == 1 && i[0].ID == 2 && i[0].SelectedAmount == 4
	}), int64(5)).Return(nil)

	cases = append(cases, testCase{
		categoryID: 1,
		selectedItems: map[int64][]int64{
			3: {2, 2, 2, 2},
		},
		catRep:  catRep,
		itemRep: itemRep,
		pres:    pres,
		err:     "",
	})

	for i, el := range cases {
		interactor := NewPublicInteractor(el.pres, nil, nil, el.catRep, el.itemRep)
		err := interactor.GetItems(context.Background(), nil, el.categoryID, el.selectedItems)

		if err != nil && err.Error() != el.err {
			t.Error("errors don't match: ", i)
		}

		if el.itemRep != nil {
			el.itemRep.AssertExpectations(t)
		}

		if el.catRep != nil {
			el.catRep.AssertExpectations(t)
		}

		if el.pres != nil {
			el.pres.AssertExpectations(t)
		}
	}
}
