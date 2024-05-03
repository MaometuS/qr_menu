package admin_interactor

import (
	"context"
	"io"
)

func (a *adminInteractor) RestorePasswordPage(ctx context.Context, w io.Writer) error {
	return a.presenter.RestorePasswordPage(w, false, false, false)
}
