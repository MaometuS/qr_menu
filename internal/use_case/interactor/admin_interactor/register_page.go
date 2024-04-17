package admin_interactor

import (
	"context"
	"io"
)

func (a *adminInteractor) RegisterPage(context context.Context, w io.Writer) error {
	return a.presenter.RegisterPage(w, false, false, false)
}
