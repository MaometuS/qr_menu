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

func TestProfileRepository_GetOne(t *testing.T) {
	repo := NewProfileRepository()

	type input struct {
		ctx    context.Context
		mockDB pgxmock.PgxPoolIface
		id     int64
	}

	type output struct {
		profile *models.Profile
		err     error
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
		"select * from profiles where id = $1",
	).WithArgs(int64(1)).WillReturnRows(
		pgxmock.NewRows(
			[]string{"id", "name", "email", "password", "new_password", "verified", "verification_code"},
		).AddRow(int64(1), "name", "email", "password", "new_password", true, "code"),
	)

	cases := []testCase{
		{
			input{
				context.Background(),
				nil,
				0,
			},
			output{
				nil,
				errors.New("connection is missing from context"),
			},
		},
		{
			input{
				context.WithValue(context.Background(), "db", entity.PgxIface(mock)),
				mock,
				1,
			},
			output{
				&models.Profile{
					ID:               1,
					Name:             "name",
					Email:            "email",
					Password:         "password",
					NewPassword:      "new_password",
					Verified:         true,
					VerificationCode: "code",
				},
				nil,
			},
		},
	}

	for i, el := range cases {
		res, err := repo.GetOne(el.input.ctx, el.input.id)
		if err != nil && el.output.err != nil && err.Error() != el.output.err.Error() {
			t.Error(fmt.Sprintf("error don't match %d, %v", i, err))
		} else if err != nil && el.output.err == nil {
			t.Error("did not expect error: " + err.Error())
		}

		if !reflect.DeepEqual(res, el.output.profile) {
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
