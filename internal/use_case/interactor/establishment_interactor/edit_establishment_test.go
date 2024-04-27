package establishment_interactor

import (
	"context"
	"errors"
	"gitlab.com/maometusu/qr_menu/internal/entity/models"
	"gitlab.com/maometusu/qr_menu/internal/use_case/repository"
	"testing"
)

func TestEstablishmentInteractor_EditEstablishment(t *testing.T) {
	type testCase struct {
		establishment *models.Establishment
		rep           *repository.EstablishmentRepositoryMock
		err           string
	}

	cases := make([]testCase, 0)

	rep := &repository.EstablishmentRepositoryMock{}
	rep.On("GetOne", context.Background(), int64(1)).Return(&models.Establishment{}, errors.New("error"))

	cases = append(cases, testCase{
		establishment: &models.Establishment{ID: 1},
		rep:           rep,
		err:           "error",
	})

	prevEst := &models.Establishment{ProfileID: 1}
	rep = &repository.EstablishmentRepositoryMock{}
	rep.On("GetOne", context.Background(), int64(1)).Return(prevEst, nil)

	cases = append(cases, testCase{
		establishment: &models.Establishment{ID: 1, ProfileID: 2},
		rep:           rep,
		err:           "not allowed for current profile",
	})

	prevEst = &models.Establishment{ProfileID: 1}
	establishment := &models.Establishment{ID: 1, ProfileID: 1}
	rep = &repository.EstablishmentRepositoryMock{}
	rep.On("GetOne", context.Background(), int64(1)).Return(prevEst, nil)
	rep.On("UpdateEstablishment", context.Background(), establishment).Return(errors.New("error1"))

	cases = append(cases, testCase{
		establishment: establishment,
		rep:           rep,
		err:           "error1",
	})

	prevEst = &models.Establishment{ProfileID: 1}
	establishment = &models.Establishment{ID: 1, ProfileID: 1}
	rep = &repository.EstablishmentRepositoryMock{}
	rep.On("GetOne", context.Background(), int64(1)).Return(prevEst, nil)
	rep.On("UpdateEstablishment", context.Background(), establishment).Return(nil)

	cases = append(cases, testCase{
		establishment: establishment,
		rep:           rep,
		err:           "",
	})

	for i, el := range cases {
		interactor := NewEstablishmentInteractor(nil, el.rep, nil, nil, nil)
		err := interactor.EditEstablishment(context.Background(), nil, nil, el.establishment)
		if err != nil && err.Error() != el.err {
			t.Error("errors don't match: ", i, err)
		}

		if el.rep != nil {
			el.rep.AssertExpectations(t)
		}
	}
}
