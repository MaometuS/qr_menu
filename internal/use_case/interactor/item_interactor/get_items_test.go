package item_interactor

import (
	"context"
	"errors"
	"gitlab.com/maometusu/qr_menu/internal/entity/models"
	"gitlab.com/maometusu/qr_menu/internal/use_case/presenter"
	"gitlab.com/maometusu/qr_menu/internal/use_case/repository"
	"testing"
)

func TestItemInteractor_GetItems(t *testing.T) {
	type testCase struct {
		categoryID int64
		catRep     *repository.CategoryRepositoryMock
		rep        *repository.ItemRepositoryMock
		pres       *presenter.ItemPresenterMock
		err        string
	}

	cases := make([]testCase, 0)

	catRep := &repository.CategoryRepositoryMock{}
	catRep.On("GetOne", context.Background(), int64(1)).Return(&models.Category{}, errors.New("error"))

	cases = append(cases, testCase{
		categoryID: 1,
		catRep:     catRep,
		rep:        nil,
		pres:       nil,
		err:        "error",
	})

	catRep = &repository.CategoryRepositoryMock{}
	catRep.On("GetOne", context.Background(), int64(1)).Return(&models.Category{}, nil)
	rep := &repository.ItemRepositoryMock{}
	rep.On("GetAll", context.Background(), int64(1)).Return(make([]models.Item, 0), errors.New("error1"))

	cases = append(cases, testCase{
		categoryID: 1,
		catRep:     catRep,
		rep:        rep,
		pres:       nil,
		err:        "error1",
	})

	category := &models.Category{
		MenuID: 5,
	}
	items := make([]models.Item, 0)
	catRep = &repository.CategoryRepositoryMock{}
	catRep.On("GetOne", context.Background(), int64(1)).Return(category, nil)
	rep = &repository.ItemRepositoryMock{}
	rep.On("GetAll", context.Background(), int64(1)).Return(items, nil)
	pres := &presenter.ItemPresenterMock{}
	pres.On("PresentItems", nil, items, int64(1), int64(5)).Return(errors.New("error2"))

	cases = append(cases, testCase{
		categoryID: 1,
		catRep:     catRep,
		rep:        rep,
		pres:       pres,
		err:        "error2",
	})

	category = &models.Category{
		MenuID: 5,
	}
	items = make([]models.Item, 0)
	catRep = &repository.CategoryRepositoryMock{}
	catRep.On("GetOne", context.Background(), int64(1)).Return(category, nil)
	rep = &repository.ItemRepositoryMock{}
	rep.On("GetAll", context.Background(), int64(1)).Return(items, nil)
	pres = &presenter.ItemPresenterMock{}
	pres.On("PresentItems", nil, items, int64(1), int64(5)).Return(nil)

	cases = append(cases, testCase{
		categoryID: 1,
		catRep:     catRep,
		rep:        rep,
		pres:       pres,
		err:        "",
	})

	for i, el := range cases {
		interactor := NewItemInteractor(el.pres, el.rep, nil, el.catRep)
		err := interactor.GetItems(context.Background(), nil, 0, el.categoryID)
		if err != nil && err.Error() != el.err {
			t.Error("errors don't match: ", i, err)
		}

		if el.rep != nil {
			el.rep.AssertExpectations(t)
		}

		if el.catRep != nil {
			el.catRep.AssertExpectations(t)
		}

		if el.pres != nil {
			el.pres.AssertExpectations(t)
		}
	}
}
