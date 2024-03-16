package controller

import (
	"html/template"
	"net/http"
)

func (c *controller) TestPage(w http.ResponseWriter, r *http.Request) {
	err := presentTestPage(w, c.tmpl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func presentTestPage(w http.ResponseWriter, tmpl *template.Template) error {
	return tmpl.ExecuteTemplate(w, "test_page", nil)
}
