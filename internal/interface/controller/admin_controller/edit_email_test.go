package admin_controller

import (
	"gitlab.com/maometusu/qr_menu/internal/use_case/interactor/admin_interactor"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestController_EditEmail(t *testing.T) {
	type testCase struct {
		req        *http.Request
		resp       *httptest.ResponseRecorder
		code       int
		interactor *admin_interactor.TestImplementation
	}

	cases := make([]testCase, 0)

	cases = append(cases, testCase{
		req:        httptest.NewRequest("POST", "https://localhost/change_password", nil),
		resp:       httptest.NewRecorder(),
		code:       500,
		interactor: nil,
	})
}
