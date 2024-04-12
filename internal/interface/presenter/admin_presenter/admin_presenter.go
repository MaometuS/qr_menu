package admin_presenter

import (
	"embed"
	"gitlab.com/maometusu/qr_menu/internal/use_case/presenter"
	"html/template"
)

//go:embed templates/*
var templates embed.FS

type adminPresenter struct {
	tmpl *template.Template
}

func NewAdminPresenter() presenter.AdminPresenter {
	tmpl, err := template.ParseFS(templates, "templates/*.gohtml")
	if err != nil {
		panic(err)
	}

	return &adminPresenter{tmpl: tmpl}
}
