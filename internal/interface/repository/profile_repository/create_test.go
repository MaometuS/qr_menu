package profile_repository

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

func TestProfileRepository_Create(t *testing.T) {
	repo := NewProfileRepository()

	type input struct {
		ctx     context.Context
		mockDB  pgxmock.PgxPoolIface
		profile *models.Profile
	}

	type output struct {
		id  int64
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
		"insert into profiles(name, email, password) values ($1, $2, $3) returning id",
	).WithArgs("name", "email", "password").WillReturnRows(pgxmock.NewRows([]string{"id"}).AddRow(int64(1)))

	cases := []testCase{
		{
			input{
				context.Background(),
				nil,
				nil,
			},
			output{
				0,
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
				1,
				nil,
			},
		},
	}

	for i, el := range cases {
		res, err := repo.Create(el.input.ctx, el.input.profile)
		if err != nil && el.output.err != nil && err.Error() != el.output.err.Error() {
			t.Error(fmt.Sprintf("error don't match %d, %v", i, err))
		} else if err != nil && el.output.err == nil {
			t.Error("did not expect error: " + err.Error())
		}

		if !reflect.DeepEqual(res, el.output.id) {
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
