package category_repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/pashagolub/pgxmock/v3"
	"gitlab.com/maometusu/qr_menu/internal/entity"
	"testing"
)

func TestCategoryRepository_CheckBelongs(t *testing.T) {
	repo := NewCategoryRepository()

	type input struct {
		ctx        context.Context
		mockDB     pgxmock.PgxPoolIface
		categoryID int64
		profileID  int64
	}

	type output struct {
		result bool
		err    error
	}

	type testCase struct {
		input  input
		output output
	}

	mock1, _ := pgxmock.NewPool(pgxmock.QueryMatcherOption(pgxmock.QueryMatcherEqual))
	mock1.ExpectQuery(`select count(id) from categories where id = $1 and menu_id in (
				select id from menus where establishment_id in (
					select id from establishments where profile_id = $2
				)
			 )`).WithArgs(int64(1), int64(2)).WillReturnRows(pgxmock.NewRows([]string{"count"}).AddRow(int64(1)))
	defer mock1.Close()

	mock2, _ := pgxmock.NewPool(pgxmock.QueryMatcherOption(pgxmock.QueryMatcherEqual))
	mock2.ExpectQuery(`select count(id) from categories where id = $1 and menu_id in (
				select id from menus where establishment_id in (
					select id from establishments where profile_id = $2
				)
			 )`).WithArgs(int64(1), int64(5)).WillReturnRows(pgxmock.NewRows([]string{"count"}).AddRow(int64(0)))
	defer mock2.Close()

	cases := []testCase{
		{
			input{
				context.Background(),
				nil,
				0,
				0,
			},
			output{
				false,
				errors.New("db not found in context"),
			},
		},
		{
			input{
				context.WithValue(context.Background(), "db", entity.PgxIface(mock1)),
				mock1,
				1,
				2,
			},
			output{
				true,
				nil,
			},
		},
		{
			input{
				context.WithValue(context.Background(), "db", entity.PgxIface(mock2)),
				mock2,
				1,
				5,
			},
			output{
				false,
				nil,
			},
		},
	}

	for i, el := range cases {
		res, err := repo.CheckBelongs(el.input.ctx, el.input.categoryID, el.input.profileID)
		if err != nil && el.output.err != nil && err.Error() != el.output.err.Error() {
			t.Error(fmt.Sprintf("error don't match %d, %v", i, err))
		} else if err != nil && el.output.err == nil {
			t.Error("did not expect error: " + err.Error())
		}

		if res != el.output.result {
			t.Error(fmt.Sprintf("result don't match %d", i))
		}

		if el.input.mockDB != nil {
			err = el.input.mockDB.ExpectationsWereMet()
			if err != nil {
				t.Error(err)
			}
		}
	}
}
