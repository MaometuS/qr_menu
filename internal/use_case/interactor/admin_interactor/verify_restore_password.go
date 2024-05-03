package admin_interactor

import (
	"context"
	"errors"
	"io"
)

func (a *adminInteractor) VerifyRestorePassword(ctx context.Context, w io.Writer, id int64, verificationCode string) error {
	prof, err := a.profileRepository.GetOne(ctx, id)
	if err != nil {
		a.presenter.VerifyRestorePasswordPage(w, id, false, true)
		return err
	}

	if prof.VerificationCode != verificationCode {
		a.presenter.VerifyRestorePasswordPage(w, id, true, false)
		return errors.New("codes don't match")
	}

	prof.Password = prof.NewPassword
	err = a.profileRepository.Update(ctx, prof)
	if err != nil {
		a.presenter.VerifyRestorePasswordPage(w, id, false, true)
		return err
	}

	return nil
}
