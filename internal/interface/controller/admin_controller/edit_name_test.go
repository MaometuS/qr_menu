package admin_controller

import (
	"context"
	"errors"
	"gitlab.com/maometusu/qr_menu/internal/use_case/interactor/admin_interactor"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestController_EditName(t *testing.T) {
	type testCase struct {
		req        *http.Request
		resp       *httptest.ResponseRecorder
		code       int
		result     string
		interactor *admin_interactor.TestImplementation
	}

	cases := make([]testCase, 0)

	cases = append(cases, testCase{
		req:        httptest.NewRequest("POST", "https://localhost/edit_name/", nil),
		resp:       httptest.NewRecorder(),
		code:       500,
		result:     "",
		interactor: nil,
	})

	body := url.Values{}
	body.Set("email", "email")
	body.Set("name", "name")
	req := httptest.NewRequest("POST", "https://localhost/admin/edit_name/", strings.NewReader(body.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req = req.WithContext(context.WithValue(req.Context(), "id", int64(1)))
	resp := httptest.NewRecorder()
	interactor := admin_interactor.NewTestImplementation()
	interactor.On("EditName", context.WithValue(req.Context(), "db", nil), "email", "name", int64(1)).Return(true, nil)

	cases = append(cases, testCase{
		req:        req,
		resp:       resp,
		code:       200,
		result:     "true",
		interactor: interactor,
	})

	body = url.Values{}
	body.Set("email", "email")
	body.Set("name", "name")
	req = httptest.NewRequest("POST", "https://localhost/admin/edit_name/", strings.NewReader(body.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req = req.WithContext(context.WithValue(req.Context(), "id", int64(1)))
	resp = httptest.NewRecorder()
	interactor = admin_interactor.NewTestImplementation()
	interactor.On("EditName", context.WithValue(req.Context(), "db", nil), "email", "name", int64(1)).Return(false, nil)

	cases = append(cases, testCase{
		req:        req,
		resp:       resp,
		code:       200,
		result:     "false",
		interactor: interactor,
	})

	body = url.Values{}
	body.Set("email", "email")
	body.Set("name", "name")
	req = httptest.NewRequest("POST", "https://localhost/admin/edit_name/", strings.NewReader(body.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req = req.WithContext(context.WithValue(req.Context(), "id", int64(1)))
	resp = httptest.NewRecorder()
	interactor = admin_interactor.NewTestImplementation()
	interactor.On("EditName", context.WithValue(req.Context(), "db", nil), "email", "name", int64(1)).Return(false, errors.New(""))

	cases = append(cases, testCase{
		req:        req,
		resp:       resp,
		code:       500,
		result:     "",
		interactor: interactor,
	})

	for i, el := range cases {
		controller := NewAdminController(el.interactor, nil)
		controller.EditName(el.resp, el.req)

		if el.resp.Code != el.code {
			t.Error("codes don't match: ", i)
		}

		if el.code == 200 && el.result != el.resp.Body.String() {
			t.Error("results don't match: ", i)
		}

		if el.interactor != nil {
			el.interactor.AssertExpectations(t)
		}
	}
}
