package establishment_interactor

import (
	"context"
	"errors"
	"gitlab.com/maometusu/qr_menu/internal/entity/models"
	"gitlab.com/maometusu/qr_menu/internal/use_case/presenter"
	"gitlab.com/maometusu/qr_menu/internal/use_case/repository"
	"testing"
)

func TestEstablishmentInteractor_EditEstablishmentPage(t *testing.T) {
	type testCase struct {
		profileID int64
		link      string
		rep       *repository.EstablishmentRepositoryMock
		menuRep   *repository.MenuRepositoryMock
		pres      *presenter.EstablishmentPresenterMock
		err       string
	}

	cases := make([]testCase, 0)

	rep := &repository.EstablishmentRepositoryMock{}
	rep.On("GetByLink", context.Background(), "link").Return(&models.Establishment{}, errors.New("error"))

	cases = append(cases, testCase{
		profileID: 1,
		link:      "link",
		rep:       rep,
		menuRep:   nil,
		pres:      nil,
		err:       "error",
	})

	establishment := &models.Establishment{
		ProfileID: 2,
	}
	rep = &repository.EstablishmentRepositoryMock{}
	rep.On("GetByLink", context.Background(), "link").Return(establishment, nil)

	cases = append(cases, testCase{
		profileID: 1,
		link:      "link",
		rep:       rep,
		menuRep:   nil,
		pres:      nil,
		err:       "not an owner",
	})

	establishment = &models.Establishment{
		ID:        5,
		ProfileID: 1,
	}
	rep = &repository.EstablishmentRepositoryMock{}
	rep.On("GetByLink", context.Background(), "link").Return(establishment, nil)
	menuRep := &repository.MenuRepositoryMock{}
	menuRep.On("GetTopID", context.Background(), int64(5)).Return(4, nil)
	pres := &presenter.EstablishmentPresenterMock{}
	pres.On("PresentEditEstablishmentPage", nil, establishment, int64(4)).Return(errors.New("error1"))

	cases = append(cases, testCase{
		profileID: 1,
		link:      "link",
		rep:       rep,
		menuRep:   menuRep,
		pres:      pres,
		err:       "error1",
	})

	establishment = &models.Establishment{
		ID:        5,
		ProfileID: 1,
	}
	rep = &repository.EstablishmentRepositoryMock{}
	rep.On("GetByLink", context.Background(), "link").Return(establishment, nil)
	menuRep = &repository.MenuRepositoryMock{}
	menuRep.On("GetTopID", context.Background(), int64(5)).Return(4, nil)
	pres = &presenter.EstablishmentPresenterMock{}
	pres.On("PresentEditEstablishmentPage", nil, establishment, int64(4)).Return(nil)

	cases = append(cases, testCase{
		profileID: 1,
		link:      "link",
		rep:       rep,
		menuRep:   menuRep,
		pres:      pres,
		err:       "",
	})

	for i, el := range cases {
		interactor := NewEstablishmentInteractor(el.pres, el.rep, el.menuRep, nil, nil)
		err := interactor.EditEstablishmentPage(context.Background(), nil, el.profileID, el.link)
		if err != nil && err.Error() != el.err {
			t.Error("errors don't match: ", i, err)
		}

		if el.rep != nil {
			el.rep.AssertExpectations(t)
		}

		if el.menuRep != nil {
			el.menuRep.AssertExpectations(t)
		}

		if el.pres != nil {
			el.pres.AssertExpectations(t)
		}
	}
}
