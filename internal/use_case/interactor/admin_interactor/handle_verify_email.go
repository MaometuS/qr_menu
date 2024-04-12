package admin_interactor

import (
	"context"
	"errors"
	"github.com/golang-jwt/jwt/v5"
)

func (a *adminInteractor) HandleVerifyEmail(context context.Context, id int64, verificationCode string) (string, error) {
	profile, err := a.profileRepository.GetOne(context, id)
	if err != nil {
		return "", err
	}

	if profile.VerificationCode != verificationCode {
		return "", errors.New("code does not match")
	}

	err = a.profileRepository.SetVerified(context, id, true)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": "qrmenu.uz",
		"sub": id,
	})

	tokenString, err := token.SignedString([]byte("asldfjof3982vu42oj3kj"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
