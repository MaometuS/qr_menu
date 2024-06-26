package admin_interactor

import (
	"context"
	"errors"
	"github.com/golang-jwt/jwt/v5"
)

func (a *adminInteractor) Auth(context context.Context, tokenString string) (int64, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}

		return []byte(a.config.JWTSignString), nil
	})
	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		id, ok := claims["sub"].(float64)
		if !ok {
			return 0, errors.New("cannot find id")
		}

		prof, err := a.profileRepository.GetOne(context, int64(id))
		if err != nil {
			return 0, err
		}

		if !prof.Verified {
			return 0, errors.New("user unverified")
		}

		return int64(id), nil
	} else {
		return 0, errors.New("jwt claims are not valid")
	}
}
