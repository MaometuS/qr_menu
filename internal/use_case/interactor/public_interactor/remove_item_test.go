package public_interactor

import (
	"context"
	"errors"
	"gitlab.com/maometusu/qr_menu/internal/use_case/repository"
	"reflect"
	"testing"
)

func TestPublicInteractor_RemoveItem(t *testing.T) {
	type testCase struct {
		itemID        int64
		selectedItems map[int64][]int64
		itemRep       *repository.ItemRepositoryMock
		res           map[int64][]int64
		err           string
	}

	cases := make([]testCase, 0)

	itemRep := &repository.ItemRepositoryMock{}
	itemRep.On("GetEstablishment", context.Background(), int64(1)).Return(0, errors.New("error"))

	cases = append(cases, testCase{
		itemID:        1,
		selectedItems: nil,
		itemRep:       itemRep,
		res:           nil,
		err:           "error",
	})

	itemRep = &repository.ItemRepositoryMock{}
	itemRep.On("GetEstablishment", context.Background(), int64(1)).Return(2, nil)

	cases = append(cases, testCase{
		itemID: 1,
		selectedItems: map[int64][]int64{
			2: {1, 3, 5},
		},
		itemRep: itemRep,
		res: map[int64][]int64{
			2: {3, 5},
		},
		err: "",
	})

	for i, el := range cases {
		interactor := NewPublicInteractor(nil, nil, nil, nil, el.itemRep)
		res, err := interactor.RemoveItem(context.Background(), el.itemID, el.selectedItems)

		if err != nil && err.Error() != el.err {
			t.Error("errors don't match: ", i)
		}

		if !reflect.DeepEqual(res, el.res) {
			t.Error("results don't match: ", i)
		}
	}
}
