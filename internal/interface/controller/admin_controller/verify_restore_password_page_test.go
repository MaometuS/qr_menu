package admin_controller

import (
	"context"
	"errors"
	"gitlab.com/maometusu/qr_menu/internal/use_case/interactor/admin_interactor"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestController_VerifyRestorePasswordPage(t *testing.T) {
	type testCase struct {
		req        *http.Request
		resp       *httptest.ResponseRecorder
		code       int
		interactor *admin_interactor.TestImplementation
	}

	cases := make([]testCase, 0)

	req := httptest.NewRequest("GET", "https://localhost/admin/verify_restore_password_page?id=1", nil)
	resp := httptest.NewRecorder()
	interactor := admin_interactor.NewTestImplementation()
	interactor.On("VerifyRestorePasswordPage", context.WithValue(req.Context(), "db", nil), resp, int64(1)).Return(errors.New("error"))

	cases = append(cases, testCase{
		req:        req,
		resp:       resp,
		code:       500,
		interactor: interactor,
	})

	req = httptest.NewRequest("GET", "https://localhost/admin/verify_restore_password_page", nil)
	resp = httptest.NewRecorder()

	cases = append(cases, testCase{
		req:        req,
		resp:       resp,
		code:       417,
		interactor: nil,
	})

	for i, el := range cases {
		controller := NewAdminController(el.interactor, nil)
		controller.VerifyRestorePasswordPage(el.resp, el.req)

		if el.resp.Code != el.code {
			t.Error("codes don't match: ", i)
		}

		if el.interactor != nil {
			el.interactor.AssertExpectations(t)
		}
	}
}
