package establishment_interactor

import (
	"context"
	"io"
)

func (e *establishmentInteractor) GetEstablishments(context context.Context, w io.Writer, profileID int64) error {
	establishments, err := e.repository.GetEstablishments(context, profileID)
	if err != nil {
		return err
	}

	currencies, err := e.commonRepository.GetCurrencies(context)
	if err != nil {
		return err
	}

	err = e.presenter.PresentEstablishments(w, establishments, currencies)
	if err != nil {
		return err
	}

	return nil
}
