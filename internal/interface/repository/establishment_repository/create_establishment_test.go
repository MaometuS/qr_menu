package establishment_repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/pashagolub/pgxmock/v3"
	"gitlab.com/maometusu/qr_menu/internal/entity"
	"gitlab.com/maometusu/qr_menu/internal/entity/models"
	"testing"
)

func TestEstablishmentRepository_CreateEstablishment(t *testing.T) {
	repo := NewEstablishmentRepository()

	type input struct {
		ctx           context.Context
		mockDB        pgxmock.PgxPoolIface
		establishment *models.Establishment
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
		"insert into establishments(name, profile_id, currency_id, link) values ($1, $2, $3, $4)",
	).WithArgs(
		"name", int64(1), int64(1), "link",
	).WillReturnResult(pgxmock.NewResult("INSERT", 1))

	cases := []testCase{
		{
			input{
				context.Background(),
				nil,
				nil,
			},
			output{
				errors.New("db not found in context"),
			},
		},
		{
			input{
				context.WithValue(context.Background(), "db", entity.PgxIface(mock)),
				mock,
				&models.Establishment{
					Name:       "name",
					ProfileID:  1,
					CurrencyID: 1,
					Link:       "link",
				},
			},
			output{
				nil,
			},
		},
	}

	for i, el := range cases {
		err := repo.CreateEstablishment(el.input.ctx, el.input.establishment)
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
