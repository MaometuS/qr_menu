package establishment_controller

import (
	"context"
	"errors"
	"gitlab.com/maometusu/qr_menu/internal/use_case/interactor/establishment_interactor"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestController_DoesLinkExist(t *testing.T) {
	type testCase struct {
		req        *http.Request
		resp       *httptest.ResponseRecorder
		res        string
		code       int
		interactor *establishment_interactor.TestImplementation
	}

	cases := make([]testCase, 0)

	req := httptest.NewRequest("GET", "/does_link_exist?link=link", nil)
	resp := httptest.NewRecorder()
	interactor := establishment_interactor.NewTestImplementation()
	interactor.On("DoesLinkExist", context.WithValue(req.Context(), "db", nil), "link").Return(true, nil)

	cases = append(cases, testCase{
		req:        req,
		resp:       resp,
		res:        "true",
		code:       http.StatusOK,
		interactor: interactor,
	})

	req = httptest.NewRequest("GET", "/does_link_exist?link=link", nil)
	resp = httptest.NewRecorder()
	interactor = establishment_interactor.NewTestImplementation()
	interactor.On("DoesLinkExist", context.WithValue(req.Context(), "db", nil), "link").Return(false, nil)

	cases = append(cases, testCase{
		req:        req,
		resp:       resp,
		res:        "false",
		code:       http.StatusOK,
		interactor: interactor,
	})

	req = httptest.NewRequest("GET", "/does_link_exist?link=link", nil)
	resp = httptest.NewRecorder()
	interactor = establishment_interactor.NewTestImplementation()
	interactor.On("DoesLinkExist", context.WithValue(req.Context(), "db", nil), "link").Return(false, errors.New(""))

	cases = append(cases, testCase{
		req:        req,
		resp:       resp,
		res:        "",
		code:       http.StatusInternalServerError,
		interactor: interactor,
	})

	for i, el := range cases {
		controller := NewEstablishmentController(el.interactor, nil)
		controller.DoesLinkExist(el.resp, el.req)

		if el.resp.Code == 200 && el.res != el.resp.Body.String() {
			t.Error("results don't match: ", i)
		}

		if el.resp.Code != el.code {
			t.Error("codes don't match: ", i)
		}

		if el.interactor != nil {
			el.interactor.AssertExpectations(t)
		}
	}
}
