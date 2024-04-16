package admin_presenter

import (
	"bytes"
	"github.com/Masterminds/sprig"
	"html/template"
	"os"
	"reflect"
	"testing"
)

func TestAdminPresenter_VerifyEmailEditPage(t *testing.T) {
	golden, err := os.ReadFile("templates/verify_email_edit_page_golden.gohtml")
	if err != nil {
		t.Error(err)
	}

	ex := bytes.NewBufferString("")

	tmpl, err := template.New("").Funcs(sprig.FuncMap()).ParseFS(templates, "templates/*.gohtml")
	if err != nil {
		t.Error(err)
	}

	err = tmpl.ExecuteTemplate(ex, "verify_email_edit_page", "email@email.email")
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(golden, ex.Bytes()) {
		t.Error(err)
	}
}

func generateGoldenVerifyEmailEditPage() {
	tmpl, err := template.ParseFS(templates, "templates/*.gohtml")
	if err != nil {
		panic(err)
	}

	f, err := os.OpenFile("templates/verify_email_edit_page_golden.gohtml", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	err = tmpl.ExecuteTemplate(f, "verify_email_edit_page", "email@email.email")
	if err != nil {
		panic(err)
	}
}
