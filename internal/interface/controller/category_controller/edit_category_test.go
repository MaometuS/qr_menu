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

type editCategoryTestCase struct {
	req        *http.Request
	resp       *httptest.ResponseRecorder
	code       int
	interactor *category_interactor.TestImplementation
}

func buildEditCategoryTestCase(category models.Category, code int, hasInteractor bool, err error) (*editCategoryTestCase, error) {
	body := &bytes.Buffer{}
	multipartWriter := multipart.NewWriter(body)

	if category.ID != 0 {
		err := multipartWriter.WriteField("id", strconv.FormatInt(category.ID, 10))
		if err != nil {
			return nil, err
		}
	}

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
		interactor.On("EditCategory", context.WithValue(req.Context(), "db", nil), int64(1), &category).Return(err)
	}

	return &editCategoryTestCase{
		req:        req,
		resp:       resp,
		code:       code,
		interactor: interactor,
	}, nil
}

func TestCategoryController_EditCategory(t *testing.T) {
	cases := make([]editCategoryTestCase, 0)

	cases = append(cases, editCategoryTestCase{
		req:        httptest.NewRequest("POST", "/edit_category/", nil),
		resp:       httptest.NewRecorder(),
		code:       500,
		interactor: nil,
	})

	testCase, err := buildEditCategoryTestCase(models.Category{
		ID:         1,
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

	testCase, err = buildEditCategoryTestCase(models.Category{
		ID:         1,
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

	testCase, err = buildEditCategoryTestCase(models.Category{
		ID:         1,
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

	testCase, err = buildEditCategoryTestCase(models.Category{
		ID:         1,
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

	testCase, err = buildEditCategoryTestCase(models.Category{
		ID:         1,
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

	testCase, err = buildEditCategoryTestCase(models.Category{
		Name:       "name",
		Background: "background",
		MenuID:     10,
		IsVisible:  true,
		OrderParam: 20,
	}, http.StatusExpectationFailed, false, nil)
	if err != nil {
		t.Error(err)
	}
	cases = append(cases, *testCase)

	for i, el := range cases {
		controller := NewCategoryController(el.interactor, nil)
		controller.EditCategory(el.resp, el.req)

		if el.resp.Code != el.code {
			t.Error("codes don't match: ", i)
		}

		if el.interactor != nil {
			el.interactor.AssertExpectations(t)
		}
	}
}
