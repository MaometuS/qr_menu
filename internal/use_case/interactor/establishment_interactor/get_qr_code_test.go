package establishment_interactor

import (
	"context"
	"errors"
	"gitlab.com/maometusu/qr_menu/internal/entity/models"
	"gitlab.com/maometusu/qr_menu/internal/use_case/presenter"
	"gitlab.com/maometusu/qr_menu/internal/use_case/repository"
	"testing"
)

func TestEstablishmentInteractor_GetQrCode(t *testing.T) {
	type testCase struct {
		hostname        string
		establishmentID int64
		rep             *repository.EstablishmentRepositoryMock
		pres            *presenter.EstablishmentPresenterMock
		err             string
	}

	cases := make([]testCase, 0)

	rep := &repository.EstablishmentRepositoryMock{}
	rep.On("GetOne", context.Background(), int64(1)).Return(&models.Establishment{}, errors.New("error"))

	cases = append(cases, testCase{
		hostname:        "link",
		establishmentID: 1,
		rep:             rep,
		pres:            nil,
		err:             "error",
	})

	establishment := &models.Establishment{Link: "link"}
	rep = &repository.EstablishmentRepositoryMock{}
	rep.On("GetOne", context.Background(), int64(1)).Return(establishment, nil)
	pres := &presenter.EstablishmentPresenterMock{}
	pres.On("PresentQrCode", nil, "host/e/link").Return(errors.New("error1"))

	cases = append(cases, testCase{
		hostname:        "host",
		establishmentID: 1,
		rep:             rep,
		pres:            pres,
		err:             "error1",
	})

	establishment = &models.Establishment{Link: "link"}
	rep = &repository.EstablishmentRepositoryMock{}
	rep.On("GetOne", context.Background(), int64(1)).Return(establishment, nil)
	pres = &presenter.EstablishmentPresenterMock{}
	pres.On("PresentQrCode", nil, "host/e/link").Return(nil)

	cases = append(cases, testCase{
		hostname:        "host",
		establishmentID: 1,
		rep:             rep,
		pres:            pres,
		err:             "",
	})

	for i, el := range cases {
		interactor := NewEstablishmentInteractor(el.pres, el.rep, nil, nil, nil)
		err := interactor.GetQrCode(context.Background(), nil, el.hostname, el.establishmentID)
		if err != nil && err.Error() != el.err {
			t.Error("errors don't match: ", i, err)
		}

		if el.rep != nil {
			el.rep.AssertExpectations(t)
		}

		if el.pres != nil {
			el.pres.AssertExpectations(t)
		}
	}
}
