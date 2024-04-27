package establishment_controller

import (
	"context"
	"errors"
	"gitlab.com/maometusu/qr_menu/internal/use_case/interactor/establishment_interactor"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestController_EditEstablishmentPage(t *testing.T) {
	type testCase struct {
		req        *http.Request
		resp       *httptest.ResponseRecorder
		code       int
		interactor *establishment_interactor.TestImplementation
	}

	cases := make([]testCase, 0)

	cases = append(cases, testCase{
		req:        httptest.NewRequest("GET", "/edit_establishment_page/", nil),
		resp:       httptest.NewRecorder(),
		code:       500,
		interactor: nil,
	})

	req := httptest.NewRequest("GET", "/edit_establishment_page/", nil)
	req = req.WithContext(context.WithValue(req.Context(), "id", int64(1)))
	resp := httptest.NewRecorder()
	interactor := establishment_interactor.NewTestImplementation()
	interactor.On("EditEstablishmentPage", context.WithValue(req.Context(), "db", nil), resp, int64(1), "").Return(nil)

	cases = append(cases, testCase{
		req:        req,
		resp:       resp,
		code:       http.StatusOK,
		interactor: interactor,
	})

	req = httptest.NewRequest("POST", "/edit_establishment_page/", nil)
	req = req.WithContext(context.WithValue(req.Context(), "id", int64(1)))
	resp = httptest.NewRecorder()
	interactor = establishment_interactor.NewTestImplementation()
	interactor.On("EditEstablishmentPage", context.WithValue(req.Context(), "db", nil), resp, int64(1), "").Return(errors.New(""))

	cases = append(cases, testCase{
		req:        req,
		resp:       resp,
		code:       http.StatusInternalServerError,
		interactor: interactor,
	})

	for i, el := range cases {
		controller := NewEstablishmentController(el.interactor, nil)
		controller.EditEstablishmentPage(el.resp, el.req)

		if el.resp.Code != el.code {
			t.Error("codes don't match: ", i)
		}

		if el.interactor != nil {
			el.interactor.AssertExpectations(t)
		}
	}
}
