package establishment_repository

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

func TestEstablishmentRepository_GetEstablishments(t *testing.T) {
	repo := NewEstablishmentRepository()

	type input struct {
		ctx    context.Context
		mockDB pgxmock.PgxPoolIface
		id     int64
	}

	type output struct {
		establishments []models.Establishment
		err            error
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
		"select * from establishments where profile_id = $1",
	).WithArgs(int64(1)).WillReturnRows(
		pgxmock.NewRows(
			[]string{"id", "name", "phone", "logo", "background", "wifi_password", "can_make_orders", "country", "city", "address", "text", "profile_id", "currency_id", "link"},
		).AddRow(int64(1), "name", "phone", "logo", "background", "password", true, "country", "city", "address", "text", int64(1), int64(1), "link"),
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
				[]models.Establishment{
					{
						1, "name", "phone", "logo", "background", "password", true, "country", "city", "address", "text", 1, 1, "link",
					},
				},
				nil,
			},
		},
	}

	for i, el := range cases {
		res, err := repo.GetEstablishments(el.input.ctx, el.input.id)
		if err != nil && el.output.err != nil && err.Error() != el.output.err.Error() {
			t.Error(fmt.Sprintf("error don't match %d, %v", i, err))
		} else if err != nil && el.output.err == nil {
			t.Error("did not expect error: " + err.Error())
		}

		if !reflect.DeepEqual(res, el.output.establishments) {
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
