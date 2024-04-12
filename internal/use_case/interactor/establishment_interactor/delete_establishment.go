package establishment_interactor

import (
	"context"
)

func (e *establishmentInteractor) DeleteEstablishment(context context.Context, profileID, establishmentID int64) error {
	return e.repository.DeleteEstablishment(context, profileID, establishmentID)
}
