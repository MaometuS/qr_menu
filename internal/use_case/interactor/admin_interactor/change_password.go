package admin_interactor

import (
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

func (a *adminInteractor) ChangePassword(ctx context.Context, password, passRepeat string, id int64) error {
	if password != passRepeat {
		return errors.New("passwords don't match")
	}

	profile, err := a.profileRepository.GetOne(ctx, id)
	if err != nil {
		return err
	}

	passHash, err := bcrypt.GenerateFromPassword([]byte(password), 6)
	if err != nil {
		return err
	}

	profile.Password = string(passHash)

	err = a.profileRepository.Update(ctx, profile)
	if err != nil {
		return err
	}

	return nil
}
