package menu_controller

import (
	"context"
	"errors"
	"gitlab.com/maometusu/qr_menu/internal/entity/models"
	"gitlab.com/maometusu/qr_menu/internal/use_case/interactor/menu_interactor"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"strings"
	"testing"
)

type editMenuTestCase struct {
	req        *http.Request
	resp       *httptest.ResponseRecorder
	code       int
	interactor *menu_interactor.TestImplementation
}

func buildEditMenuTestCase(item models.Menu, code int, hasInteractor bool, err error) (*editMenuTestCase, error) {
	vals := url.Values{}

	if item.ID != 0 {
		vals.Set("id", strconv.FormatInt(item.ID, 10))
	}

	if item.Name != "" {
		vals.Set("name", item.Name)
	}

	if item.IsVisible {
		vals.Set("is_visible", "on")
	}

	if item.EstablishmentID != 0 {
		vals.Set("establishment_id", strconv.FormatInt(item.EstablishmentID, 10))
	}

	if item.OrderParam != 0 {
		vals.Set("order", strconv.FormatInt(item.OrderParam, 10))
	}

	req := httptest.NewRequest("POST", "/create_item/", strings.NewReader(vals.Encode()))
	req = req.WithContext(context.WithValue(req.Context(), "id", int64(1)))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp := httptest.NewRecorder()

	var interactor *menu_interactor.TestImplementation
	if hasInteractor {
		interactor = menu_interactor.NewTestImplementation()
		interactor.On("EditMenu", context.WithValue(req.Context(), "db", nil), int64(1), &item).Return(err)
	}

	return &editMenuTestCase{
		req:        req,
		resp:       resp,
		code:       code,
		interactor: interactor,
	}, nil
}

func TestMenuController_EditMenu(t *testing.T) {
	cases := make([]editMenuTestCase, 0)

	cases = append(cases, editMenuTestCase{
		req:        httptest.NewRequest("POST", "/create_menu/", nil),
		resp:       httptest.NewRecorder(),
		code:       500,
		interactor: nil,
	})

	testCase, err := buildEditMenuTestCase(models.Menu{
		ID:              1,
		Name:            "name",
		IsVisible:       true,
		EstablishmentID: 10,
		OrderParam:      20,
	}, http.StatusSeeOther, true, nil)
	if err != nil {
		t.Error(err)
	}
	cases = append(cases, *testCase)

	testCase, err = buildEditMenuTestCase(models.Menu{
		ID:              1,
		Name:            "name",
		IsVisible:       true,
		EstablishmentID: 10,
		OrderParam:      20,
	}, http.StatusInternalServerError, true, errors.New(""))
	if err != nil {
		t.Error(err)
	}
	cases = append(cases, *testCase)

	testCase, err = buildEditMenuTestCase(models.Menu{
		ID:              1,
		Name:            "name",
		IsVisible:       true,
		EstablishmentID: 0,
		OrderParam:      20,
	}, http.StatusExpectationFailed, false, nil)
	if err != nil {
		t.Error(err)
	}
	cases = append(cases, *testCase)

	testCase, err = buildEditMenuTestCase(models.Menu{
		ID:              1,
		Name:            "name",
		IsVisible:       true,
		EstablishmentID: 10,
		OrderParam:      0,
	}, http.StatusExpectationFailed, false, nil)
	if err != nil {
		t.Error(err)
	}
	cases = append(cases, *testCase)

	testCase, err = buildEditMenuTestCase(models.Menu{
		ID:              0,
		Name:            "name",
		IsVisible:       true,
		EstablishmentID: 10,
		OrderParam:      20,
	}, http.StatusExpectationFailed, false, nil)
	if err != nil {
		t.Error(err)
	}
	cases = append(cases, *testCase)

	for i, el := range cases {
		controller := NewMenuController(el.interactor, nil)
		controller.EditMenu(el.resp, el.req)

		if el.resp.Code != el.code {
			t.Error("codes don't match:", i, "expected code", el.code, "got code", el.resp.Code)
		}

		if el.interactor != nil {
			el.interactor.AssertExpectations(t)
		}
	}
}
