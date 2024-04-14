package establishment_repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/pashagolub/pgxmock/v3"
	"gitlab.com/maometusu/qr_menu/internal/entity"
	"testing"
)

func TestEstablishmentRepository_DeleteEstablishment(t *testing.T) {
	repo := NewEstablishmentRepository()

	type input struct {
		ctx             context.Context
		mockDB          pgxmock.PgxPoolIface
		profileID       int64
		establishmentID int64
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
		"delete from establishments where id = $1 and profile_id = $2",
	).WithArgs(int64(1), int64(1)).WillReturnResult(pgxmock.NewResult("DELETE", 1))

	cases := []testCase{
		{
			input{
				context.Background(),
				nil,
				0,
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
				1,
			},
			output{
				nil,
			},
		},
	}

	for i, el := range cases {
		err := repo.DeleteEstablishment(el.input.ctx, el.input.profileID, el.input.establishmentID)
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
