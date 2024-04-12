package admin_interactor

import (
	"context"
	"errors"
)

func (a *adminInteractor) EditEmail(ctx context.Context, email, verificationCode string, id int64) error {
	profile, err := a.profileRepository.GetOne(ctx, id)
	if err != nil {
		return err
	}

	if profile.VerificationCode != verificationCode {
		return errors.New("wrong verification code")
	}

	profile.Email = email
	err = a.profileRepository.Update(ctx, profile)
	if err != nil {
		return err
	}

	return nil
}
