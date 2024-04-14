package category_repository

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

func TestCategoryRepository_GetOne(t *testing.T) {
	repo := NewCategoryRepository()

	type input struct {
		ctx    context.Context
		mockDB pgxmock.PgxPoolIface
		id     int64
	}

	type output struct {
		category *models.Category
		err      error
	}

	type testCase struct {
		input  input
		output output
	}

	mock, err := pgxmock.NewPool(pgxmock.QueryMatcherOption(pgxmock.QueryMatcherEqual))
	if err != nil {
		t.Error(err)
	}

	mock.ExpectQuery(
		"select * from categories where id = $1",
	).WithArgs(int64(1)).WillReturnRows(
		pgxmock.NewRows(
			[]string{"id", "name", "background", "menu_id", "is_visible", "order_param"},
		).AddRow(int64(1), "name", "background", int64(1), true, int64(1)),
	)

	cases := []testCase{
		{
			input{
				context.Background(),
				nil,
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
				1,
			},
			output{
				&models.Category{
					ID:         1,
					Name:       "name",
					Background: "background",
					MenuID:     1,
					IsVisible:  true,
					OrderParam: 1,
				},
				nil,
			},
		},
	}

	for i, el := range cases {
		res, err := repo.GetOne(el.input.ctx, el.input.id)
		if err != nil && el.output.err != nil && err.Error() != el.output.err.Error() {
			t.Error(fmt.Sprintf("error don't match %d, %v", i, err))
		} else if err != nil && el.output.err == nil {
			t.Error("did not expect error: " + err.Error())
		}

		if !reflect.DeepEqual(res, el.output.category) {
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
