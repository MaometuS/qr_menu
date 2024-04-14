package menu_repository

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

func TestMenuRepository_GetOne(t *testing.T) {
	repo := NewMenuRepository()

	type input struct {
		ctx    context.Context
		mockDB pgxmock.PgxPoolIface
		id     int64
	}

	type output struct {
		menu *models.Menu
		err  error
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
		"select * from menus where id = $1",
	).WithArgs(int64(1)).WillReturnRows(
		pgxmock.NewRows(
			[]string{"id", "name", "is_visible", "establishment_id", "order_param"},
		).AddRow(int64(1), "name", true, int64(1), int64(1)),
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
				&models.Menu{
					ID:              1,
					Name:            "name",
					IsVisible:       true,
					EstablishmentID: 1,
					OrderParam:      1,
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

		if !reflect.DeepEqual(res, el.output.menu) {
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
