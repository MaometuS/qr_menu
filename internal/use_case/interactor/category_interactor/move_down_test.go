package category_interactor

import (
	"context"
	"errors"
	"gitlab.com/maometusu/qr_menu/internal/use_case/repository"
	"testing"
)

func TestCategoryInteractor_MoveDown(t *testing.T) {
	type testCase struct {
		profileID  int64
		categoryID int64
		rep        *repository.CategoryRepositoryMock
		err        string
	}

	cases := make([]testCase, 0)

	rep := &repository.CategoryRepositoryMock{}
	rep.On("CheckBelongs", context.Background(), int64(2), int64(1)).Return(false, errors.New("error"))

	cases = append(cases, testCase{
		profileID:  1,
		categoryID: 2,
		rep:        rep,
		err:        "error",
	})

	rep = &repository.CategoryRepositoryMock{}
	rep.On("CheckBelongs", context.Background(), int64(2), int64(1)).Return(false, nil)

	cases = append(cases, testCase{
		profileID:  1,
		categoryID: 2,
		rep:        rep,
		err:        "category does not belong to current profile",
	})

	rep = &repository.CategoryRepositoryMock{}
	rep.On("CheckBelongs", context.Background(), int64(2), int64(1)).Return(true, nil)
	rep.On("MoveDown", context.Background(), int64(2)).Return(errors.New("error1"))

	cases = append(cases, testCase{
		profileID:  1,
		categoryID: 2,
		rep:        rep,
		err:        "error1",
	})

	rep = &repository.CategoryRepositoryMock{}
	rep.On("CheckBelongs", context.Background(), int64(2), int64(1)).Return(true, nil)
	rep.On("MoveDown", context.Background(), int64(2)).Return(nil)

	cases = append(cases, testCase{
		profileID:  1,
		categoryID: 2,
		rep:        rep,
		err:        "",
	})

	for i, el := range cases {
		interactor := NewCategoryInteractor(nil, el.rep, nil, nil)
		err := interactor.MoveDown(context.Background(), el.profileID, el.categoryID)
		if err != nil && err.Error() != el.err {
			t.Error("errors don't match: ", i, err)
		}

		if el.rep != nil {
			el.rep.AssertExpectations(t)
		}
	}
}
