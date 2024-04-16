package admin_controller

import (
	"context"
	"errors"
	"gitlab.com/maometusu/qr_menu/internal/use_case/interactor/admin_interactor"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestController_ChangePassword(t *testing.T) {
	type testCase struct {
		req        *http.Request
		resp       *httptest.ResponseRecorder
		code       int
		interactor *admin_interactor.TestImplementation
	}

	cases := make([]testCase, 0)

	cases = append(cases, testCase{
		req:        httptest.NewRequest("POST", "https://localhost/change_password", nil),
		resp:       httptest.NewRecorder(),
		code:       500,
		interactor: nil,
	})

	req := httptest.NewRequest("POST", "https://localhost/change_password?password=123&repeat_password=123", nil)
	req = req.WithContext(context.WithValue(req.Context(), "id", int64(1)))
	interactor := admin_interactor.NewTestImplementation()
	interactor.On("ChangePassword", context.WithValue(req.Context(), "db", nil), "123", "123", int64(1)).Return(nil)

	cases = append(cases, testCase{
		req:        req,
		resp:       httptest.NewRecorder(),
		code:       200,
		interactor: interactor,
	})

	req = httptest.NewRequest("POST", "https://localhost/change_password?password=123&repeat_password=123", nil)
	req = req.WithContext(context.WithValue(req.Context(), "id", int64(1)))
	interactor = admin_interactor.NewTestImplementation()
	interactor.On("ChangePassword", context.WithValue(req.Context(), "db", nil), "123", "123", int64(1)).Return(errors.New(""))

	cases = append(cases, testCase{
		req:        req,
		resp:       httptest.NewRecorder(),
		code:       500,
		interactor: interactor,
	})

	for _, el := range cases {
		controller := NewAdminController(el.interactor, nil)
		controller.ChangePassword(el.resp, el.req)

		if el.resp.Code != el.code {
			t.Error("codes don't match")
		}

		if el.interactor != nil {
			el.interactor.AssertExpectations(t)
		}
	}
}
