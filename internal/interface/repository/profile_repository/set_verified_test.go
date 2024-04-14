package profile_repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/pashagolub/pgxmock/v3"
	"gitlab.com/maometusu/qr_menu/internal/entity"
	"testing"
)

func TestProfileRepository_SetVerified(t *testing.T) {
	repo := NewProfileRepository()

	type input struct {
		ctx      context.Context
		mockDB   pgxmock.PgxPoolIface
		id       int64
		verified bool
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
		"update profiles set verified = $2 where id = $1",
	).WithArgs(int64(1), true).WillReturnResult(pgxmock.NewResult("UPDATE", 1))

	cases := []testCase{
		{
			input{
				context.Background(),
				nil,
				0,
				false,
			},
			output{
				errors.New("connection is missing from context"),
			},
		},
		{
			input{
				context.WithValue(context.Background(), "db", entity.PgxIface(mock)),
				mock,
				1,
				true,
			},
			output{
				nil,
			},
		},
	}

	for i, el := range cases {
		err := repo.SetVerified(el.input.ctx, el.input.id, el.input.verified)
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
