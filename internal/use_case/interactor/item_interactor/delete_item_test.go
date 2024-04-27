package item_interactor

import (
	"context"
	"errors"
	"gitlab.com/maometusu/qr_menu/internal/use_case/repository"
	"testing"
)

func TestItemInteractor_DeleteItem(t *testing.T) {
	type testCase struct {
		profileID int64
		itemID    int64
		rep       *repository.ItemRepositoryMock
		err       string
	}

	cases := make([]testCase, 0)

	rep := &repository.ItemRepositoryMock{}
	rep.On("CheckBelongs", context.Background(), int64(1), int64(2)).Return(false, errors.New("error"))

	cases = append(cases, testCase{
		profileID: 2,
		itemID:    1,
		rep:       rep,
		err:       "error",
	})

	rep = &repository.ItemRepositoryMock{}
	rep.On("CheckBelongs", context.Background(), int64(1), int64(2)).Return(false, nil)

	cases = append(cases, testCase{
		profileID: 2,
		itemID:    1,
		rep:       rep,
		err:       "item does not belong to current profile",
	})

	rep = &repository.ItemRepositoryMock{}
	rep.On("CheckBelongs", context.Background(), int64(1), int64(2)).Return(true, nil)
	rep.On("DeleteItem", context.Background(), int64(1)).Return(errors.New("error1"))

	cases = append(cases, testCase{
		profileID: 2,
		itemID:    1,
		rep:       rep,
		err:       "error1",
	})

	rep = &repository.ItemRepositoryMock{}
	rep.On("CheckBelongs", context.Background(), int64(1), int64(2)).Return(true, nil)
	rep.On("DeleteItem", context.Background(), int64(1)).Return(nil)

	cases = append(cases, testCase{
		profileID: 2,
		itemID:    1,
		rep:       rep,
		err:       "",
	})

	for i, el := range cases {
		interactor := NewItemInteractor(nil, el.rep, nil, nil)
		err := interactor.DeleteItem(context.Background(), el.profileID, el.itemID)
		if err != nil && err.Error() != el.err {
			t.Error("errors don't match: ", i, err)
		}

		if el.rep != nil {
			el.rep.AssertExpectations(t)
		}
	}
}
