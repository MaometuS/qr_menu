package common_controller

import (
	"gitlab.com/maometusu/qr_menu/internal/entity"
	"gitlab.com/maometusu/qr_menu/internal/interface/controller/common_controller/static"
	"html/template"
	"log"
	"net/http"
)

type CommonController interface {
	TestPage(w http.ResponseWriter, r *http.Request)
	ServeAssets(w http.ResponseWriter, r *http.Request)
}

type controller struct {
	db          entity.PgxIface
	tmpl        *template.Template
	assetServer http.Handler
}

func NewCommonController(db entity.PgxIface) CommonController {
	tmpl, err := template.ParseFS(static.Templates, "templates/*.gohtml")
	if err != nil {
		log.Fatal(err)
	}

	return &controller{
		db:          db,
		tmpl:        tmpl,
		assetServer: http.FileServer(http.FS(static.Assets)),
	}
}
