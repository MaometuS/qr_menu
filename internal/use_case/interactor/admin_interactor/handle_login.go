package admin_interactor

import (
	"context"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"io"
)

func (a *adminInteractor) HandleLogin(context context.Context, w io.Writer, email, password string) (string, error) {
	profile, err := a.profileRepository.GetOneByEmail(context, email)
	if err != nil {
		a.presenter.LoginPage(w, true, false)
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(profile.Password), []byte(password))
	if err != nil {
		a.presenter.LoginPage(w, true, false)
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": "qrmenu.uz",
		"sub": profile.ID,
	})

	tokenString, err := token.SignedString([]byte(a.config.JWTSignString))
	if err != nil {
		a.presenter.LoginPage(w, false, true)
		return "", err
	}

	return tokenString, nil
}
