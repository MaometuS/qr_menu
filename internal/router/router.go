package router

import (
	"gitlab.com/maometusu/qr_menu/internal/controller"
	"net/http"
)

func NewRouter(controller controller.Controller) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /test/", controller.TestPage)
	mux.HandleFunc("GET /assets/", controller.ServeAssets)

	return mux
}
