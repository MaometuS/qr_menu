package admin_presenter

import (
	"bytes"
	"github.com/Masterminds/sprig"
	"gitlab.com/maometusu/qr_menu/internal/entity/models"
	"html/template"
	"os"
	"reflect"
	"testing"
)

func TestAdminPresenter_ProfilePage(t *testing.T) {
	golden, err := os.ReadFile("templates/profile_page_golden.gohtml")
	if err != nil {
		t.Error(err)
	}

	ex := bytes.NewBufferString("")

	tmpl, err := template.New("categories").Funcs(sprig.FuncMap()).ParseFS(templates, "templates/*.gohtml")
	if err != nil {
		t.Error(err)
	}

	err = tmpl.ExecuteTemplate(ex, "profile_page", models.Profile{
		ID:               1,
		Name:             "name",
		Email:            "email",
		Password:         "password",
		Verified:         true,
		VerificationCode: "code",
	})
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(golden, ex.Bytes()) {
		t.Error(err)
	}
}

func generateGoldenProfilePage() {
	tmpl, err := template.ParseFS(templates, "templates/*.gohtml")
	if err != nil {
		panic(err)
	}

	f, err := os.OpenFile("templates/profile_page_golden.gohtml", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}

	err = tmpl.ExecuteTemplate(f, "profile_page", models.Profile{
		ID:               1,
		Name:             "name",
		Email:            "email",
		Password:         "password",
		Verified:         true,
		VerificationCode: "code",
	})
	if err != nil {
		panic(err)
	}
}
