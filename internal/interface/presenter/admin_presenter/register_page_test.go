package admin_presenter

import (
	"bytes"
	"github.com/Masterminds/sprig"
	"html/template"
	"os"
	"reflect"
	"testing"
)

func TestAdminPresenter_RegisterPage(t *testing.T) {
	golden, err := os.ReadFile("templates/register_page_golden.gohtml")
	if err != nil {
		t.Error(err)
	}

	ex := bytes.NewBufferString("")

	tmpl, err := template.New("categories").Funcs(sprig.FuncMap()).ParseFS(templates, "templates/*.gohtml")
	if err != nil {
		t.Error(err)
	}

	err = tmpl.ExecuteTemplate(ex, "register_page", map[string]any{
		"EmailExists":       true,
		"PasswordsNotMatch": true,
		"Unexpected":        true,
	})
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(golden, ex.Bytes()) {
		t.Error(err)
	}
}

func generateGoldenRegisterPage() {
	tmpl, err := template.ParseFS(templates, "templates/*.gohtml")
	if err != nil {
		panic(err)
	}

	f, err := os.OpenFile("templates/register_page_golden.gohtml", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	err = tmpl.ExecuteTemplate(f, "register_page", map[string]any{
		"EmailExists":       true,
		"PasswordsNotMatch": true,
		"Unexpected":        true,
	})
	if err != nil {
		panic(err)
	}
}
