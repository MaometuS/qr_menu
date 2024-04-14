package item_repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/pashagolub/pgxmock/v3"
	"gitlab.com/maometusu/qr_menu/internal/entity"
	"gitlab.com/maometusu/qr_menu/internal/entity/models"
	"reflect"
	"testing"
)

func TestItemRepository_Search(t *testing.T) {
	repo := NewItemRepository()

	type input struct {
		ctx             context.Context
		mockDB          pgxmock.PgxPoolIface
		search          string
		establishmentID int64
	}

	type output struct {
		items []models.Item
		err   error
	}

	type testCase struct {
		input  input
		output output
	}

	mock, err := pgxmock.NewPool(pgxmock.QueryMatcherOption(pgxmock.QueryMatcherEqual))
	if err != nil {
		t.Error(err)
	}

	mock.ExpectQuery(`select * from items where 
		(name like concat('%', $1::varchar, '%') or additional_info like concat('%', $1::varchar, '%')) and
		category_id in (select id from categories where menu_id in (select id from menus where establishment_id = $2))`,
	).WithArgs(
		"search", int64(1),
	).WillReturnRows(
		pgxmock.NewRows(
			[]string{"id", "name", "weight", "price", "additional_info", "is_visible", "is_available", "image", "category_id", "order_param"},
		).AddRow(
			int64(1), "name", "weight", float64(1), "additional_info", true, true, "image", int64(1), int64(1),
		),
	)

	cases := []testCase{
		{
			input{
				context.Background(),
				nil,
				"",
				0,
			},
			output{
				nil,
				errors.New("db not found in context"),
			},
		},
		{
			input{
				context.WithValue(context.Background(), "db", entity.PgxIface(mock)),
				mock,
				"search",
				1,
			},
			output{
				[]models.Item{
					{int64(1), "name", "weight", float64(1), "additional_info", true, true, "image", int64(1), int64(1), 0},
				},
				nil,
			},
		},
	}

	for i, el := range cases {
		res, err := repo.Search(el.input.ctx, el.input.search, el.input.establishmentID)
		if err != nil && el.output.err != nil && err.Error() != el.output.err.Error() {
			t.Error(fmt.Sprintf("error don't match %d, %v", i, err))
		} else if err != nil && el.output.err == nil {
			t.Error("did not expect error: " + err.Error())
		}

		if !reflect.DeepEqual(res, el.output.items) {
			t.Error("results don't match")
		}

		if el.input.mockDB != nil {
			err = el.input.mockDB.ExpectationsWereMet()
			if err != nil {
				t.Error(err)
			}
		}
	}
}
