package establishment_controller

import (
	"context"
	"errors"
	"gitlab.com/maometusu/qr_menu/internal/use_case/interactor/establishment_interactor"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestController_GetQrCode(t *testing.T) {
	type testCase struct {
		req        *http.Request
		resp       *httptest.ResponseRecorder
		code       int
		interactor *establishment_interactor.TestImplementation
	}

	cases := make([]testCase, 0)

	req := httptest.NewRequest("GET", "/establishment/", nil)
	resp := httptest.NewRecorder()

	cases = append(cases, testCase{
		req:  req,
		resp: resp,
		code: 417,
	})

	req = httptest.NewRequest("GET", "/establishment?id=2", nil)
	resp = httptest.NewRecorder()
	interactor := establishment_interactor.NewTestImplementation()
	interactor.On("GetQrCode", context.WithValue(req.Context(), "db", nil), resp, "https://example.com", int64(2)).Return(nil)

	cases = append(cases, testCase{
		req:        req,
		resp:       resp,
		code:       http.StatusOK,
		interactor: interactor,
	})

	req = httptest.NewRequest("GET", "/establishment?id=2", nil)
	resp = httptest.NewRecorder()
	interactor = establishment_interactor.NewTestImplementation()
	interactor.On("GetQrCode", context.WithValue(req.Context(), "db", nil), resp, "https://example.com", int64(2)).Return(errors.New(""))

	cases = append(cases, testCase{
		req:        req,
		resp:       resp,
		code:       http.StatusInternalServerError,
		interactor: interactor,
	})

	for i, el := range cases {
		controller := NewEstablishmentController(el.interactor, nil)
		controller.GetQrCode(el.resp, el.req)

		if el.resp.Code != el.code {
			t.Error("codes don't match: ", i)
		}

		if el.interactor != nil {
			el.interactor.AssertExpectations(t)
		}
	}
}
