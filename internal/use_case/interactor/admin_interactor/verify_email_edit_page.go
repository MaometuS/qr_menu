package admin_interactor

import (
	"context"
	"io"
)

func (a *adminInteractor) VerifyEmailEditPage(ctx context.Context, w io.Writer, id int64, email string) error {
	err := sendVerificationCode(ctx, a.profileRepository, a.emailRepository, id, email)
	if err != nil {
		return err
	}

	return a.presenter.VerifyEmailEditPage(w, email)
}
