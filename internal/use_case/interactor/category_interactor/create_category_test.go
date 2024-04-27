package category_interactor

import (
	"context"
	"errors"
	"gitlab.com/maometusu/qr_menu/internal/entity/models"
	"gitlab.com/maometusu/qr_menu/internal/use_case/repository"
	"testing"
)

func TestCategoryInteractor_CreateCategory(t *testing.T) {
	type testCase struct {
		profileID int64
		category  *models.Category
		menuRep   *repository.MenuRepositoryMock
		fileRep   *repository.FileRepositoryMock
		rep       *repository.CategoryRepositoryMock
		err       string
	}
	cases := make([]testCase, 0)

	menuRep := &repository.MenuRepositoryMock{}
	menuRep.On("CheckBelongs", context.Background(), int64(1), int64(2)).Return(false, errors.New("error"))

	cases = append(cases, testCase{
		profileID: 2,
		category: &models.Category{
			MenuID: 1,
		},
		menuRep: menuRep,
		err:     "error",
	})

	menuRep = &repository.MenuRepositoryMock{}
	menuRep.On("CheckBelongs", context.Background(), int64(1), int64(2)).Return(false, nil)

	cases = append(cases, testCase{
		profileID: 2,
		category: &models.Category{
			MenuID: 1,
		},
		menuRep: menuRep,
		err:     "menu does not belong to current profile",
	})

	menuRep = &repository.MenuRepositoryMock{}
	menuRep.On("CheckBelongs", context.Background(), int64(1), int64(2)).Return(true, nil)
	fileRep := &repository.FileRepositoryMock{}
	fileRep.On("SaveFile", context.Background(), nil).Return("", errors.New("error1"))

	cases = append(cases, testCase{
		profileID: 2,
		category: &models.Category{
			MenuID: 1,
		},
		menuRep: menuRep,
		fileRep: fileRep,
		err:     "error1",
	})

	category := &models.Category{
		Background: "background",
		MenuID:     1,
	}
	menuRep = &repository.MenuRepositoryMock{}
	menuRep.On("CheckBelongs", context.Background(), int64(1), int64(2)).Return(true, nil)
	fileRep = &repository.FileRepositoryMock{}
	fileRep.On("SaveFile", context.Background(), nil).Return("background", nil)
	rep := &repository.CategoryRepositoryMock{}
	rep.On("CreateCategory", context.Background(), category).Return(errors.New("error2"))

	cases = append(cases, testCase{
		profileID: 2,
		category:  category,
		menuRep:   menuRep,
		fileRep:   fileRep,
		rep:       rep,
		err:       "error2",
	})

	category = &models.Category{
		Background: "background",
		MenuID:     1,
	}
	menuRep = &repository.MenuRepositoryMock{}
	menuRep.On("CheckBelongs", context.Background(), int64(1), int64(2)).Return(true, nil)
	fileRep = &repository.FileRepositoryMock{}
	fileRep.On("SaveFile", context.Background(), nil).Return("background", nil)
	rep = &repository.CategoryRepositoryMock{}
	rep.On("CreateCategory", context.Background(), category).Return(nil)

	cases = append(cases, testCase{
		profileID: 2,
		category:  category,
		menuRep:   menuRep,
		fileRep:   fileRep,
		rep:       rep,
	})

	for i, el := range cases {
		interactor := NewCategoryInteractor(nil, el.rep, el.menuRep, el.fileRep)
		err := interactor.CreateCategory(context.Background(), el.profileID, nil, el.category)
		if err != nil && err.Error() != el.err {
			t.Error("errors don't match: ", i, err)
		}

		if el.rep != nil {
			el.rep.AssertExpectations(t)
		}

		if el.menuRep != nil {
			el.menuRep.AssertExpectations(t)
		}

		if el.fileRep != nil {
			el.fileRep.AssertExpectations(t)
		}
	}
}
