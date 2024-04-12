package category_presenter

import (
	"embed"
	"github.com/Masterminds/sprig"
	"gitlab.com/maometusu/qr_menu/internal/entity/models"
	"gitlab.com/maometusu/qr_menu/internal/use_case/presenter"
	"html/template"
	"io"
)

//go:embed templates/*
var templates embed.FS

type categoryPresenter struct {
	tmpl *template.Template
}

func (c *categoryPresenter) PresentCategories(w io.Writer, categories []models.Category, menuID int64) error {
	return c.tmpl.ExecuteTemplate(w, "categories", map[string]any{
		"Categories": categories,
		"MenuID":     menuID,
	})
}

func NewCategoryPresenter() presenter.CategoryPresenter {
	tmpl, err := template.New("categories").Funcs(sprig.FuncMap()).ParseFS(templates, "templates/*.gohtml")
	if err != nil {
		panic(err)
	}

	return &categoryPresenter{tmpl}
}
