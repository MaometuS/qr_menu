package common_controller

import "net/http"

func (c *controller) ServeAssets(w http.ResponseWriter, r *http.Request) {
	c.assetServer.ServeHTTP(w, r)
}
