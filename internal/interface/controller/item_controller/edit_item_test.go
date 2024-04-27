package item_controller

import (
	"bytes"
	"context"
	"errors"
	"gitlab.com/maometusu/qr_menu/internal/entity/models"
	"gitlab.com/maometusu/qr_menu/internal/use_case/interactor/item_interactor"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

type editItemTestCase struct {
	req        *http.Request
	resp       *httptest.ResponseRecorder
	code       int
	interactor *item_interactor.TestImplementation
}

func buildEditItemTestCase(item models.Item, code int, hasInteractor bool, err error) (*editItemTestCase, error) {
	body := &bytes.Buffer{}
	multipartWriter := multipart.NewWriter(body)

	if item.ID != 0 {
		err := multipartWriter.WriteField("id", strconv.FormatInt(item.ID, 10))
		if err != nil {
			return nil, err
		}
	}

	if item.Name != "" {
		err := multipartWriter.WriteField("name", item.Name)
		if err != nil {
			return nil, err
		}
	}

	if item.Weight != "" {
		err := multipartWriter.WriteField("weight", item.Weight)
		if err != nil {
			return nil, err
		}
	}

	if item.Price != 0 {
		err := multipartWriter.WriteField("price", strconv.FormatInt(int64(item.Price), 10))
		if err != nil {
			return nil, err
		}
	}

	if item.AdditionalInfo != "" {
		err := multipartWriter.WriteField("additional_info", item.AdditionalInfo)
		if err != nil {
			return nil, err
		}
	}

	if item.IsVisible {
		err := multipartWriter.WriteField("is_visible", "on")
		if err != nil {
			return nil, err
		}
	}

	if item.IsAvailable {
		err := multipartWriter.WriteField("is_available", "on")
		if err != nil {
			return nil, err
		}
	}

	if item.CategoryID != 0 {
		err := multipartWriter.WriteField("category_id", strconv.FormatInt(item.CategoryID, 10))
		if err != nil {
			return nil, err
		}
	}

	if item.OrderParam != 0 {
		err := multipartWriter.WriteField("order_param", strconv.FormatInt(item.OrderParam, 10))
		if err != nil {
			return nil, err
		}
	}

	if item.Image != "" {
		fileWriter, err := multipartWriter.CreateFormFile("image", "image.jpg")
		if err != nil {
			return nil, err
		}

		_, err = fileWriter.Write([]byte(item.Image))
		if err != nil {
			return nil, err
		}

		err = multipartWriter.Close()
		if err != nil {
			return nil, err
		}
	}

	req := httptest.NewRequest("POST", "/edit_item/", body)
	req = req.WithContext(context.WithValue(req.Context(), "id", int64(1)))
	req.Header.Set("Content-Type", multipartWriter.FormDataContentType())
	resp := httptest.NewRecorder()

	var interactor *item_interactor.TestImplementation
	if hasInteractor {
		interactor = item_interactor.NewTestImplementation()
		item.Image = ""
		interactor.On("EditItem", context.WithValue(req.Context(), "db", nil), int64(1), &item).Return(err)
	}

	return &editItemTestCase{
		req:        req,
		resp:       resp,
		code:       code,
		interactor: interactor,
	}, nil
}

func TestItemController_EditItem(t *testing.T) {
	cases := make([]editItemTestCase, 0)

	cases = append(cases, editItemTestCase{
		req:        httptest.NewRequest("POST", "/edit_item/", nil),
		resp:       httptest.NewRecorder(),
		code:       500,
		interactor: nil,
	})

	testCase, err := buildEditItemTestCase(models.Item{
		ID:             1,
		Name:           "name",
		Weight:         "weight",
		Price:          100,
		AdditionalInfo: "additional_info",
		IsVisible:      true,
		IsAvailable:    true,
		Image:          "image",
		CategoryID:     10,
		OrderParam:     20,
	}, http.StatusSeeOther, true, nil)
	if err != nil {
		t.Error(err)
	}
	cases = append(cases, *testCase)

	testCase, err = buildEditItemTestCase(models.Item{
		ID:             1,
		Name:           "name",
		Weight:         "weight",
		Price:          100,
		AdditionalInfo: "additional_info",
		IsVisible:      true,
		IsAvailable:    true,
		Image:          "image",
		CategoryID:     10,
		OrderParam:     20,
	}, http.StatusInternalServerError, true, errors.New(""))
	if err != nil {
		t.Error(err)
	}
	cases = append(cases, *testCase)

	testCase, err = buildEditItemTestCase(models.Item{
		ID:             1,
		Name:           "name",
		Weight:         "weight",
		Price:          0,
		AdditionalInfo: "additional_info",
		IsVisible:      true,
		IsAvailable:    true,
		Image:          "image",
		CategoryID:     10,
		OrderParam:     20,
	}, http.StatusExpectationFailed, false, nil)
	if err != nil {
		t.Error(err)
	}
	cases = append(cases, *testCase)

	testCase, err = buildEditItemTestCase(models.Item{
		ID:             1,
		Name:           "name",
		Weight:         "weight",
		Price:          100,
		AdditionalInfo: "additional_info",
		IsVisible:      true,
		IsAvailable:    true,
		Image:          "image",
		CategoryID:     0,
		OrderParam:     20,
	}, http.StatusExpectationFailed, false, nil)
	if err != nil {
		t.Error(err)
	}
	cases = append(cases, *testCase)

	testCase, err = buildEditItemTestCase(models.Item{
		ID:             1,
		Name:           "name",
		Weight:         "weight",
		Price:          100,
		AdditionalInfo: "additional_info",
		IsVisible:      true,
		IsAvailable:    true,
		Image:          "image",
		CategoryID:     10,
		OrderParam:     0,
	}, http.StatusExpectationFailed, false, nil)
	if err != nil {
		t.Error(err)
	}
	cases = append(cases, *testCase)

	testCase, err = buildEditItemTestCase(models.Item{
		ID:             0,
		Name:           "name",
		Weight:         "weight",
		Price:          100,
		AdditionalInfo: "additional_info",
		IsVisible:      true,
		IsAvailable:    true,
		Image:          "image",
		CategoryID:     10,
		OrderParam:     20,
	}, http.StatusExpectationFailed, false, nil)
	if err != nil {
		t.Error(err)
	}
	cases = append(cases, *testCase)

	for i, el := range cases {
		controller := NewItemController(el.interactor, nil)
		controller.EditItem(el.resp, el.req)

		if el.resp.Code != el.code {
			t.Error("codes don't match:", i, "expected code", el.code, "got code", el.resp.Code)
		}

		if el.interactor != nil {
			el.interactor.AssertExpectations(t)
		}
	}
}
