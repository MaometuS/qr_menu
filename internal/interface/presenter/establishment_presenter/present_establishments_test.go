package establishment_presenter

import (
	"bytes"
	"gitlab.com/maometusu/qr_menu/internal/entity/models"
	"html/template"
	"os"
	"reflect"
	"testing"
)

func TestEstablishmentPresenter_PresentEstablishments(t *testing.T) {
	golden, err := os.ReadFile("templates/establishments_golden.gohtml")
	if err != nil {
		t.Error(err)
	}

	ex := bytes.NewBufferString("")

	tmpl, err := template.ParseFS(templates, "templates/*.gohtml")
	if err != nil {
		t.Error(err)
	}

	err = tmpl.ExecuteTemplate(ex, "establishments_page", map[string]any{
		"Establishments": []models.Establishment{
			{
				ID:            1,
				Name:          "name1",
				Phone:         "123",
				Logo:          "logo1",
				Background:    "background1",
				WifiPassword:  "password1",
				CanMakeOrders: true,
				Country:       "uz",
				City:          "xiva",
				Address:       "address1",
				Text:          "text1",
				ProfileID:     1,
				CurrencyID:    1,
				Link:          "link1",
			},
			{
				ID:            2,
				Name:          "name2",
				Phone:         "432",
				Logo:          "logo2",
				Background:    "background2",
				WifiPassword:  "password2",
				CanMakeOrders: false,
				Country:       "ru",
				City:          "piter",
				Address:       "address2",
				Text:          "text2",
				ProfileID:     1,
				CurrencyID:    1,
				Link:          "link2",
			},
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

func generateGoldenEstablishments() {
	tmpl, err := template.ParseFS(templates, "templates/*.gohtml")
	if err != nil {
		panic(err)
	}

	f, err := os.OpenFile("templates/establishments_golden.gohtml", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	err = tmpl.ExecuteTemplate(f, "establishments_page", map[string]any{
		"Establishments": []models.Establishment{
			{
				ID:            1,
				Name:          "name1",
				Phone:         "123",
				Logo:          "logo1",
				Background:    "background1",
				WifiPassword:  "password1",
				CanMakeOrders: true,
				Country:       "uz",
				City:          "xiva",
				Address:       "address1",
				Text:          "text1",
				ProfileID:     1,
				CurrencyID:    1,
				Link:          "link1",
			},
			{
				ID:            2,
				Name:          "name2",
				Phone:         "432",
				Logo:          "logo2",
				Background:    "background2",
				WifiPassword:  "password2",
				CanMakeOrders: false,
				Country:       "ru",
				City:          "piter",
				Address:       "address2",
				Text:          "text2",
				ProfileID:     1,
				CurrencyID:    1,
				Link:          "link2",
			},
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
