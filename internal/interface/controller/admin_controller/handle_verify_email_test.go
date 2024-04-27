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

func TestController_HandleVerifyEmail(t *testing.T) {
	type testCase struct {
		req        *http.Request
		resp       *httptest.ResponseRecorder
		code       int
		interactor *admin_interactor.TestImplementation
	}

	cases := make([]testCase, 0)

	body := url.Values{}
	body.Set("code", "code")
	req := httptest.NewRequest("POST", "https://localhost/admin/verify_email/", strings.NewReader(body.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp := httptest.NewRecorder()

	cases = append(cases, testCase{
		req:        req,
		resp:       resp,
		code:       http.StatusExpectationFailed,
		interactor: nil,
	})

	body = url.Values{}
	body.Set("id", "1")
	body.Set("code", "code")
	req = httptest.NewRequest("POST", "https://localhost/admin/verify_email/", strings.NewReader(body.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp = httptest.NewRecorder()
	interactor := admin_interactor.NewTestImplementation()
	interactor.On("HandleVerifyEmail", context.WithValue(req.Context(), "db", nil), resp, int64(1), "code").Return("", nil)

	cases = append(cases, testCase{
		req:        req,
		resp:       resp,
		code:       http.StatusSeeOther,
		interactor: interactor,
	})

	body = url.Values{}
	body.Set("id", "1")
	body.Set("code", "code")
	req = httptest.NewRequest("POST", "https://localhost/admin/verify_email/", strings.NewReader(body.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp = httptest.NewRecorder()
	interactor = admin_interactor.NewTestImplementation()
	interactor.On("HandleVerifyEmail", context.WithValue(req.Context(), "db", nil), resp, int64(1), "code").Return("", errors.New(""))

	cases = append(cases, testCase{
		req:        req,
		resp:       resp,
		code:       200,
		interactor: interactor,
	})

	for i, el := range cases {
		controller := NewAdminController(el.interactor, nil)
		controller.HandleVerifyEmail(el.resp, el.req)

		if el.resp.Code != el.code {
			t.Error("codes don't match: ", i)
		}

		if el.interactor != nil {
			el.interactor.AssertExpectations(t)
		}
	}
}
