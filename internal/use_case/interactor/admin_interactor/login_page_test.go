package admin_interactor

import (
	"context"
	"errors"
	"gitlab.com/maometusu/qr_menu/internal/use_case/presenter"
	"testing"
)

func TestAdminInteractor_LoginPage(t *testing.T) {
	pres := &presenter.AdminPresenterMock{}
	pres.On("LoginPage", nil, false, false).Return(errors.New("error"))

	interactor := NewAdminInteractor(pres, nil, nil, nil)
	err := interactor.LoginPage(context.Background(), nil)
	if err.Error() != "error" {
		t.Error("errors don't match")
	}

	pres.AssertExpectations(t)
}
