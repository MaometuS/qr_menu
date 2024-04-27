package establishment_interactor

import (
	"context"
	"errors"
	"gitlab.com/maometusu/qr_menu/internal/entity/models"
	"gitlab.com/maometusu/qr_menu/internal/use_case/presenter"
	"gitlab.com/maometusu/qr_menu/internal/use_case/repository"
	"testing"
)

func TestEstablishmentInteractor_GetEstablishments(t *testing.T) {
	type testCase struct {
		profileID int64
		rep       *repository.EstablishmentRepositoryMock
		comRep    *repository.CommonRepositoryMock
		pres      *presenter.EstablishmentPresenterMock
		err       string
	}

	cases := make([]testCase, 0)

	rep := &repository.EstablishmentRepositoryMock{}
	rep.On("GetEstablishments", context.Background(), int64(1)).Return(make([]models.Establishment, 0), errors.New("error"))

	cases = append(cases, testCase{
		profileID: 1,
		rep:       rep,
		comRep:    nil,
		pres:      nil,
		err:       "error",
	})

	establishments := make([]models.Establishment, 0)
	rep = &repository.EstablishmentRepositoryMock{}
	rep.On("GetEstablishments", context.Background(), int64(1)).Return(establishments, nil)
	comRep := &repository.CommonRepositoryMock{}
	comRep.On("GetCurrencies", context.Background()).Return(make([]models.Currency, 0), errors.New("error1"))

	cases = append(cases, testCase{
		profileID: 1,
		rep:       rep,
		comRep:    comRep,
		pres:      nil,
		err:       "error1",
	})

	establishments = make([]models.Establishment, 0)
	currencies := make([]models.Currency, 0)
	rep = &repository.EstablishmentRepositoryMock{}
	rep.On("GetEstablishments", context.Background(), int64(1)).Return(establishments, nil)
	comRep = &repository.CommonRepositoryMock{}
	comRep.On("GetCurrencies", context.Background()).Return(currencies, nil)
	pres := &presenter.EstablishmentPresenterMock{}
	pres.On("PresentEstablishments", nil, establishments, currencies).Return(errors.New("error2"))

	cases = append(cases, testCase{
		profileID: 1,
		rep:       rep,
		comRep:    comRep,
		pres:      pres,
		err:       "error2",
	})

	establishments = make([]models.Establishment, 0)
	currencies = make([]models.Currency, 0)
	rep = &repository.EstablishmentRepositoryMock{}
	rep.On("GetEstablishments", context.Background(), int64(1)).Return(establishments, nil)
	comRep = &repository.CommonRepositoryMock{}
	comRep.On("GetCurrencies", context.Background()).Return(currencies, nil)
	pres = &presenter.EstablishmentPresenterMock{}
	pres.On("PresentEstablishments", nil, establishments, currencies).Return(nil)

	cases = append(cases, testCase{
		profileID: 1,
		rep:       rep,
		comRep:    comRep,
		pres:      pres,
		err:       "",
	})

	for i, el := range cases {
		interactor := NewEstablishmentInteractor(el.pres, el.rep, nil, nil, el.comRep)
		err := interactor.GetEstablishments(context.Background(), nil, el.profileID)
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
