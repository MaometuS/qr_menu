package item_presenter

import (
	"bytes"
	"github.com/Masterminds/sprig"
	"gitlab.com/maometusu/qr_menu/internal/entity/models"
	"html/template"
	"os"
	"reflect"
	"testing"
)

func TestItemPresenter_PresentItems(t *testing.T) {
	golden, err := os.ReadFile("templates/items_golden.gohtml")
	if err != nil {
		t.Error(err)
	}

	ex := bytes.NewBufferString("")

	tmpl, err := template.New("").Funcs(sprig.FuncMap()).ParseFS(templates, "templates/*.gohtml")
	if err != nil {
		t.Error(err)
	}

	tmpl.ExecuteTemplate(ex, "items", map[string]any{
		"Items": []models.Item{
			{
				ID:             1,
				Name:           "name",
				Weight:         "weight",
				Price:          10,
				AdditionalInfo: "additional_info",
				IsVisible:      true,
				IsAvailable:    true,
				Image:          "image",
				CategoryID:     1,
				OrderParam:     1,
				SelectedAmount: 0,
			},
		},
		"CategoryID": 1,
		"MenuID":     1,
	})

	if !reflect.DeepEqual(golden, ex.Bytes()) {
		t.Error("results don't match")
	}
}

func generateGolden() {
	tmpl, err := template.New("").Funcs(sprig.FuncMap()).ParseFS(templates, "templates/*.gohtml")
	if err != nil {
		panic(err)
	}

	f, err := os.OpenFile("templates/items_golden.gohtml", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	tmpl.ExecuteTemplate(f, "items", map[string]any{
		"Items": []models.Item{
			{
				ID:             1,
				Name:           "name",
				Weight:         "weight",
				Price:          10,
				AdditionalInfo: "additional_info",
				IsVisible:      true,
				IsAvailable:    true,
				Image:          "image",
				CategoryID:     1,
				OrderParam:     1,
				SelectedAmount: 0,
			},
		},
		"CategoryID": 1,
		"MenuID":     1,
	})
}
