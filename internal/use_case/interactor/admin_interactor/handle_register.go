package admin_interactor

import (
	"context"
	"errors"
	"gitlab.com/maometusu/qr_menu/internal/entity/models"
	"golang.org/x/crypto/bcrypt"
)

func (a *adminInteractor) HandleRegister(context context.Context, name, email, password string) error {
	exists, err := a.profileRepository.CheckExistence(context, email)
	if err != nil {
		return err
	}

	if exists {
		return errors.New("user already exists")
	}

	passHash, err := bcrypt.GenerateFromPassword([]byte(password), 6)
	if err != nil {
		return err
	}

	_, err = a.profileRepository.Create(context, &models.Profile{
		Name:     name,
		Email:    email,
		Password: string(passHash),
	})
	if err != nil {
		return err
	}

	return nil
}
