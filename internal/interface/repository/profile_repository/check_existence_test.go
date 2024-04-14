package profile_repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/pashagolub/pgxmock/v3"
	"gitlab.com/maometusu/qr_menu/internal/entity"
	"testing"
)

func TestProfileRepository_CheckExistence(t *testing.T) {
	repo := NewProfileRepository()

	type input struct {
		ctx    context.Context
		mockDB pgxmock.PgxPoolIface
		email  string
	}

	type output struct {
		exists bool
		err    error
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
		"select count(id) from profiles where email = $1 and verified = true",
	).WithArgs("email").WillReturnRows(pgxmock.NewRows([]string{"count"}).AddRow(int64(1)))

	mock2, err := pgxmock.NewPool(pgxmock.QueryMatcherOption(pgxmock.QueryMatcherEqual))
	if err != nil {
		t.Error(err)
	}
	mock2.ExpectQuery(
		"select count(id) from profiles where email = $1 and verified = true",
	).WithArgs("email").WillReturnRows(pgxmock.NewRows([]string{"count"}).AddRow(int64(0)))

	cases := []testCase{
		{
			input{
				context.Background(),
				nil,
				"",
			},
			output{
				false,
				errors.New("connection is missing from context"),
			},
		},
		{
			input{
				context.WithValue(context.Background(), "db", entity.PgxIface(mock1)),
				mock1,
				"email",
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
				"email",
			},
			output{
				false,
				nil,
			},
		},
	}

	for i, el := range cases {
		belongs, err := repo.CheckExistence(el.input.ctx, el.input.email)
		if err != nil && el.output.err != nil && err.Error() != el.output.err.Error() {
			t.Error(fmt.Sprintf("error don't match %d, %v", i, err))
		} else if err != nil && el.output.err == nil {
			t.Error("did not expect error: " + err.Error())
		}

		if belongs != el.output.exists {
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
