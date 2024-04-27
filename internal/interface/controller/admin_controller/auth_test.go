package admin_controller

import (
	"context"
	"errors"
	"gitlab.com/maometusu/qr_menu/internal/use_case/interactor/admin_interactor"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestController_Auth(t *testing.T) {
	type testCase struct {
		req        *http.Request
		resp       *httptest.ResponseRecorder
		code       int
		interactor *admin_interactor.TestImplementation
	}

	cases := make([]testCase, 0)

	cases = append(cases, testCase{
		req:        httptest.NewRequest("GET", "/profile_page/", nil),
		resp:       httptest.NewRecorder(),
		code:       http.StatusUnauthorized,
		interactor: nil,
	})

	cases = append(cases, testCase{
		req:        httptest.NewRequest("GET", "/admin/", nil),
		resp:       httptest.NewRecorder(),
		code:       http.StatusPermanentRedirect,
		interactor: nil,
	})

	cases = append(cases, testCase{
		req:        httptest.NewRequest("GET", "/admin/e/t", nil),
		resp:       httptest.NewRecorder(),
		code:       http.StatusPermanentRedirect,
		interactor: nil,
	})

	req := httptest.NewRequest("GET", "https://localhost/admin/login/", nil)
	req.AddCookie(&http.Cookie{
		Name:     "auth_token",
		Value:    "cookie",
		Path:     "/",
		Expires:  time.Now().AddDate(1, 0, 0),
		Secure:   false,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	})
	resp := httptest.NewRecorder()
	interactor := admin_interactor.NewTestImplementation()
	interactor.On("Auth", context.WithValue(req.Context(), "db", nil), "cookie").Return(1, nil)

	cases = append(cases, testCase{
		req:        req,
		resp:       resp,
		code:       200,
		interactor: interactor,
	})

	req = httptest.NewRequest("GET", "https://localhost/admin/login/", nil)
	req.AddCookie(&http.Cookie{
		Name:     "auth_token",
		Value:    "cookie",
		Path:     "/",
		Expires:  time.Now().AddDate(1, 0, 0),
		Secure:   false,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	})
	resp = httptest.NewRecorder()
	interactor = admin_interactor.NewTestImplementation()
	interactor.On("Auth", context.WithValue(req.Context(), "db", nil), "cookie").Return(1, errors.New(""))

	cases = append(cases, testCase{
		req:        req,
		resp:       resp,
		code:       http.StatusUnauthorized,
		interactor: interactor,
	})

	for i, el := range cases {
		controller := NewAdminController(el.interactor, nil)
		controller.Auth(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {})).ServeHTTP(el.resp, el.req)

		if el.resp.Code != el.code {
			t.Error("codes don't match: ", i)
		}

		if el.interactor != nil {
			el.interactor.AssertExpectations(t)
		}
	}
}
