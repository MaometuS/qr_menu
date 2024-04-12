package establishment_interactor

import (
	"context"
	"errors"
	"io"
)

func (e *establishmentInteractor) EditEstablishmentPage(context context.Context, w io.Writer, profileID int64, link string) error {
	establishment, err := e.repository.GetByLink(context, link)
	if err != nil {
		return err
	}

	if establishment.ProfileID != profileID {
		return errors.New("not an owner")
	}

	topMenuID, _ := e.menuRepository.GetTopID(context, establishment.ID)

	return e.presenter.PresentEditEstablishmentPage(w, establishment, topMenuID)
}
