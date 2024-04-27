package establishment_interactor

import (
	"context"
	"errors"
	"gitlab.com/maometusu/qr_menu/internal/use_case/repository"
	"testing"
)

func TestEstablishmentInteractor_DoesLinkExist(t *testing.T) {
	type testCase struct {
		link string
		rep  *repository.EstablishmentRepositoryMock
		err  string
		res  bool
	}

	cases := make([]testCase, 0)

	rep := &repository.EstablishmentRepositoryMock{}
	rep.On("DoesLinkExist", context.Background(), "link").Return(false, errors.New("error"))

	cases = append(cases, testCase{
		link: "link",
		rep:  rep,
		err:  "error",
		res:  false,
	})

	rep = &repository.EstablishmentRepositoryMock{}
	rep.On("DoesLinkExist", context.Background(), "link").Return(true, nil)

	cases = append(cases, testCase{
		link: "link",
		rep:  rep,
		err:  "",
		res:  true,
	})

	for i, el := range cases {
		interactor := NewEstablishmentInteractor(nil, el.rep, nil, nil, nil)
		res, err := interactor.DoesLinkExist(context.Background(), el.link)
		if err != nil && err.Error() != el.err {
			t.Error("errors don't match: ", i, err)
		}

		if res != el.res {
			t.Error("results don't match: ", i)
		}

		if el.rep != nil {
			el.rep.AssertExpectations(t)
		}
	}
}
