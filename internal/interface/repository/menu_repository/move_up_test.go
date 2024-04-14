package menu_repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/pashagolub/pgxmock/v3"
	"gitlab.com/maometusu/qr_menu/internal/entity"
	"testing"
)

func TestMenuRepository_MoveUp(t *testing.T) {
	repo := NewMenuRepository()

	type input struct {
		ctx    context.Context
		mockDB pgxmock.PgxPoolIface
		id     int64
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

	mock.ExpectQuery(
		"select order_param, establishment_id from menus where id = $1",
	).WithArgs(int64(1)).WillReturnRows(pgxmock.NewRows(
		[]string{"order_param", "menu_id"},
	).AddRow(int64(2), int64(1)))

	mock.ExpectExec(
		"update menus set order_param=order_param+1 where order_param = $1 and establishment_id = $2",
	).WithArgs(int64(1), int64(1)).WillReturnResult(pgxmock.NewResult("UPDATE", 1))

	mock.ExpectExec(
		"update menus set order_param=order_param-1 where id = $1",
	).WithArgs(int64(1)).WillReturnResult(pgxmock.NewResult("UPDATE", 1))

	cases := []testCase{
		{
			input{
				context.Background(),
				nil,
				0,
			},
			output{
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
				nil,
			},
		},
	}

	for i, el := range cases {
		err := repo.MoveUp(el.input.ctx, el.input.id)
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
