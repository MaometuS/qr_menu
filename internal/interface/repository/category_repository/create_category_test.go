package category_repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/pashagolub/pgxmock/v3"
	"gitlab.com/maometusu/qr_menu/internal/entity"
	"gitlab.com/maometusu/qr_menu/internal/entity/models"
	"testing"
)

func TestCategoryRepository_CreateCategory(t *testing.T) {
	repo := NewCategoryRepository()

	type input struct {
		ctx      context.Context
		mockDB   pgxmock.PgxPoolIface
		category *models.Category
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
		t.Error(err.Error())
	}
	mock.ExpectExec(
		"update categories set order_param=order_param+1 where order_param>=$1 and menu_id = $2",
	).WithArgs(int64(1), int64(2)).WillReturnResult(pgxmock.NewResult("UPDATE", 2))
	mock.ExpectExec(
		"insert into categories(name, background, menu_id, is_visible, order_param) values ($1, $2, $3, $4, $5)",
	).WithArgs(
		"name",
		"background",
		int64(2),
		true,
		int64(1),
	).WillReturnResult(pgxmock.NewResult("INSERT", 1))

	cases := []testCase{
		{
			input{
				context.Background(),
				nil,
				&models.Category{
					ID:         0,
					Name:       "",
					Background: "",
					MenuID:     0,
					IsVisible:  false,
					OrderParam: 0,
				},
			},
			output{errors.New("db not found in context")},
		},
		{
			input{
				context.WithValue(context.Background(), "db", entity.PgxIface(mock)),
				mock,
				&models.Category{
					Name:       "name",
					Background: "background",
					MenuID:     2,
					IsVisible:  true,
					OrderParam: 1,
				},
			},
			output{
				nil,
			},
		},
	}

	for i, el := range cases {
		err := repo.CreateCategory(el.input.ctx, el.input.category)
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
