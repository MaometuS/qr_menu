package admin_interactor

import (
	"context"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"io"
)

func (a *adminInteractor) HandleVerifyEmail(context context.Context, w io.Writer, id int64, verificationCode string) (string, error) {
	profile, err := a.profileRepository.GetOne(context, id)
	if err != nil {
		a.presenter.VerifyPage(w, id, false, true)
		return "", err
	}

	if profile.VerificationCode != verificationCode {
		a.presenter.VerifyPage(w, id, true, false)
		return "", errors.New("code does not match")
	}

	err = a.profileRepository.SetVerified(context, id, true)
	if err != nil {
		a.presenter.VerifyPage(w, id, false, true)
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": "qrmenu.uz",
		"sub": id,
	})

	tokenString, err := token.SignedString([]byte(a.config.JWTSignString))
	if err != nil {
		a.presenter.VerifyPage(w, id, false, true)
		return "", err
	}

	return tokenString, nil
}
