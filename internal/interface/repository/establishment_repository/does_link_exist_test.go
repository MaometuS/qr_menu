package establishment_repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/pashagolub/pgxmock/v3"
	"gitlab.com/maometusu/qr_menu/internal/entity"
	"testing"
)

func TestEstablishmentRepository_DoesLinkExist(t *testing.T) {
	repo := NewEstablishmentRepository()

	type input struct {
		ctx    context.Context
		mockDB pgxmock.PgxPoolIface
		link   string
	}

	type output struct {
		exists bool
		err    error
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
		"select count(id) from establishments where link = $1",
	).WithArgs("link").WillReturnRows(pgxmock.NewRows([]string{"count"}).AddRow(int64(1)))

	mock2, err := pgxmock.NewPool(pgxmock.QueryMatcherOption(pgxmock.QueryMatcherEqual))
	if err != nil {
		t.Error(err)
	}
	mock2.ExpectQuery(
		"select count(id) from establishments where link = $1",
	).WithArgs("link").WillReturnRows(pgxmock.NewRows([]string{"count"}).AddRow(int64(0)))

	cases := []testCase{
		{
			input{
				context.Background(),
				nil,
				"",
			},
			output{
				false,
				errors.New("db not found in context"),
			},
		},
		{
			input{
				context.WithValue(context.Background(), "db", entity.PgxIface(mock)),
				mock,
				"link",
			},
			output{
				true,
				nil,
			},
		},
		{
			input{
				context.WithValue(context.Background(), "db", entity.PgxIface(mock2)),
				mock2,
				"link",
			},
			output{
				false,
				nil,
			},
		},
	}

	for i, el := range cases {
		exists, err := repo.DoesLinkExist(el.input.ctx, el.input.link)
		if err != nil && el.output.err != nil && err.Error() != el.output.err.Error() {
			t.Error(fmt.Sprintf("error don't match %d, %v", i, err))
		} else if err != nil && el.output.err == nil {
			t.Error("did not expect error: " + err.Error())
		}

		if exists != el.output.exists {
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
