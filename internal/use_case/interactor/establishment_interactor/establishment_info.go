package establishment_interactor

import (
	"context"
	"errors"
	"io"
)

func (e *establishmentInteractor) EstablishmentInfo(context context.Context, w io.Writer, profileID, establishmentID int64) error {
	establishment, err := e.repository.GetOne(context, establishmentID)
	if err != nil {
		return err
	}

	if establishment.ProfileID != profileID {
		return errors.New("not allowed for current profile")
	}

	currencies, err := e.commonRepository.GetCurrencies(context)
	if err != nil {
		return err
	}

	return e.presenter.PresentEstablishmentInfo(w, establishment, currencies)
}
