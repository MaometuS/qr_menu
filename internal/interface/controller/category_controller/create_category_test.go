package category_controller

import (
	"bytes"
	"context"
	"errors"
	"gitlab.com/maometusu/qr_menu/internal/entity/models"
	"gitlab.com/maometusu/qr_menu/internal/use_case/interactor/category_interactor"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

type createCategoryTestCase struct {
	req        *http.Request
	resp       *httptest.ResponseRecorder
	code       int
	interactor *category_interactor.TestImplementation
}

func buildCreateCategoryTestCase(category models.Category, code int, hasInteractor bool, err error) (*createCategoryTestCase, error) {
	body := &bytes.Buffer{}
	multipartWriter := multipart.NewWriter(body)

	if category.Name != "" {
		err := multipartWriter.WriteField("name", "name")
		if err != nil {
			return nil, err
		}
	}

	if category.MenuID != 0 {
		err := multipartWriter.WriteField("menu_id", strconv.FormatInt(category.MenuID, 10))
		if err != nil {
			return nil, err
		}
	}

	if category.IsVisible {
		err := multipartWriter.WriteField("is_visible", "on")
		if err != nil {
			return nil, err
		}
	}

	if category.OrderParam != 0 {
		err := multipartWriter.WriteField("order_param", strconv.FormatInt(category.OrderParam, 10))
		if err != nil {
			return nil, err
		}
	}

	if category.Background != "" {
		fileWriter, err := multipartWriter.CreateFormFile("background", "background.jpg")
		if err != nil {
			return nil, err
		}

		_, err = fileWriter.Write([]byte(category.Background))
		if err != nil {
			return nil, err
		}

		err = multipartWriter.Close()
		if err != nil {
			return nil, err
		}
	}

	req := httptest.NewRequest("POST", "/create_category/", body)
	req = req.WithContext(context.WithValue(req.Context(), "id", int64(1)))
	req.Header.Set("Content-Type", multipartWriter.FormDataContentType())
	resp := httptest.NewRecorder()

	var interactor *category_interactor.TestImplementation
	if hasInteractor {
		interactor = category_interactor.NewTestImplementation()
		category.Background = ""
		interactor.On("CreateCategory", context.WithValue(req.Context(), "db", nil), int64(1), &category).Return(err)
	}

	return &createCategoryTestCase{
		req:        req,
		resp:       resp,
		code:       code,
		interactor: interactor,
	}, nil
}

func TestCategoryController_CreateCategory(t *testing.T) {
	cases := make([]createCategoryTestCase, 0)

	cases = append(cases, createCategoryTestCase{
		req:        httptest.NewRequest("POST", "/create_category/", nil),
		resp:       httptest.NewRecorder(),
		code:       500,
		interactor: nil,
	})

	testCase, err := buildCreateCategoryTestCase(models.Category{
		Name:       "name",
		Background: "background",
		MenuID:     10,
		IsVisible:  true,
		OrderParam: 20,
	}, http.StatusSeeOther, true, nil)
	if err != nil {
		t.Error(err)
	}
	cases = append(cases, *testCase)

	testCase, err = buildCreateCategoryTestCase(models.Category{
		Name:       "name",
		Background: "background",
		MenuID:     10,
		IsVisible:  true,
		OrderParam: 20,
	}, http.StatusInternalServerError, true, errors.New(""))
	if err != nil {
		t.Error(err)
	}
	cases = append(cases, *testCase)

	testCase, err = buildCreateCategoryTestCase(models.Category{
		Name:       "name",
		Background: "",
		MenuID:     10,
		IsVisible:  true,
		OrderParam: 20,
	}, http.StatusExpectationFailed, false, nil)
	if err != nil {
		t.Error(err)
	}
	cases = append(cases, *testCase)

	testCase, err = buildCreateCategoryTestCase(models.Category{
		Name:       "name",
		Background: "background",
		MenuID:     0,
		IsVisible:  true,
		OrderParam: 20,
	}, http.StatusExpectationFailed, false, nil)
	if err != nil {
		t.Error(err)
	}
	cases = append(cases, *testCase)

	testCase, err = buildCreateCategoryTestCase(models.Category{
		Name:       "name",
		Background: "background",
		MenuID:     10,
		IsVisible:  true,
		OrderParam: 0,
	}, http.StatusExpectationFailed, false, nil)
	if err != nil {
		t.Error(err)
	}
	cases = append(cases, *testCase)

	for i, el := range cases {
		controller := NewCategoryController(el.interactor, nil)
		controller.CreateCategory(el.resp, el.req)

		if el.resp.Code != el.code {
			t.Error("codes don't match: ", i)
		}

		if el.interactor != nil {
			el.interactor.AssertExpectations(t)
		}
	}
}
