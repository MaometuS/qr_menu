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

func TestEstablishmentRepository_UpdateEstablishment(t *testing.T) {
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
		`update establishments set 
				name = $1, 
				phone = $2, 
				wifi_password = $3, 
				can_make_orders = $4,
				country = $5,
				city = $6,
				address = $7,
				text = $8,
				currency_id = $9
			where id = $10`,
	).WithArgs(
		"name", "phone", "wifi_password", true, "country", "city", "address", "text", int64(1), int64(1),
	).WillReturnResult(pgxmock.NewResult("UPDATE", 1))

	mock.ExpectExec(
		"update establishments set logo = $1 where id = $2",
	).WithArgs("logo", int64(1)).WillReturnResult(pgxmock.NewResult("UPDATE", 1))

	mock.ExpectExec(
		"update establishments set background = $1 where id = $2",
	).WithArgs("background", int64(1)).WillReturnResult(pgxmock.NewResult("UPDATE", 1))

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
					ID:            1,
					Name:          "name",
					Phone:         "phone",
					Logo:          "logo",
					Background:    "background",
					WifiPassword:  "wifi_password",
					CanMakeOrders: true,
					Country:       "country",
					City:          "city",
					Address:       "address",
					Text:          "text",
					CurrencyID:    1,
				},
			},
			output{
				nil,
			},
		},
	}

	for i, el := range cases {
		err := repo.UpdateEstablishment(el.input.ctx, el.input.establishment)
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
