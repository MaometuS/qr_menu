package item_interactor

import (
	"context"
	"errors"
	"gitlab.com/maometusu/qr_menu/internal/entity/models"
	"gitlab.com/maometusu/qr_menu/internal/use_case/repository"
	"testing"
)

func TestItemInteractor_CreateItem(t *testing.T) {
	type testCase struct {
		profileID int64
		item      *models.Item
		catRep    *repository.CategoryRepositoryMock
		rep       *repository.ItemRepositoryMock
		err       string
	}

	cases := make([]testCase, 0)

	catRep := &repository.CategoryRepositoryMock{}
	catRep.On("CheckBelongs", context.Background(), int64(1), int64(2)).Return(false, errors.New("error"))

	cases = append(cases, testCase{
		profileID: 2,
		item:      &models.Item{CategoryID: 1},
		catRep:    catRep,
		rep:       nil,
		err:       "error",
	})

	catRep = &repository.CategoryRepositoryMock{}
	catRep.On("CheckBelongs", context.Background(), int64(1), int64(2)).Return(false, nil)

	cases = append(cases, testCase{
		profileID: 2,
		item:      &models.Item{CategoryID: 1},
		catRep:    catRep,
		rep:       nil,
		err:       "category does not belong to current profile",
	})

	item := &models.Item{
		CategoryID: 1,
		Image:      "image",
	}
	catRep = &repository.CategoryRepositoryMock{}
	catRep.On("CheckBelongs", context.Background(), int64(1), int64(2)).Return(true, nil)
	rep := &repository.ItemRepositoryMock{}
	rep.On("UpdateItem", context.Background(), item).Return(errors.New("error2"))

	cases = append(cases, testCase{
		profileID: 2,
		item:      item,
		catRep:    catRep,
		rep:       rep,
		err:       "error2",
	})

	item = &models.Item{
		CategoryID: 1,
		Image:      "image",
	}
	catRep = &repository.CategoryRepositoryMock{}
	catRep.On("CheckBelongs", context.Background(), int64(1), int64(2)).Return(true, nil)
	rep = &repository.ItemRepositoryMock{}
	rep.On("UpdateItem", context.Background(), item).Return(nil)

	cases = append(cases, testCase{
		profileID: 2,
		item:      item,
		catRep:    catRep,
		rep:       rep,
		err:       "",
	})

	for i, el := range cases {
		interactor := NewItemInteractor(nil, el.rep, nil, el.catRep)
		err := interactor.EditItem(context.Background(), el.profileID, nil, el.item)
		if err != nil && err.Error() != el.err {
			t.Error("errors don't match: ", i, err)
		}

		if el.rep != nil {
			el.rep.AssertExpectations(t)
		}

		if el.catRep != nil {
			el.catRep.AssertExpectations(t)
		}
	}
}
