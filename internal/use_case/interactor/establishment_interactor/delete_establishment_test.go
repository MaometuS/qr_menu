package establishment_interactor

import (
	"context"
	"errors"
	"gitlab.com/maometusu/qr_menu/internal/use_case/repository"
	"testing"
)

func TestEstablishmentInteractor_DeleteEstablishment(t *testing.T) {
	type testCase struct {
		profileID  int64
		categoryID int64
		rep        *repository.EstablishmentRepositoryMock
		err        string
	}

	cases := make([]testCase, 0)

	rep := &repository.EstablishmentRepositoryMock{}
	rep.On("DeleteEstablishment", context.Background(), int64(1), int64(2)).Return(errors.New("error"))

	cases = append(cases, testCase{
		profileID:  1,
		categoryID: 2,
		rep:        rep,
		err:        "error",
	})

	rep = &repository.EstablishmentRepositoryMock{}
	rep.On("DeleteEstablishment", context.Background(), int64(1), int64(2)).Return(nil)

	cases = append(cases, testCase{
		profileID:  1,
		categoryID: 2,
		rep:        rep,
		err:        "",
	})

	for i, el := range cases {
		interactor := NewEstablishmentInteractor(nil, el.rep, nil, nil, nil)
		err := interactor.DeleteEstablishment(context.Background(), el.profileID, el.categoryID)
		if err != nil && err.Error() != el.err {
			t.Error("errors don't match: ", i, err)
		}

		if el.rep != nil {
			el.rep.AssertExpectations(t)
		}
	}
}
