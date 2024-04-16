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

func TestPublicPresenter_PresentItem(t *testing.T) {
	golden, err := os.ReadFile("templates/item_golden.gohtml")
	if err != nil {
		t.Error(err)
	}

	ex := bytes.NewBufferString("")

	tmpl, err := template.New("").Funcs(sprig.FuncMap()).ParseFS(templates, "templates/*.gohtml")
	if err != nil {
		t.Error(err)
	}

	err = tmpl.ExecuteTemplate(ex, "item", map[string]any{
		"I": models.Item{
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
			SelectedAmount: 1,
		},
		"HasSelected": true,
	})
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(ex.Bytes(), golden) {
		t.Error("results don't match")
	}
}

func generateGoldenItem() {
	tmpl, err := template.ParseFS(templates, "templates/*.gohtml")
	if err != nil {
		panic(err)
	}

	f, err := os.OpenFile("templates/item_golden.gohtml", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	err = tmpl.ExecuteTemplate(f, "item", map[string]any{
		"I": models.Item{
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
			SelectedAmount: 1,
		},
		"HasSelected": true,
	})
	if err != nil {
		panic(err)
	}
}
