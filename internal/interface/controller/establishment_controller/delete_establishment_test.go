package establishment_controller

import (
	"context"
	"errors"
	"gitlab.com/maometusu/qr_menu/internal/use_case/interactor/establishment_interactor"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestController_DeleteEstablishment(t *testing.T) {
	type testCase struct {
		req        *http.Request
		resp       *httptest.ResponseRecorder
		code       int
		interactor *establishment_interactor.TestImplementation
	}

	cases := make([]testCase, 0)

	cases = append(cases, testCase{
		req:        httptest.NewRequest("POST", "/delete_menu/", nil),
		resp:       httptest.NewRecorder(),
		code:       500,
		interactor: nil,
	})

	req := httptest.NewRequest("POST", "/delete_menu/", nil)
	req = req.WithContext(context.WithValue(req.Context(), "id", int64(1)))
	resp := httptest.NewRecorder()

	cases = append(cases, testCase{
		req:  req,
		resp: resp,
		code: 417,
	})

	req = httptest.NewRequest("POST", "/delete_menu?id=2", nil)
	req = req.WithContext(context.WithValue(req.Context(), "id", int64(1)))
	resp = httptest.NewRecorder()
	interactor := establishment_interactor.NewTestImplementation()
	interactor.On("DeleteEstablishment", context.WithValue(req.Context(), "db", nil), int64(1), int64(2)).Return(nil)

	cases = append(cases, testCase{
		req:        req,
		resp:       resp,
		code:       http.StatusSeeOther,
		interactor: interactor,
	})

	req = httptest.NewRequest("POST", "/delete_item?id=2", nil)
	req = req.WithContext(context.WithValue(req.Context(), "id", int64(1)))
	resp = httptest.NewRecorder()
	interactor = establishment_interactor.NewTestImplementation()
	interactor.On("DeleteEstablishment", context.WithValue(req.Context(), "db", nil), int64(1), int64(2)).Return(errors.New(""))

	cases = append(cases, testCase{
		req:        req,
		resp:       resp,
		code:       http.StatusInternalServerError,
		interactor: interactor,
	})

	for i, el := range cases {
		controller := NewEstablishmentController(el.interactor, nil)
		controller.DeleteEstablishment(el.resp, el.req)

		if el.resp.Code != el.code {
			t.Error("codes don't match: ", i)
		}

		if el.interactor != nil {
			el.interactor.AssertExpectations(t)
		}
	}
}
