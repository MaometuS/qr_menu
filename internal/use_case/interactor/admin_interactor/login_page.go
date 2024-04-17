package admin_interactor

import (
	"context"
	"io"
)

func (a *adminInteractor) LoginPage(context context.Context, w io.Writer) error {
	return a.presenter.LoginPage(w, false, false)
}
