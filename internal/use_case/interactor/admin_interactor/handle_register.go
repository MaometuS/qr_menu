package admin_interactor

import (
	"context"
	"errors"
	"gitlab.com/maometusu/qr_menu/internal/entity/models"
	"golang.org/x/crypto/bcrypt"
	"io"
)

func (a *adminInteractor) HandleRegister(context context.Context, w io.Writer, name, email, password, passRepeat string) error {
	exists, err := a.profileRepository.CheckExistence(context, email)
	if err != nil {
		a.presenter.RegisterPage(w, false, false, true)
		return err
	}

	if exists {
		a.presenter.RegisterPage(w, true, false, false)
		return errors.New("email already exists")
	}

	if password != passRepeat {
		a.presenter.RegisterPage(w, false, true, false)
		return errors.New("passwords don't match")
	}

	passHash, err := bcrypt.GenerateFromPassword([]byte(password), 6)
	if err != nil {
		a.presenter.RegisterPage(w, false, false, true)
		return err
	}

	_, err = a.profileRepository.Create(context, &models.Profile{
		Name:     name,
		Email:    email,
		Password: string(passHash),
	})
	if err != nil {
		a.presenter.RegisterPage(w, false, false, true)
		return err
	}

	return nil
}
