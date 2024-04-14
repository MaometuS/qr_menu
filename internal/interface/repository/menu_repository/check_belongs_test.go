package menu_repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/pashagolub/pgxmock/v3"
	"gitlab.com/maometusu/qr_menu/internal/entity"
	"testing"
)

func TestMenuRepository_CheckBelongs(t *testing.T) {
	repo := NewMenuRepository()

	type input struct {
		ctx       context.Context
		mockDB    pgxmock.PgxPoolIface
		menuID    int64
		profileID int64
	}

	type output struct {
		belongs bool
		err     error
	}

	type testCase struct {
		input  input
		output output
	}

	mock1, err := pgxmock.NewPool(pgxmock.QueryMatcherOption(pgxmock.QueryMatcherEqual))
	if err != nil {
		t.Error(err)
	}
	mock1.ExpectQuery(
		"select count(id) from menus where id = $1 and establishment_id in (select id from establishments where profile_id = $2)",
	).WithArgs(int64(1), int64(1)).WillReturnRows(
		pgxmock.NewRows([]string{"id"}).AddRow(int64(1)),
	)

	mock2, err := pgxmock.NewPool(pgxmock.QueryMatcherOption(pgxmock.QueryMatcherEqual))
	if err != nil {
		t.Error(err)
	}
	mock2.ExpectQuery(
		"select count(id) from menus where id = $1 and establishment_id in (select id from establishments where profile_id = $2)",
	).WithArgs(int64(1), int64(1)).WillReturnRows(
		pgxmock.NewRows([]string{"id"}).AddRow(int64(0)),
	)

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
				1,
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
				1,
			},
			output{
				false,
				nil,
			},
		},
	}

	for i, el := range cases {
		belongs, err := repo.CheckBelongs(el.input.ctx, el.input.menuID, el.input.profileID)
		if err != nil && el.output.err != nil && err.Error() != el.output.err.Error() {
			t.Error(fmt.Sprintf("error don't match %d, %v", i, err))
		} else if err != nil && el.output.err == nil {
			t.Error("did not expect error: " + err.Error())
		}

		if belongs != el.output.belongs {
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
