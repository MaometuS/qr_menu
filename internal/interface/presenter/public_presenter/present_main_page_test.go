package public_presenter

import (
	"bytes"
	"github.com/Masterminds/sprig"
	"gitlab.com/maometusu/qr_menu/internal/entity/models"
	"html/template"
	"os"
	"reflect"
	"testing"
)

func TestPublicPresenter_PresentMainPage(t *testing.T) {
	golden, err := os.ReadFile("templates/main_page_golden.gohtml")
	if err != nil {
		t.Error(err)
	}

	ex := bytes.NewBufferString("")

	tmpl, err := template.New("").Funcs(sprig.FuncMap()).ParseFS(templates, "templates/*.gohtml")
	if err != nil {
		t.Error(err)
	}

	err = tmpl.ExecuteTemplate(ex, "main_page", map[string]any{
		"Establishment": models.Establishment{
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
		"Menus": []models.Menu{
			{
				ID:              1,
				Name:            "name",
				IsVisible:       true,
				EstablishmentID: 1,
				OrderParam:      1,
			},
			{
				ID:              2,
				Name:            "name2",
				IsVisible:       true,
				EstablishmentID: 2,
				OrderParam:      2,
			},
		},
		"MenuID":      1,
		"HasSelected": true,
	})
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(ex.Bytes(), golden) {
		t.Error("results don't match")
	}
}

func generateGoldenMainPage() {
	tmpl, err := template.ParseFS(templates, "templates/*.gohtml")
	if err != nil {
		panic(err)
	}

	f, err := os.OpenFile("templates/main_page_golden.gohtml", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	err = tmpl.ExecuteTemplate(f, "main_page", map[string]any{
		"Establishment": models.Establishment{
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
		"Menus": []models.Menu{
			{
				ID:              1,
				Name:            "name",
				IsVisible:       true,
				EstablishmentID: 1,
				OrderParam:      1,
			},
			{
				ID:              2,
				Name:            "name2",
				IsVisible:       true,
				EstablishmentID: 2,
				OrderParam:      2,
			},
		},
		"MenuID":      1,
		"HasSelected": true,
	})
	if err != nil {
		panic(err)
	}
}
