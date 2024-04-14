package menu_repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/pashagolub/pgxmock/v3"
	"gitlab.com/maometusu/qr_menu/internal/entity"
	"gitlab.com/maometusu/qr_menu/internal/entity/models"
	"testing"
)

func TestMenuRepository_CreateMenu(t *testing.T) {
	repo := NewMenuRepository()

	type input struct {
		ctx    context.Context
		mockDB pgxmock.PgxPoolIface
		menu   *models.Menu
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
		"update menus set order_param=order_param+1 where order_param >= $1 and establishment_id = $2",
	).WithArgs(int64(1), int64(1)).WillReturnResult(pgxmock.NewResult("UPDATE", 1))

	mock.ExpectExec(
		"insert into menus(name, is_visible, establishment_id, order_param) values ($1, $2, $3, $4)",
	).WithArgs("name", true, int64(1), int64(1)).WillReturnResult(pgxmock.NewResult("INSERT", 1))

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
				&models.Menu{
					Name:            "name",
					IsVisible:       true,
					EstablishmentID: 1,
					OrderParam:      1,
				},
			},
			output{
				nil,
			},
		},
	}

	for i, el := range cases {
		err := repo.CreateMenu(el.input.ctx, el.input.menu)
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
