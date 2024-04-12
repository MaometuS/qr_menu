package menu_presenter

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

type menuPresenter struct {
	tmpl *template.Template
}

func (m *menuPresenter) PresentMenus(w io.Writer, menus []models.Menu, establishmentID int64) error {
	err := m.tmpl.ExecuteTemplate(w, "menus", map[string]any{
		"Menus":           menus,
		"EstablishmentID": establishmentID,
	})
	if err != nil {
		return err
	}

	return nil
}

func NewMenuPresenter() presenter.MenuPresenter {
	tmpl, err := template.New("menus").Funcs(sprig.FuncMap()).ParseFS(templates, "templates/*.gohtml")
	if err != nil {
		panic(err)
	}

	return &menuPresenter{tmpl: tmpl}
}
