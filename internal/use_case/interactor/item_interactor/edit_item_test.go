package item_interactor

import (
	"context"
	"errors"
	"gitlab.com/maometusu/qr_menu/internal/entity/models"
	"gitlab.com/maometusu/qr_menu/internal/use_case/repository"
	"testing"
)

func TestItemInteractor_EditItem(t *testing.T) {
	type testCase struct {
		profileID int64
		item      *models.Item
		catRep    *repository.CategoryRepositoryMock
		fileRep   *repository.FileRepositoryMock
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
		fileRep:   nil,
		rep:       nil,
		err:       "error",
	})

	catRep = &repository.CategoryRepositoryMock{}
	catRep.On("CheckBelongs", context.Background(), int64(1), int64(2)).Return(false, nil)

	cases = append(cases, testCase{
		profileID: 2,
		item:      &models.Item{CategoryID: 1},
		catRep:    catRep,
		fileRep:   nil,
		rep:       nil,
		err:       "category does not belong to current profile",
	})

	catRep = &repository.CategoryRepositoryMock{}
	catRep.On("CheckBelongs", context.Background(), int64(1), int64(2)).Return(true, nil)
	fileRep := &repository.FileRepositoryMock{}
	fileRep.On("SaveFile", context.Background(), nil).Return("", errors.New("error1"))

	cases = append(cases, testCase{
		profileID: 2,
		item:      &models.Item{CategoryID: 1},
		catRep:    catRep,
		fileRep:   fileRep,
		rep:       nil,
		err:       "error1",
	})

	item := &models.Item{
		CategoryID: 1,
		Image:      "image",
	}
	catRep = &repository.CategoryRepositoryMock{}
	catRep.On("CheckBelongs", context.Background(), int64(1), int64(2)).Return(true, nil)
	fileRep = &repository.FileRepositoryMock{}
	fileRep.On("SaveFile", context.Background(), nil).Return("image", nil)
	rep := &repository.ItemRepositoryMock{}
	rep.On("CreateItem", context.Background(), item).Return(errors.New("error2"))

	cases = append(cases, testCase{
		profileID: 2,
		item:      item,
		catRep:    catRep,
		fileRep:   fileRep,
		rep:       rep,
		err:       "error2",
	})

	item = &models.Item{
		CategoryID: 1,
		Image:      "image",
	}
	catRep = &repository.CategoryRepositoryMock{}
	catRep.On("CheckBelongs", context.Background(), int64(1), int64(2)).Return(true, nil)
	fileRep = &repository.FileRepositoryMock{}
	fileRep.On("SaveFile", context.Background(), nil).Return("image", nil)
	rep = &repository.ItemRepositoryMock{}
	rep.On("CreateItem", context.Background(), item).Return(nil)

	cases = append(cases, testCase{
		profileID: 2,
		item:      item,
		catRep:    catRep,
		fileRep:   fileRep,
		rep:       rep,
		err:       "",
	})

	for i, el := range cases {
		interactor := NewItemInteractor(nil, el.rep, el.fileRep, el.catRep)
		err := interactor.CreateItem(context.Background(), el.profileID, nil, el.item)
		if err != nil && err.Error() != el.err {
			t.Error("errors don't match: ", i, err)
		}

		if el.rep != nil {
			el.rep.AssertExpectations(t)
		}

		if el.fileRep != nil {
			el.fileRep.AssertExpectations(t)
		}

		if el.catRep != nil {
			el.catRep.AssertExpectations(t)
		}
	}
}
