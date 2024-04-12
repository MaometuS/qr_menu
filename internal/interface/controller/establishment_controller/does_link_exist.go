package establishment_controller

import (
	"context"
	"net/http"
	"strconv"
)

func (c *controller) DoesLinkExist(w http.ResponseWriter, r *http.Request) {
	exist, err := c.interactor.DoesLinkExist(context.WithValue(r.Context(), "db", c.db), r.URL.Query().Get("link"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(strconv.FormatBool(exist)))
}
