package establishment_interactor

import (
	"context"
	"errors"
	"gitlab.com/maometusu/qr_menu/internal/entity/models"
	"gitlab.com/maometusu/qr_menu/internal/use_case/presenter"
	"gitlab.com/maometusu/qr_menu/internal/use_case/repository"
	"testing"
)

func TestEstablishmentInteractor_EstablishmentInfo(t *testing.T) {
	type testCase struct {
		profileID       int64
		establishmentID int64
		rep             *repository.EstablishmentRepositoryMock
		comRep          *repository.CommonRepositoryMock
		pres            *presenter.EstablishmentPresenterMock
		err             string
	}

	cases := make([]testCase, 0)

	rep := &repository.EstablishmentRepositoryMock{}
	rep.On("GetOne", context.Background(), int64(1)).Return(&models.Establishment{}, errors.New("error"))

	cases = append(cases, testCase{
		profileID:       2,
		establishmentID: 1,
		rep:             rep,
		comRep:          nil,
		pres:            nil,
		err:             "error",
	})

	establishment := &models.Establishment{
		ProfileID: 1,
	}
	rep = &repository.EstablishmentRepositoryMock{}
	rep.On("GetOne", context.Background(), int64(1)).Return(establishment, nil)

	cases = append(cases, testCase{
		profileID:       2,
		establishmentID: 1,
		rep:             rep,
		comRep:          nil,
		pres:            nil,
		err:             "not allowed for current profile",
	})

	establishment = &models.Establishment{
		ProfileID: 2,
	}
	rep = &repository.EstablishmentRepositoryMock{}
	rep.On("GetOne", context.Background(), int64(1)).Return(establishment, nil)
	commRep := &repository.CommonRepositoryMock{}
	commRep.On("GetCurrencies", context.Background()).Return(make([]models.Currency, 0), errors.New("error1"))

	cases = append(cases, testCase{
		profileID:       2,
		establishmentID: 1,
		rep:             rep,
		comRep:          commRep,
		pres:            nil,
		err:             "error1",
	})

	establishment = &models.Establishment{
		ProfileID: 2,
	}
	currencies := make([]models.Currency, 0)
	rep = &repository.EstablishmentRepositoryMock{}
	rep.On("GetOne", context.Background(), int64(1)).Return(establishment, nil)
	commRep = &repository.CommonRepositoryMock{}
	commRep.On("GetCurrencies", context.Background()).Return(currencies, nil)
	pres := &presenter.EstablishmentPresenterMock{}
	pres.On("PresentEstablishmentInfo", nil, establishment, currencies).Return(errors.New("error2"))

	cases = append(cases, testCase{
		profileID:       2,
		establishmentID: 1,
		rep:             rep,
		comRep:          commRep,
		pres:            pres,
		err:             "error2",
	})

	establishment = &models.Establishment{
		ProfileID: 2,
	}
	currencies = make([]models.Currency, 0)
	rep = &repository.EstablishmentRepositoryMock{}
	rep.On("GetOne", context.Background(), int64(1)).Return(establishment, nil)
	commRep = &repository.CommonRepositoryMock{}
	commRep.On("GetCurrencies", context.Background()).Return(currencies, nil)
	pres = &presenter.EstablishmentPresenterMock{}
	pres.On("PresentEstablishmentInfo", nil, establishment, currencies).Return(nil)

	cases = append(cases, testCase{
		profileID:       2,
		establishmentID: 1,
		rep:             rep,
		comRep:          commRep,
		pres:            pres,
		err:             "",
	})

	for i, el := range cases {
		interactor := NewEstablishmentInteractor(el.pres, el.rep, nil, nil, el.comRep)
		err := interactor.EstablishmentInfo(context.Background(), nil, el.profileID, el.establishmentID)
		if err != nil && err.Error() != el.err {
			t.Error("errors don't match: ", i, err)
		}

		if el.rep != nil {
			el.rep.AssertExpectations(t)
		}

		if el.comRep != nil {
			el.comRep.AssertExpectations(t)
		}

		if el.pres != nil {
			el.pres.AssertExpectations(t)
		}
	}
}
