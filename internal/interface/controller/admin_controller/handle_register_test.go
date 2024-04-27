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

func TestController_HandleRegister(t *testing.T) {
	type testCase struct {
		req        *http.Request
		resp       *httptest.ResponseRecorder
		code       int
		interactor *admin_interactor.TestImplementation
	}

	cases := make([]testCase, 0)

	body := url.Values{}
	body.Set("name", "name")
	body.Set("email", "email")
	body.Set("password", "password")
	body.Set("pass_repeat", "pass_repeat")
	req := httptest.NewRequest("POST", "https://localhost/admin/handle_register/", strings.NewReader(body.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp := httptest.NewRecorder()
	interactor := admin_interactor.NewTestImplementation()
	interactor.On("HandleRegister", context.WithValue(req.Context(), "db", nil), resp, "name", "email", "password", "pass_repeat").Return(1, nil)

	cases = append(cases, testCase{
		req:        req,
		resp:       resp,
		code:       http.StatusSeeOther,
		interactor: interactor,
	})

	body = url.Values{}
	body.Set("name", "name")
	body.Set("email", "email")
	body.Set("password", "password")
	body.Set("pass_repeat", "pass_repeat")
	req = httptest.NewRequest("POST", "https://localhost/admin/handle_register/", strings.NewReader(body.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp = httptest.NewRecorder()
	interactor = admin_interactor.NewTestImplementation()
	interactor.On("HandleRegister", context.WithValue(req.Context(), "db", nil), resp, "name", "email", "password", "pass_repeat").Return(1, errors.New(""))

	cases = append(cases, testCase{
		req:        req,
		resp:       resp,
		code:       200,
		interactor: interactor,
	})

	for i, el := range cases {
		controller := NewAdminController(el.interactor, nil)
		controller.HandleRegister(el.resp, el.req)

		if el.resp.Code != el.code {
			t.Error("codes don't match: ", i)
		}

		if el.interactor != nil {
			el.interactor.AssertExpectations(t)
		}
	}
}
