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

func TestPublicPresenter_PresentSearch(t *testing.T) {
	golden, err := os.ReadFile("templates/search_golden.gohtml")
	if err != nil {
		t.Error(err)
	}

	ex := bytes.NewBufferString("")

	tmpl, err := template.New("").Funcs(sprig.FuncMap()).ParseFS(templates, "templates/*.gohtml")
	if err != nil {
		t.Error(err)
	}

	err = tmpl.ExecuteTemplate(ex, "search", map[string]any{
		"Search": "search",
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
				SelectedAmount: 1,
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

func generateGoldenSearch() {
	tmpl, err := template.ParseFS(templates, "templates/*.gohtml")
	if err != nil {
		panic(err)
	}

	f, err := os.OpenFile("templates/search_golden.gohtml", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	err = tmpl.ExecuteTemplate(f, "search", map[string]any{
		"Search": "search",
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
				SelectedAmount: 1,
			},
		},
		"EstablishmentID": 1,
		"MenuID":          1,
	})
	if err != nil {
		panic(err)
	}
}
