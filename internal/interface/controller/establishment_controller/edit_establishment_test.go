package establishment_controller

import (
	"bytes"
	"context"
	"errors"
	"gitlab.com/maometusu/qr_menu/internal/entity/models"
	"gitlab.com/maometusu/qr_menu/internal/use_case/interactor/establishment_interactor"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

type editEstablishmentTestCase struct {
	req        *http.Request
	resp       *httptest.ResponseRecorder
	code       int
	interactor *establishment_interactor.TestImplementation
}

func buildEditEstablishmentTestCase(item models.Establishment, code int, hasInteractor bool, err error) (*editEstablishmentTestCase, error) {
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

	if item.CurrencyID != 0 {
		err := multipartWriter.WriteField("currency", strconv.FormatInt(item.CurrencyID, 10))
		if err != nil {
			return nil, err
		}
	}

	if item.Phone != "" {
		err := multipartWriter.WriteField("phone", item.Phone)
		if err != nil {
			return nil, err
		}
	}

	if item.Text != "" {
		err := multipartWriter.WriteField("additional_info", item.Text)
		if err != nil {
			return nil, err
		}
	}

	if item.WifiPassword != "" {
		err := multipartWriter.WriteField("wifi_password", item.WifiPassword)
		if err != nil {
			return nil, err
		}
	}

	if item.CanMakeOrders {
		err := multipartWriter.WriteField("can_make_orders", "on")
		if err != nil {
			return nil, err
		}
	}

	if item.Country != "" {
		err := multipartWriter.WriteField("country", item.Country)
		if err != nil {
			return nil, err
		}
	}

	if item.City != "" {
		err := multipartWriter.WriteField("city", item.City)
		if err != nil {
			return nil, err
		}
	}

	if item.Address != "" {
		err := multipartWriter.WriteField("address", item.Address)
		if err != nil {
			return nil, err
		}
	}

	if item.Background != "" {
		fileWriter, err := multipartWriter.CreateFormFile("background", "background.jpg")
		if err != nil {
			return nil, err
		}

		_, err = fileWriter.Write([]byte(item.Background))
		if err != nil {
			return nil, err
		}

		err = multipartWriter.Close()
		if err != nil {
			return nil, err
		}
	}

	if item.Logo != "" {
		fileWriter, err := multipartWriter.CreateFormFile("logo", "logo.jpg")
		if err != nil {
			return nil, err
		}

		_, err = fileWriter.Write([]byte(item.Logo))
		if err != nil {
			return nil, err
		}

		err = multipartWriter.Close()
		if err != nil {
			return nil, err
		}
	}

	req := httptest.NewRequest("POST", "/edit_establishment/", body)
	req = req.WithContext(context.WithValue(req.Context(), "id", int64(1)))
	req.Header.Set("Content-Type", multipartWriter.FormDataContentType())
	resp := httptest.NewRecorder()

	var interactor *establishment_interactor.TestImplementation
	if hasInteractor {
		interactor = establishment_interactor.NewTestImplementation()
		item.ProfileID = 1
		item.Logo = ""
		item.Background = ""
		interactor.On("EditEstablishment", context.WithValue(req.Context(), "db", nil), &item).Return(err)
	}

	return &editEstablishmentTestCase{
		req:        req,
		resp:       resp,
		code:       code,
		interactor: interactor,
	}, nil
}

func TestItemController_EditEstablishment(t *testing.T) {
	cases := make([]editEstablishmentTestCase, 0)

	cases = append(cases, editEstablishmentTestCase{
		req:        httptest.NewRequest("POST", "/create_item/", nil),
		resp:       httptest.NewRecorder(),
		code:       500,
		interactor: nil,
	})

	testCase, err := buildEditEstablishmentTestCase(models.Establishment{
		ID:            1,
		Name:          "name",
		Phone:         "phone",
		Logo:          "logo",
		Background:    "background",
		WifiPassword:  "wifi_password",
		CanMakeOrders: true,
		Country:       "country",
		City:          "city",
		Address:       "address",
		Text:          "text",
		CurrencyID:    29,
	}, http.StatusSeeOther, true, nil)
	if err != nil {
		t.Error(err)
	}
	cases = append(cases, *testCase)

	testCase, err = buildEditEstablishmentTestCase(models.Establishment{
		ID:            1,
		Name:          "name",
		Phone:         "phone",
		Logo:          "logo",
		Background:    "background",
		WifiPassword:  "wifi_password",
		CanMakeOrders: true,
		Country:       "country",
		City:          "city",
		Address:       "address",
		Text:          "text",
		CurrencyID:    29,
	}, http.StatusInternalServerError, true, errors.New(""))
	if err != nil {
		t.Error(err)
	}
	cases = append(cases, *testCase)

	testCase, err = buildEditEstablishmentTestCase(models.Establishment{
		ID:            0,
		Name:          "name",
		Phone:         "phone",
		Logo:          "logo",
		Background:    "background",
		WifiPassword:  "wifi_password",
		CanMakeOrders: true,
		Country:       "country",
		City:          "city",
		Address:       "address",
		Text:          "text",
		CurrencyID:    29,
	}, http.StatusExpectationFailed, false, nil)
	if err != nil {
		t.Error(err)
	}
	cases = append(cases, *testCase)

	testCase, err = buildEditEstablishmentTestCase(models.Establishment{
		ID:            1,
		Name:          "name",
		Phone:         "phone",
		Logo:          "logo",
		Background:    "background",
		WifiPassword:  "wifi_password",
		CanMakeOrders: true,
		Country:       "country",
		City:          "city",
		Address:       "address",
		Text:          "text",
		ProfileID:     10,
		CurrencyID:    0,
	}, http.StatusExpectationFailed, false, nil)
	if err != nil {
		t.Error(err)
	}
	cases = append(cases, *testCase)

	for i, el := range cases {
		controller := NewEstablishmentController(el.interactor, nil)
		controller.EditEstablishment(el.resp, el.req)

		if el.resp.Code != el.code {
			t.Error("codes don't match:", i, "expected code", el.code, "got code", el.resp.Code)
		}

		if el.interactor != nil {
			el.interactor.AssertExpectations(t)
		}
	}
}
