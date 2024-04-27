package admin_controller

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestController_Logout(t *testing.T) {
	req := httptest.NewRequest("GET", "http://localhost/logout", nil)
	resp := httptest.NewRecorder()

	controller := NewAdminController(nil, nil)
	controller.Logout(resp, req)

	if resp.Code != http.StatusPermanentRedirect {
		t.Error("wrong code")
	}
}
