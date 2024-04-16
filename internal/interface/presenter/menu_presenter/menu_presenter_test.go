package menu_presenter

import (
	"bytes"
	"github.com/Masterminds/sprig"
	"gitlab.com/maometusu/qr_menu/internal/entity/models"
	"html/template"
	"os"
	"reflect"
	"testing"
)

func TestMenuPresenter_PresentMenus(t *testing.T) {
	golden, err := os.ReadFile("templates/menus_golden.gohtml")
	if err != nil {
		t.Error(err)
	}

	ex := bytes.NewBufferString("")

	tmpl, err := template.New("menus").Funcs(sprig.FuncMap()).ParseFS(templates, "templates/*.gohtml")
	if err != nil {
		t.Error(err)
	}

	err = tmpl.ExecuteTemplate(ex, "menus", map[string]any{
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
				EstablishmentID: 1,
				OrderParam:      2,
			},
		},
		"EstablishmentID": 1,
	})
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(ex.Bytes(), golden) {
		t.Error("results don't match")
	}
}

func generateGolden() {
	tmpl, err := template.New("menus").Funcs(sprig.FuncMap()).ParseFS(templates, "templates/*.gohtml")
	if err != nil {
		panic(err)
	}

	f, err := os.OpenFile("templates/menus_golden.gohtml", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	err = tmpl.ExecuteTemplate(f, "menus", map[string]any{
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
				EstablishmentID: 1,
				OrderParam:      2,
			},
		},
		"EstablishmentID": 1,
	})
	if err != nil {
		panic(err)
	}
}
