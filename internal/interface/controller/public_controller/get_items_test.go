package public_controller

import (
	"context"
	"errors"
	"gitlab.com/maometusu/qr_menu/internal/use_case/interactor/public_interactor"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPublicController_GetItems(t *testing.T) {
	type testCase struct {
		req        *http.Request
		resp       *httptest.ResponseRecorder
		code       int
		interactor *public_interactor.TestImplementation
	}

	cases := make([]testCase, 0)

	req := httptest.NewRequest("GET", "/get_items/", nil)
	resp := httptest.NewRecorder()

	cases = append(cases, testCase{
		req:  req,
		resp: resp,
		code: 417,
	})

	req = httptest.NewRequest("GET", "/get_items?category_id=2", nil)
	resp = httptest.NewRecorder()
	interactor := public_interactor.NewTestImplementation()
	interactor.On("GetItems", context.WithValue(req.Context(), "db", nil), resp, int64(2), make(map[int64][]int64)).Return(nil)

	cases = append(cases, testCase{
		req:        req,
		resp:       resp,
		code:       http.StatusOK,
		interactor: interactor,
	})

	req = httptest.NewRequest("GET", "/get_items?category_id=2", nil)
	resp = httptest.NewRecorder()
	interactor = public_interactor.NewTestImplementation()
	interactor.On("GetItems", context.WithValue(req.Context(), "db", nil), resp, int64(2), make(map[int64][]int64)).Return(errors.New(""))

	cases = append(cases, testCase{
		req:        req,
		resp:       resp,
		code:       http.StatusInternalServerError,
		interactor: interactor,
	})

	for i, el := range cases {
		controller := NewPublicController(el.interactor, nil)
		controller.GetItems(el.resp, el.req)

		if el.resp.Code != el.code {
			t.Error("codes don't match: ", i)
		}

		if el.interactor != nil {
			el.interactor.AssertExpectations(t)
		}
	}
}
