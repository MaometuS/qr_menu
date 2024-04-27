package admin_controller

import (
	"context"
	"errors"
	"gitlab.com/maometusu/qr_menu/internal/use_case/interactor/admin_interactor"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestController_ProfilePage(t *testing.T) {
	type testCase struct {
		req        *http.Request
		resp       *httptest.ResponseRecorder
		code       int
		interactor *admin_interactor.TestImplementation
	}

	cases := make([]testCase, 0)

	cases = append(cases, testCase{
		req:        httptest.NewRequest("POST", "https://localhost/profile_page", nil),
		resp:       httptest.NewRecorder(),
		code:       500,
		interactor: nil,
	})

	req := httptest.NewRequest("GET", "https://localhost/admin/profile_page/", nil)
	req = req.WithContext(context.WithValue(req.Context(), "id", int64(1)))
	resp := httptest.NewRecorder()
	interactor := admin_interactor.NewTestImplementation()
	interactor.On("ProfilePage", context.WithValue(req.Context(), "db", nil), resp, int64(1)).Return(nil)

	cases = append(cases, testCase{
		req:        req,
		resp:       resp,
		code:       200,
		interactor: interactor,
	})

	req = httptest.NewRequest("GET", "https://localhost/admin/profile_page/", nil)
	req = req.WithContext(context.WithValue(req.Context(), "id", int64(1)))
	resp = httptest.NewRecorder()
	interactor = admin_interactor.NewTestImplementation()
	interactor.On("ProfilePage", context.WithValue(req.Context(), "db", nil), resp, int64(1)).Return(errors.New(""))

	cases = append(cases, testCase{
		req:        req,
		resp:       resp,
		code:       500,
		interactor: interactor,
	})

	for i, el := range cases {
		controller := NewAdminController(el.interactor, nil)
		controller.ProfilePage(el.resp, el.req)

		if el.resp.Code != el.code {
			t.Error("codes don't match: ", i)
		}

		if el.interactor != nil {
			el.interactor.AssertExpectations(t)
		}
	}
}
