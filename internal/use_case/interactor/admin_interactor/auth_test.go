package admin_interactor

import (
	"context"
	"gitlab.com/maometusu/qr_menu/internal/entity"
	"testing"
)

func TestAdminInteractor_Auth(t *testing.T) {
	type testCase struct {
		ctx         context.Context
		tokenString string
		userID      int64
		err         string
		jwtSecret   string
	}

	cases := []testCase{
		{
			ctx:         context.Background(),
			tokenString: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.3NGZqf-6q9ZoM95B0EvF4Nu0No-Z50dFvnPs51bqKXI",
			userID:      0,
			err:         "token signature is invalid: signature is invalid",
			jwtSecret:   "secret",
		},
		{
			ctx:         context.Background(),
			tokenString: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiSm9obiBEb2UiLCJpYXQiOjE1MTYyMzkwMjJ9.8nYFUX869Y1mnDDDU4yL11aANgVRuifoxrE8BHZY1iE",
			userID:      0,
			err:         "cannot find id",
			jwtSecret:   "secret",
		},
		{
			ctx:         context.Background(),
			tokenString: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOjEsIm5hbWUiOiJKb2huIERvZSIsImlhdCI6MTUxNjIzOTAyMn0.L7T2OCycF4IAnj2GxQPPD7_OBDsK0QwKONL7INO1eEI",
			userID:      1,
			err:         "",
			jwtSecret:   "secret",
		},
	}

	for i, el := range cases {
		interactor := NewAdminInteractor(nil, nil, nil, &entity.Config{JWTSignString: el.jwtSecret})
		userID, err := interactor.Auth(el.ctx, el.tokenString)
		if err != nil && err.Error() != el.err {
			t.Error("errors don't match: ", i, err)
		}

		if el.userID != userID {
			t.Error("user id don't match: ", i, "expected", el.userID, "got", userID)
		}
	}
}
