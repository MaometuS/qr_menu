package item_presenter

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

type itemPresenter struct {
	tmpl *template.Template
}

func (i *itemPresenter) PresentItems(w io.Writer, items []models.Item, categoryID, menuID int64) error {
	return i.tmpl.ExecuteTemplate(w, "items", map[string]any{
		"Items":      items,
		"CategoryID": categoryID,
		"MenuID":     menuID,
	})
}

func NewItemPresenter() presenter.ItemPresenter {
	tmpl, err := template.New("").Funcs(sprig.FuncMap()).ParseFS(templates, "templates/*.gohtml")
	if err != nil {
		panic(err)
	}

	return &itemPresenter{tmpl: tmpl}
}
