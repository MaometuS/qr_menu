package public_controller

import (
	"context"
	"errors"
	"gitlab.com/maometusu/qr_menu/internal/use_case/interactor/public_interactor"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPublicController_AddOrder(t *testing.T) {
	type testCase struct {
		req        *http.Request
		resp       *httptest.ResponseRecorder
		code       int
		interactor *public_interactor.TestImplementation
	}

	cases := make([]testCase, 0)

	req := httptest.NewRequest("GET", "/add_order/", nil)
	resp := httptest.NewRecorder()

	cases = append(cases, testCase{
		req:  req,
		resp: resp,
		code: 417,
	})

	req = httptest.NewRequest("GET", "/add_order?id=2", nil)
	resp = httptest.NewRecorder()

	cases = append(cases, testCase{
		req:  req,
		resp: resp,
		code: 417,
	})

	req = httptest.NewRequest("GET", "/add_order?id=2&menu_id=3", nil)
	resp = httptest.NewRecorder()
	interactor := public_interactor.NewTestImplementation()
	interactor.On("AddItem", context.WithValue(req.Context(), "db", nil), int64(2), make(map[int64][]int64)).Return(make(map[int64][]int64), nil)

	cases = append(cases, testCase{
		req:        req,
		resp:       resp,
		code:       http.StatusSeeOther,
		interactor: interactor,
	})

	req = httptest.NewRequest("GET", "/add_order?id=2&menu_id=3", nil)
	resp = httptest.NewRecorder()
	interactor = public_interactor.NewTestImplementation()
	interactor.On("AddItem", context.WithValue(req.Context(), "db", nil), int64(2), make(map[int64][]int64)).Return(make(map[int64][]int64), errors.New(""))

	cases = append(cases, testCase{
		req:        req,
		resp:       resp,
		code:       500,
		interactor: interactor,
	})

	for i, el := range cases {
		controller := NewPublicController(el.interactor, nil)
		controller.AddOrder(el.resp, el.req)

		if el.resp.Code != el.code {
			t.Error("codes don't match: ", i)
		}

		if el.interactor != nil {
			el.interactor.AssertExpectations(t)
		}
	}
}
