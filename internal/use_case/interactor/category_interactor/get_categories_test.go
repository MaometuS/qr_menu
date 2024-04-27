package category_interactor

import (
	"context"
	"errors"
	"gitlab.com/maometusu/qr_menu/internal/entity/models"
	"gitlab.com/maometusu/qr_menu/internal/use_case/presenter"
	"gitlab.com/maometusu/qr_menu/internal/use_case/repository"
	"testing"
)

func TestCategoryInteractor_GetCategories(t *testing.T) {
	type testCase struct {
		menuID int64
		rep    *repository.CategoryRepositoryMock
		pres   *presenter.CategoryPresenterMock
		err    string
	}

	cases := make([]testCase, 0)

	rep := &repository.CategoryRepositoryMock{}
	rep.On("GetAll", context.Background(), int64(1)).Return(make([]models.Category, 0), errors.New("error"))

	cases = append(cases, testCase{
		menuID: 1,
		rep:    rep,
		pres:   nil,
		err:    "error",
	})

	categories := make([]models.Category, 0)
	rep = &repository.CategoryRepositoryMock{}
	rep.On("GetAll", context.Background(), int64(1)).Return(categories, nil)
	pres := &presenter.CategoryPresenterMock{}
	pres.On("PresentCategories", nil, categories, int64(1)).Return(errors.New("error1"))

	cases = append(cases, testCase{
		menuID: 1,
		rep:    rep,
		pres:   pres,
		err:    "error1",
	})

	categories = make([]models.Category, 0)
	rep = &repository.CategoryRepositoryMock{}
	rep.On("GetAll", context.Background(), int64(1)).Return(categories, nil)
	pres = &presenter.CategoryPresenterMock{}
	pres.On("PresentCategories", nil, categories, int64(1)).Return(nil)

	cases = append(cases, testCase{
		menuID: 1,
		rep:    rep,
		pres:   pres,
		err:    "",
	})

	for i, el := range cases {
		interactor := NewCategoryInteractor(el.pres, el.rep, nil, nil)
		err := interactor.GetCategories(context.Background(), nil, 0, el.menuID)
		if err != nil && err.Error() != el.err {
			t.Error("errors don't match: ", i, err)
		}

		if el.rep != nil {
			el.rep.AssertExpectations(t)
		}
	}
}
