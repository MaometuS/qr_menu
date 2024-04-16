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

func TestPublicPresenter_PresentCategories(t *testing.T) {
	golden, err := os.ReadFile("templates/categories_golden.gohtml")
	if err != nil {
		t.Error(err)
	}

	ex := bytes.NewBufferString("")

	tmpl, err := template.New("").Funcs(sprig.FuncMap()).ParseFS(templates, "templates/*.gohtml")
	if err != nil {
		t.Error(err)
	}

	err = tmpl.ExecuteTemplate(ex, "categories", map[string]any{
		"Menu": "menu",
		"Categories": []models.Category{
			{
				ID:         1,
				Name:       "name",
				Background: "background",
				MenuID:     1,
				IsVisible:  true,
				OrderParam: 1,
			},
			{
				ID:         2,
				Name:       "name2",
				Background: "background2",
				MenuID:     2,
				IsVisible:  true,
				OrderParam: 2,
			},
		},
		"EstablishmentID": 1,
		"MenuID":          1,
	})
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(ex.Bytes(), golden) {
		t.Error("results don't match")
	}
}

func generateGoldenCategories() {
	tmpl, err := template.ParseFS(templates, "templates/*.gohtml")
	if err != nil {
		panic(err)
	}

	f, err := os.OpenFile("templates/categories_golden.gohtml", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	err = tmpl.ExecuteTemplate(f, "categories", map[string]any{
		"Menu": "menu",
		"Categories": []models.Category{
			{
				ID:         1,
				Name:       "name",
				Background: "background",
				MenuID:     1,
				IsVisible:  true,
				OrderParam: 1,
			},
			{
				ID:         2,
				Name:       "name2",
				Background: "background2",
				MenuID:     2,
				IsVisible:  true,
				OrderParam: 2,
			},
		},
		"EstablishmentID": 1,
		"MenuID":          1,
	})
	if err != nil {
		panic(err)
	}
}
