package admin_interactor

import (
	"context"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func (a *adminInteractor) HandleLogin(context context.Context, email, password string) (string, error) {
	profile, err := a.profileRepository.GetOneByEmail(context, email)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(profile.Password), []byte(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": "qrmenu.uz",
		"sub": profile.ID,
	})

	tokenString, err := token.SignedString([]byte("asldfjof3982vu42oj3kj"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
