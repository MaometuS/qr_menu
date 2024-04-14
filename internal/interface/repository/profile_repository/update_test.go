package profile_repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/pashagolub/pgxmock/v3"
	"gitlab.com/maometusu/qr_menu/internal/entity"
	"gitlab.com/maometusu/qr_menu/internal/entity/models"
	"testing"
)

func TestProfileRepository_Update(t *testing.T) {
	repo := NewProfileRepository()

	type input struct {
		ctx     context.Context
		mockDB  pgxmock.PgxPoolIface
		profile *models.Profile
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
		"update profiles set name = $1, email = $2, password = $3",
	).WithArgs("name", "email", "password").WillReturnResult(pgxmock.NewResult("UPDATE", 1))

	cases := []testCase{
		{
			input{
				context.Background(),
				nil,
				nil,
			},
			output{
				errors.New("connection is missing from context"),
			},
		},
		{
			input{
				context.WithValue(context.Background(), "db", entity.PgxIface(mock)),
				mock,
				&models.Profile{
					Name:     "name",
					Email:    "email",
					Password: "password",
				},
			},
			output{
				nil,
			},
		},
	}

	for i, el := range cases {
		err := repo.Update(el.input.ctx, el.input.profile)
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
