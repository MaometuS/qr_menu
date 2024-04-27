package public_interactor

import (
	"context"
	"github.com/stretchr/testify/mock"
	"io"
)

type TestImplementation struct {
	mock.Mock
}

func (t *TestImplementation) GetMainPage(ctx context.Context, w io.Writer, link string, selectedItems map[int64][]int64) error {
	args := t.Called(ctx, w, link, selectedItems)
	return args.Error(0)
}

func (t *TestImplementation) GetCategories(ctx context.Context, w io.Writer, menuID int64) error {
	args := t.Called(ctx, w, menuID)
	return args.Error(0)
}

func (t *TestImplementation) GetItems(ctx context.Context, w io.Writer, categoryID int64, selectedItems map[int64][]int64) error {
	args := t.Called(ctx, w, categoryID, selectedItems)
	return args.Error(0)
}

func (t *TestImplementation) GetSearch(ctx context.Context, w io.Writer, search string, establishmentID int64, menuID int64, selectedItems map[int64][]int64) error {
	args := t.Called(ctx, w, search, establishmentID, menuID, selectedItems)
	return args.Error(0)
}

func (t *TestImplementation) GetItem(ctx context.Context, w io.Writer, itemID int64, selectedItems map[int64][]int64) error {
	args := t.Called(ctx, w, itemID, selectedItems)
	return args.Error(0)
}

func (t *TestImplementation) AddItem(ctx context.Context, itemID int64, selectedItems map[int64][]int64) (map[int64][]int64, error) {
	args := t.Called(ctx, itemID, selectedItems)
	return args.Get(0).(map[int64][]int64), args.Error(1)
}

func (t *TestImplementation) RemoveItem(ctx context.Context, itemID int64, selectedItems map[int64][]int64) (map[int64][]int64, error) {
	args := t.Called(ctx, itemID, selectedItems)
	return args.Get(0).(map[int64][]int64), args.Error(1)
}

func (t *TestImplementation) GetOrder(ctx context.Context, w io.Writer, menuID int64, selectedItems map[int64][]int64) error {
	args := t.Called(ctx, w, menuID, selectedItems)
	return args.Error(0)
}

func NewTestImplementation() *TestImplementation {
	return &TestImplementation{}
}
