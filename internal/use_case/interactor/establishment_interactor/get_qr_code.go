package establishment_interactor

import (
	"context"
	"io"
)

func (e *establishmentInteractor) GetQrCode(ctx context.Context, w io.Writer, hostname string, establishmentID int64) error {
	establishment, err := e.repository.GetOne(ctx, establishmentID)
	if err != nil {
		return err
	}

	return e.presenter.PresentQrCode(w, hostname+"/e/"+establishment.Link)
}
