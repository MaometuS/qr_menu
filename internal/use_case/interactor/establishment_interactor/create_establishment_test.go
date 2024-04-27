package establishment_interactor

import (
	"context"
	"errors"
	"gitlab.com/maometusu/qr_menu/internal/entity/models"
	"gitlab.com/maometusu/qr_menu/internal/use_case/repository"
	"testing"
)

func TestEstablishmentInteractor_CreateEstablishment(t *testing.T) {
	type testCase struct {
		establishment *models.Establishment
		rep           *repository.EstablishmentRepositoryMock
		err           string
	}

	cases := make([]testCase, 0)

	establishment := &models.Establishment{}
	rep := &repository.EstablishmentRepositoryMock{}
	rep.On("CreateEstablishment", context.Background(), establishment).Return(errors.New("error"))

	cases = append(cases, testCase{
		establishment: establishment,
		rep:           rep,
		err:           "error",
	})

	establishment = &models.Establishment{}
	rep = &repository.EstablishmentRepositoryMock{}
	rep.On("CreateEstablishment", context.Background(), establishment).Return(nil)

	cases = append(cases, testCase{
		establishment: establishment,
		rep:           rep,
		err:           "",
	})

	for i, el := range cases {
		interactor := NewEstablishmentInteractor(nil, el.rep, nil, nil, nil)
		err := interactor.CreateEstablishment(context.Background(), el.establishment)
		if err != nil && err.Error() != el.err {
			t.Error("errors don't match: ", i, err)
		}

		if el.rep != nil {
			el.rep.AssertExpectations(t)
		}
	}
}
