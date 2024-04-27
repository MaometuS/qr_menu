package menu_interactor

import (
	"context"
	"errors"
	"gitlab.com/maometusu/qr_menu/internal/use_case/repository"
	"testing"
)

func TestInteractor_MoveUp(t *testing.T) {
	type testCase struct {
		profileID  int64
		categoryID int64
		rep        *repository.MenuRepositoryMock
		err        string
	}

	cases := make([]testCase, 0)

	rep := &repository.MenuRepositoryMock{}
	rep.On("CheckBelongs", context.Background(), int64(2), int64(1)).Return(false, errors.New("error"))

	cases = append(cases, testCase{
		profileID:  1,
		categoryID: 2,
		rep:        rep,
		err:        "error",
	})

	rep = &repository.MenuRepositoryMock{}
	rep.On("CheckBelongs", context.Background(), int64(2), int64(1)).Return(false, nil)

	cases = append(cases, testCase{
		profileID:  1,
		categoryID: 2,
		rep:        rep,
		err:        "not allowed",
	})

	rep = &repository.MenuRepositoryMock{}
	rep.On("CheckBelongs", context.Background(), int64(2), int64(1)).Return(true, nil)
	rep.On("MoveUp", context.Background(), int64(2)).Return(errors.New("error1"))

	cases = append(cases, testCase{
		profileID:  1,
		categoryID: 2,
		rep:        rep,
		err:        "error1",
	})

	rep = &repository.MenuRepositoryMock{}
	rep.On("CheckBelongs", context.Background(), int64(2), int64(1)).Return(true, nil)
	rep.On("MoveUp", context.Background(), int64(2)).Return(nil)

	cases = append(cases, testCase{
		profileID:  1,
		categoryID: 2,
		rep:        rep,
		err:        "",
	})

	for i, el := range cases {
		interactor := NewMenuInteractor(nil, el.rep, nil)
		err := interactor.MoveUp(context.Background(), el.profileID, el.categoryID)
		if err != nil && err.Error() != el.err {
			t.Error("errors don't match: ", i, err)
		}

		if el.rep != nil {
			el.rep.AssertExpectations(t)
		}
	}
}
