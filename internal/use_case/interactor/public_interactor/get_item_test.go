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

func TestPublicInteractor_GetItem(t *testing.T) {
	type testCase struct {
		itemID        int64
		selectedItems map[int64][]int64
		itemRep       *repository.ItemRepositoryMock
		pres          *presenter.PublicPresenterMock
		err           string
	}

	cases := make([]testCase, 0)

	itemRep := &repository.ItemRepositoryMock{}
	itemRep.On("GetOne", context.Background(), int64(1)).Return(&models.Item{}, errors.New("error"))

	cases = append(cases, testCase{
		itemID:        1,
		selectedItems: nil,
		itemRep:       itemRep,
		pres:          nil,
		err:           "error",
	})

	itemRep = &repository.ItemRepositoryMock{}
	itemRep.On("GetOne", context.Background(), int64(1)).Return(&models.Item{}, nil)
	itemRep.On("GetEstablishment", context.Background(), int64(1)).Return(0, errors.New("error1"))

	cases = append(cases, testCase{
		itemID:        1,
		selectedItems: nil,
		itemRep:       itemRep,
		pres:          nil,
		err:           "error1",
	})

	item := &models.Item{
		ID: 1,
	}
	itemRep = &repository.ItemRepositoryMock{}
	itemRep.On("GetOne", context.Background(), int64(1)).Return(item, nil)
	itemRep.On("GetEstablishment", context.Background(), int64(1)).Return(2, nil)
	pres := &presenter.PublicPresenterMock{}
	pres.On("PresentItem", nil, mock.MatchedBy(func(i *models.Item) bool {
		return i.ID == 1 && i.SelectedAmount == 3
	}), true).Return(errors.New("error2"))

	cases = append(cases, testCase{
		itemID: 1,
		selectedItems: map[int64][]int64{
			2: {1, 1, 1},
		},
		itemRep: itemRep,
		pres:    pres,
		err:     "error2",
	})

	item = &models.Item{
		ID: 1,
	}
	itemRep = &repository.ItemRepositoryMock{}
	itemRep.On("GetOne", context.Background(), int64(1)).Return(item, nil)
	itemRep.On("GetEstablishment", context.Background(), int64(1)).Return(2, nil)
	pres = &presenter.PublicPresenterMock{}
	pres.On("PresentItem", nil, mock.MatchedBy(func(i *models.Item) bool {
		return i.ID == 1 && i.SelectedAmount == 3
	}), true).Return(nil)

	cases = append(cases, testCase{
		itemID: 1,
		selectedItems: map[int64][]int64{
			2: {1, 1, 1},
		},
		itemRep: itemRep,
		pres:    pres,
		err:     "",
	})

	for i, el := range cases {
		interactor := NewPublicInteractor(el.pres, nil, nil, nil, el.itemRep)
		err := interactor.GetItem(context.Background(), nil, el.itemID, el.selectedItems)

		if err != nil && err.Error() != el.err {
			t.Error("errors don't match: ", i)
		}

		if el.itemRep != nil {
			el.itemRep.AssertExpectations(t)
		}
	}
}
