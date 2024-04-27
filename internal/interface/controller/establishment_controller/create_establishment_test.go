package establishment_controller

import (
	"context"
	"errors"
	"gitlab.com/maometusu/qr_menu/internal/entity/models"
	"gitlab.com/maometusu/qr_menu/internal/use_case/interactor/establishment_interactor"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"strings"
	"testing"
)

type createEstablishmentTestCase struct {
	req        *http.Request
	resp       *httptest.ResponseRecorder
	code       int
	interactor *establishment_interactor.TestImplementation
}

func buildCreateEstablishmentTestCase(item models.Establishment, code int, hasInteractor bool, err error) (*createEstablishmentTestCase, error) {
	vals := url.Values{}

	if item.Name != "" {
		vals.Set("name", item.Name)
	}

	if item.Link != "" {
		vals.Set("link", item.Link)
	}

	if item.CurrencyID != 0 {
		vals.Set("currency", strconv.FormatInt(item.CurrencyID, 10))
	}

	req := httptest.NewRequest("POST", "/create_establishment/", strings.NewReader(vals.Encode()))
	req = req.WithContext(context.WithValue(req.Context(), "id", int64(1)))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp := httptest.NewRecorder()

	var interactor *establishment_interactor.TestImplementation
	if hasInteractor {
		interactor = establishment_interactor.NewTestImplementation()
		item.ProfileID = 1
		interactor.On("CreateEstablishment", context.WithValue(req.Context(), "db", nil), &item).Return(err)
	}

	return &createEstablishmentTestCase{
		req:        req,
		resp:       resp,
		code:       code,
		interactor: interactor,
	}, nil
}

func TestEstablishmentController_CreateEstablishment(t *testing.T) {
	cases := make([]createEstablishmentTestCase, 0)

	cases = append(cases, createEstablishmentTestCase{
		req:        httptest.NewRequest("POST", "/create_establishment/", nil),
		resp:       httptest.NewRecorder(),
		code:       500,
		interactor: nil,
	})

	testCase, err := buildCreateEstablishmentTestCase(models.Establishment{
		Name:       "name",
		Link:       "link",
		CurrencyID: 10,
	}, http.StatusSeeOther, true, nil)
	if err != nil {
		t.Error(err)
	}
	cases = append(cases, *testCase)

	testCase, err = buildCreateEstablishmentTestCase(models.Establishment{
		Name:       "name",
		Link:       "link",
		CurrencyID: 10,
	}, http.StatusInternalServerError, true, errors.New(""))
	if err != nil {
		t.Error(err)
	}
	cases = append(cases, *testCase)

	testCase, err = buildCreateEstablishmentTestCase(models.Establishment{
		Name:       "name",
		Link:       "link",
		CurrencyID: 0,
	}, http.StatusExpectationFailed, false, nil)
	if err != nil {
		t.Error(err)
	}
	cases = append(cases, *testCase)

	for i, el := range cases {
		controller := NewEstablishmentController(el.interactor, nil)
		controller.CreateEstablishment(el.resp, el.req)

		if el.resp.Code != el.code {
			t.Error("codes don't match:", i, "expected code", el.code, "got code", el.resp.Code)
		}

		if el.interactor != nil {
			el.interactor.AssertExpectations(t)
		}
	}
}
