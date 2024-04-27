package category_controller

import (
	"context"
	"errors"
	"gitlab.com/maometusu/qr_menu/internal/use_case/interactor/category_interactor"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCategoryController_GetCategories(t *testing.T) {
	type testCase struct {
		req        *http.Request
		resp       *httptest.ResponseRecorder
		code       int
		interactor *category_interactor.TestImplementation
	}

	cases := make([]testCase, 0)

	cases = append(cases, testCase{
		req:        httptest.NewRequest("GET", "/categories/", nil),
		resp:       httptest.NewRecorder(),
		code:       500,
		interactor: nil,
	})

	req := httptest.NewRequest("GET", "/categories/", nil)
	req = req.WithContext(context.WithValue(req.Context(), "id", int64(1)))
	resp := httptest.NewRecorder()

	cases = append(cases, testCase{
		req:  req,
		resp: resp,
		code: 417,
	})

	req = httptest.NewRequest("GET", "/categories?menu_id=2", nil)
	req = req.WithContext(context.WithValue(req.Context(), "id", int64(1)))
	resp = httptest.NewRecorder()
	interactor := category_interactor.NewTestImplementation()
	interactor.On("GetCategories", context.WithValue(req.Context(), "db", nil), resp, int64(1), int64(2)).Return(nil)

	cases = append(cases, testCase{
		req:        req,
		resp:       resp,
		code:       200,
		interactor: interactor,
	})

	req = httptest.NewRequest("GET", "/categories?menu_id=2", nil)
	req = req.WithContext(context.WithValue(req.Context(), "id", int64(1)))
	resp = httptest.NewRecorder()
	interactor = category_interactor.NewTestImplementation()
	interactor.On("GetCategories", context.WithValue(req.Context(), "db", nil), resp, int64(1), int64(2)).Return(errors.New(""))

	cases = append(cases, testCase{
		req:        req,
		resp:       resp,
		code:       http.StatusInternalServerError,
		interactor: interactor,
	})

	for i, el := range cases {
		controller := NewCategoryController(el.interactor, nil)
		controller.GetCategories(el.resp, el.req)

		if el.resp.Code != el.code {
			t.Error("codes don't match: ", i)
		}

		if el.interactor != nil {
			el.interactor.AssertExpectations(t)
		}
	}
}
