package item_repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/pashagolub/pgxmock/v3"
	"gitlab.com/maometusu/qr_menu/internal/entity"
	"gitlab.com/maometusu/qr_menu/internal/entity/models"
	"testing"
)

func TestItemRepository_CreateItem(t *testing.T) {
	repo := NewItemRepository()

	type input struct {
		ctx    context.Context
		mockDB pgxmock.PgxPoolIface
		item   *models.Item
	}

	type output struct {
		err error
	}

	type testCase struct {
		input  input
		output output
	}

	mock, err := pgxmock.NewPool(pgxmock.QueryMatcherOption(pgxmock.QueryMatcherEqual))
	if err != nil {
		t.Error(err)
	}

	mock.ExpectExec(
		"update items set order_param=order_param+1 where order_param>=$1 and category_id = $2",
	).WithArgs(int64(1), int64(1)).WillReturnResult(pgxmock.NewResult("UPDATE", 1))

	mock.ExpectExec(
		"insert into items(name, weight, additional_info, is_visible, is_available, image, category_id, order_param, price) values ($1, $2, $3, $4, $5, $6, $7, $8, $9)",
	).WithArgs("name", "weight", "additional_info", true, true, "image", int64(1), int64(1), float64(1)).WillReturnResult(pgxmock.NewResult("INSERT", 1))

	cases := []testCase{
		{
			input{
				context.Background(),
				nil,
				nil,
			},
			output{
				errors.New("db not found in context"),
			},
		},
		{
			input{
				context.WithValue(context.Background(), "db", entity.PgxIface(mock)),
				mock,
				&models.Item{
					Name:           "name",
					Weight:         "weight",
					Price:          1,
					AdditionalInfo: "additional_info",
					IsVisible:      true,
					IsAvailable:    true,
					Image:          "image",
					CategoryID:     1,
					OrderParam:     1,
				},
			},
			output{
				nil,
			},
		},
	}

	for i, el := range cases {
		err := repo.CreateItem(el.input.ctx, el.input.item)
		if err != nil && el.output.err != nil && err.Error() != el.output.err.Error() {
			t.Error(fmt.Sprintf("error don't match %d, %v", i, err))
		} else if err != nil && el.output.err == nil {
			t.Error("did not expect error: " + err.Error())
		}

		if el.input.mockDB != nil {
			err = el.input.mockDB.ExpectationsWereMet()
			if err != nil {
				t.Error(err)
			}
		}
	}
}
