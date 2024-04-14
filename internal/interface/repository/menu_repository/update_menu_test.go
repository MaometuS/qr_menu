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

func TestMenuRepository_UpdateMenu(t *testing.T) {
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
		"update menus set name = $1, is_visible = $2, establishment_id = $3 where id = $4",
	).WithArgs("name", true, int64(1), int64(1)).WillReturnResult(pgxmock.NewResult("UPDATE", 1))

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
					ID:              1,
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
		err := repo.UpdateMenu(el.input.ctx, el.input.menu)
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
