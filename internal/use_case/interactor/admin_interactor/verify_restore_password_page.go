package admin_interactor

import (
	"context"
	"io"
)

func (a *adminInteractor) VerifyRestorePasswordPage(ctx context.Context, w io.Writer, id int64) error {
	return a.presenter.VerifyRestorePasswordPage(w, id, false, false)
}
