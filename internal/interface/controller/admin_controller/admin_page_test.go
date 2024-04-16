package admin_controller

import (
	"context"
	"errors"
	"gitlab.com/maometusu/qr_menu/internal/use_case/interactor/admin_interactor"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestController_AdminPage(t *testing.T) {
	type testCase struct {
		req        *http.Request
		resp       *httptest.ResponseRecorder
		code       int
		interactor *admin_interactor.TestImplementation
	}

	cases := []testCase{
		{
			httptest.NewRequest("GET", "https://localhost/admin_page", nil),
			httptest.NewRecorder(),
			500,
			nil,
		},
		{
			httptest.NewRequest("GET", "https://localhost/admin_page", nil),
			httptest.NewRecorder(),
			200,
			admin_interactor.NewTestImplementation(),
		},
		{
			httptest.NewRequest("GET", "https://localhost/admin_page", nil),
			httptest.NewRecorder(),
			500,
			admin_interactor.NewTestImplementation(),
		},
	}

	cases[1].req = cases[1].req.WithContext(context.WithValue(cases[1].req.Context(), "id", int64(1)))
	cases[1].interactor.On(
		"AdminPage",
		context.WithValue(cases[1].req.Context(), "db", nil),
		cases[1].resp,
		int64(1),
	).Return(nil)

	cases[2].req = cases[2].req.WithContext(context.WithValue(cases[2].req.Context(), "id", int64(2)))
	cases[2].interactor.On(
		"AdminPage",
		context.WithValue(cases[2].req.Context(), "db", nil),
		cases[2].resp,
		int64(2),
	).Return(errors.New(""))

	for _, el := range cases {
		controller := NewAdminController(el.interactor, nil)
		controller.AdminPage(el.resp, el.req)

		if el.resp.Code != el.code {
			t.Error("codes don't match")
		}

		if el.interactor != nil {
			el.interactor.AssertExpectations(t)
		}
	}
}
