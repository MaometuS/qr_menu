package admin_interactor

import (
	"context"
	"errors"
	"gitlab.com/maometusu/qr_menu/internal/entity/models"
	"golang.org/x/crypto/bcrypt"
	"io"
)

func (a *adminInteractor) HandleRegister(context context.Context, w io.Writer, name, email, password, passRepeat string) (int64, error) {
	exists, err := a.profileRepository.CheckExistence(context, email)
	if err != nil {
		a.presenter.RegisterPage(w, false, false, true)
		return 0, err
	}

	if exists {
		a.presenter.RegisterPage(w, true, false, false)
		return 0, errors.New("email already exists")
	}

	if password != passRepeat {
		a.presenter.RegisterPage(w, false, true, false)
		return 0, errors.New("passwords don't match")
	}

	passHash, err := bcrypt.GenerateFromPassword([]byte(password), 6)
	if err != nil {
		a.presenter.RegisterPage(w, false, false, true)
		return 0, err
	}

	id, err := a.profileRepository.Create(context, &models.Profile{
		Name:     name,
		Email:    email,
		Password: string(passHash),
	})
	if err != nil {
		a.presenter.RegisterPage(w, false, false, true)
		return 0, err
	}

	return id, nil
}
