package controller

import (
	"gitlab.com/maometusu/qr_menu/internal/static"
	"html/template"
	"net/http"
)

type Controller interface {
	TestPage(w http.ResponseWriter, r *http.Request)
	ServeAssets(w http.ResponseWriter, r *http.Request)
}

type controller struct {
	tmpl        *template.Template
	assetServer http.Handler
}

func NewController() Controller {
	tmpl, err := template.ParseFS(static.Templates, "templates/*.gohtml")
	if err != nil {
		panic(err)
	}

	return &controller{
		tmpl:        tmpl,
		assetServer: http.FileServer(http.FS(static.Assets)),
	}
}
