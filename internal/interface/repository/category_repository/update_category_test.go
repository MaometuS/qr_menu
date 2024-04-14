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

func TestCategoryRepository_UpdateCategory(t *testing.T) {
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
		t.Error(err)
	}
	mock.ExpectExec(
		"update categories set name = $1, menu_id = $2, is_visible = $3, order_param = $4 where id = $5",
	).WithArgs(
		"name",
		int64(2),
		true,
		int64(1),
		int64(1),
	).WillReturnResult(pgxmock.NewResult("UPDATE", 1))

	mock.ExpectExec(
		"update categories set background = $1 where id = $2",
	).WithArgs("background", int64(1)).WillReturnResult(pgxmock.NewResult("UPDATE", int64(1)))

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
					ID:         1,
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
		err := repo.UpdateCategory(el.input.ctx, el.input.category)
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
