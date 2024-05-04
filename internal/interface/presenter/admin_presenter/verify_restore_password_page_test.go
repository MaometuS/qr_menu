package admin_presenter

import (
	"bytes"
	"github.com/Masterminds/sprig"
	"html/template"
	"os"
	"reflect"
	"testing"
)

func TestAdminPresenter_VerifyRestorePasswordPage(t *testing.T) {
	golden, err := os.ReadFile("templates/verify_restore_password_page_golden.gohtml")
	if err != nil {
		t.Error(err)
	}

	ex := bytes.NewBufferString("")

	tmpl, err := template.New("").Funcs(sprig.FuncMap()).ParseFS(templates, "templates/*.gohtml")
	if err != nil {
		t.Error(err)
	}

	err = tmpl.ExecuteTemplate(ex, "verify_restore_password_page", nil)
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(golden, ex.Bytes()) {
		t.Error(err)
	}
}

func generateGoldenVerifyRestorePasswordPage() {
	tmpl, err := template.ParseFS(templates, "templates/*.gohtml")
	if err != nil {
		panic(err)
	}

	f, err := os.OpenFile("templates/verify_restore_password_page_golden.gohtml", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	err = tmpl.ExecuteTemplate(f, "verify_restore_password_page", nil)
	if err != nil {
		panic(err)
	}
}
