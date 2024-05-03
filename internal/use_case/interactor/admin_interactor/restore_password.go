package admin_interactor

import (
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"io"
)

func (a *adminInteractor) RestorePassword(ctx context.Context, w io.Writer, email, pass, passRepeat string) (int64, error) {
	prof, err := a.profileRepository.GetOneByEmail(ctx, email)
	if err != nil {
		a.presenter.RestorePasswordPage(w, true, false, false)
		return 0, err
	}

	if pass != passRepeat {
		a.presenter.RestorePasswordPage(w, false, true, false)
		return 0, errors.New("passwords don't match")
	}

	passHash, err := bcrypt.GenerateFromPassword([]byte(pass), 6)
	if err != nil {
		a.presenter.RestorePasswordPage(w, false, false, true)
		return 0, err
	}

	prof.NewPassword = string(passHash)
	err = a.profileRepository.Update(ctx, prof)
	if err != nil {
		a.presenter.RestorePasswordPage(w, false, false, true)
		return 0, err
	}

	return prof.ID, nil
}
