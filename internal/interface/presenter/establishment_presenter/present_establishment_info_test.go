package establishment_presenter

import (
	"bytes"
	"gitlab.com/maometusu/qr_menu/internal/entity/models"
	"html/template"
	"os"
	"reflect"
	"testing"
)

func TestEstablishmentPresenter_PresentEstablishmentInfo(t *testing.T) {
	golden, err := os.ReadFile("templates/establishment_info_golden.gohtml")
	if err != nil {
		t.Error(err)
	}

	ex := bytes.NewBufferString("")

	tmpl, err := template.ParseFS(templates, "templates/*.gohtml")
	if err != nil {
		t.Error(err)
	}

	err = tmpl.ExecuteTemplate(ex, "establishment_info", map[string]any{
		"E": models.Establishment{
			ID:            1,
			Name:          "name",
			Phone:         "123321",
			Logo:          "logo",
			Background:    "background",
			WifiPassword:  "password",
			CanMakeOrders: true,
			Country:       "country",
			City:          "city",
			Address:       "address",
			Text:          "text",
			ProfileID:     1,
			CurrencyID:    1,
			Link:          "link",
		},
		"Currencies": []models.Currency{
			{
				ID:   1,
				Name: "Sum",
				Code: "UZS",
			},
		},
	})
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(golden, ex.Bytes()) {
		t.Error("results don't match")
	}
}

func generateGoldenEstablishmentInfo() {
	tmpl, err := template.ParseFS(templates, "templates/*.gohtml")
	if err != nil {
		panic(err)
	}

	f, err := os.OpenFile("templates/establishment_info_golden.gohtml", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	err = tmpl.ExecuteTemplate(f, "establishment_info", map[string]any{
		"E": models.Establishment{
			ID:            1,
			Name:          "name",
			Phone:         "123321",
			Logo:          "logo",
			Background:    "background",
			WifiPassword:  "password",
			CanMakeOrders: true,
			Country:       "country",
			City:          "city",
			Address:       "address",
			Text:          "text",
			ProfileID:     1,
			CurrencyID:    1,
			Link:          "link",
		},
		"Currencies": []models.Currency{
			{
				ID:   1,
				Name: "Sum",
				Code: "UZS",
			},
		},
	})
	if err != nil {
		panic(err)
	}
}
