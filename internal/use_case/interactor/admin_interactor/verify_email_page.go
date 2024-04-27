package admin_interactor

import (
	"context"
	"io"
)

func (a *adminInteractor) VerifyEmailPage(context context.Context, w io.Writer, id int64) error {
	err := sendVerificationCode(context, a.profileRepository, a.emailRepository, id, "")
	if err != nil {
		return err
	}

	err = a.presenter.VerifyPage(w, id, false, false)
	if err != nil {
		return err
	}

	return nil
}
