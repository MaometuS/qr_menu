package public_presenter

import (
	"embed"
	"gitlab.com/maometusu/qr_menu/internal/use_case/presenter"
	"html/template"
)

//go:embed templates/*
var templates embed.FS

type publicPresenter struct {
	tmpl *template.Template
}

func NewPublicPresenter() presenter.PublicPresenter {
	tmpl, err := template.ParseFS(templates, "templates/*.gohtml")
	if err != nil {
		panic(err)
	}

	return &publicPresenter{tmpl: tmpl}
}
