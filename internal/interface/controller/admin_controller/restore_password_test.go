package admin_controller

import (
	"context"
	"errors"
	"gitlab.com/maometusu/qr_menu/internal/use_case/interactor/admin_interactor"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestController_RestorePassword(t *testing.T) {
	type testCase struct {
		req        *http.Request
		resp       *httptest.ResponseRecorder
		code       int
		interactor *admin_interactor.TestImplementation
	}

	cases := make([]testCase, 0)

	body := url.Values{}
	body.Set("email", "email")
	body.Set("password", "password")
	body.Set("password_repeat", "password_repeat")
	req := httptest.NewRequest("POST", "https://localhost/admin/restore_password/", strings.NewReader(body.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp := httptest.NewRecorder()
	interactor := admin_interactor.NewTestImplementation()
	interactor.On("RestorePassword", context.WithValue(req.Context(), "db", nil), resp, "email", "password", "password_repeat").Return(0, errors.New("error"))

	cases = append(cases, testCase{
		req:        req,
		resp:       resp,
		code:       200,
		interactor: interactor,
	})

	body = url.Values{}
	body.Set("email", "email")
	body.Set("password", "password")
	body.Set("password_repeat", "password_repeat")
	req = httptest.NewRequest("POST", "https://localhost/admin/restore_password/", strings.NewReader(body.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp = httptest.NewRecorder()
	interactor = admin_interactor.NewTestImplementation()
	interactor.On("RestorePassword", context.WithValue(req.Context(), "db", nil), resp, "email", "password", "password_repeat").Return(0, nil)

	cases = append(cases, testCase{
		req:        req,
		resp:       resp,
		code:       http.StatusSeeOther,
		interactor: interactor,
	})

	for i, el := range cases {
		controller := NewAdminController(el.interactor, nil)
		controller.RestorePassword(el.resp, el.req)

		if el.resp.Code != el.code {
			t.Error("codes don't match: ", i)
		}

		if el.interactor != nil {
			el.interactor.AssertExpectations(t)
		}
	}
}
