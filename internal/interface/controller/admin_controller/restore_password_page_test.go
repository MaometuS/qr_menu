package admin_controller

import (
	"errors"
	"gitlab.com/maometusu/qr_menu/internal/use_case/interactor/admin_interactor"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestController_RestorePasswordPage(t *testing.T) {
	type testCase struct {
		req        *http.Request
		resp       *httptest.ResponseRecorder
		code       int
		interactor *admin_interactor.TestImplementation
	}

	cases := make([]testCase, 0)

	req := httptest.NewRequest("GET", "/restore_page/", nil)
	resp := httptest.NewRecorder()
	interactor := admin_interactor.NewTestImplementation()
	interactor.On("RestorePasswordPage", req.Context(), resp).Return(errors.New("error"))

	cases = append(cases, testCase{
		req:        req,
		resp:       resp,
		code:       500,
		interactor: interactor,
	})

	req = httptest.NewRequest("GET", "/restore_page/", nil)
	resp = httptest.NewRecorder()
	interactor = admin_interactor.NewTestImplementation()
	interactor.On("RestorePasswordPage", req.Context(), resp).Return(nil)

	cases = append(cases, testCase{
		req:        req,
		resp:       resp,
		code:       200,
		interactor: interactor,
	})

	for i, el := range cases {
		controller := NewAdminController(el.interactor, nil)
		controller.RestorePasswordPage(el.resp, el.req)

		if el.resp.Code != el.code {
			t.Error("codes don't match: ", i)
		}

		if el.interactor != nil {
			el.interactor.AssertExpectations(t)
		}
	}
}
