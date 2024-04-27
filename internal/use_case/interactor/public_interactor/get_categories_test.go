package public_interactor

import (
	"context"
	"errors"
	"gitlab.com/maometusu/qr_menu/internal/entity/models"
	"gitlab.com/maometusu/qr_menu/internal/use_case/presenter"
	"gitlab.com/maometusu/qr_menu/internal/use_case/repository"
	"testing"
)

func TestPublicInteractor_GetCategories(t *testing.T) {
	type testCase struct {
		menuID  int64
		menuRep *repository.MenuRepositoryMock
		catRep  *repository.CategoryRepositoryMock
		pres    *presenter.PublicPresenterMock
		err     string
	}

	cases := make([]testCase, 0)

	menuRep := &repository.MenuRepositoryMock{}
	menuRep.On("GetOne", context.Background(), int64(1)).Return(&models.Menu{}, errors.New("error"))

	cases = append(cases, testCase{
		menuID:  1,
		menuRep: menuRep,
		catRep:  nil,
		pres:    nil,
		err:     "error",
	})

	menuRep = &repository.MenuRepositoryMock{}
	menuRep.On("GetOne", context.Background(), int64(1)).Return(&models.Menu{}, nil)
	catRep := &repository.CategoryRepositoryMock{}
	catRep.On("GetAll", context.Background(), int64(1)).Return(make([]models.Category, 0), errors.New("error1"))

	cases = append(cases, testCase{
		menuID:  1,
		menuRep: menuRep,
		catRep:  catRep,
		pres:    nil,
		err:     "error1",
	})

	menu := &models.Menu{
		Name:            "name",
		EstablishmentID: 5,
	}
	categories := make([]models.Category, 0)
	menuRep = &repository.MenuRepositoryMock{}
	menuRep.On("GetOne", context.Background(), int64(1)).Return(menu, nil)
	catRep = &repository.CategoryRepositoryMock{}
	catRep.On("GetAll", context.Background(), int64(1)).Return(categories, nil)
	pres := &presenter.PublicPresenterMock{}
	pres.On("PresentCategories", nil, "name", categories, int64(5), int64(1)).Return(errors.New("error2"))

	cases = append(cases, testCase{
		menuID:  1,
		menuRep: menuRep,
		catRep:  catRep,
		pres:    pres,
		err:     "error2",
	})

	menu = &models.Menu{
		Name:            "name",
		EstablishmentID: 5,
	}
	categories = make([]models.Category, 0)
	menuRep = &repository.MenuRepositoryMock{}
	menuRep.On("GetOne", context.Background(), int64(1)).Return(menu, nil)
	catRep = &repository.CategoryRepositoryMock{}
	catRep.On("GetAll", context.Background(), int64(1)).Return(categories, nil)
	pres = &presenter.PublicPresenterMock{}
	pres.On("PresentCategories", nil, "name", categories, int64(5), int64(1)).Return(nil)

	cases = append(cases, testCase{
		menuID:  1,
		menuRep: menuRep,
		catRep:  catRep,
		pres:    pres,
		err:     "",
	})

	for i, el := range cases {
		interactor := NewPublicInteractor(el.pres, nil, el.menuRep, el.catRep, nil)
		err := interactor.GetCategories(context.Background(), nil, el.menuID)

		if err != nil && err.Error() != el.err {
			t.Error("errors don't match: ", i)
		}

		if el.menuRep != nil {
			el.menuRep.AssertExpectations(t)
		}

		if el.catRep != nil {
			el.catRep.AssertExpectations(t)
		}
	}
}
