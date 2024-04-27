package admin_interactor

import (
	"context"
	"errors"
	"gitlab.com/maometusu/qr_menu/internal/use_case/presenter"
	"testing"
)

func TestAdminInteractor_RegisterPage(t *testing.T) {
	pres := &presenter.AdminPresenterMock{}
	pres.On("RegisterPage", nil, false, false, false).Return(errors.New("error"))

	interactor := NewAdminInteractor(pres, nil, nil, nil)
	err := interactor.RegisterPage(context.Background(), nil)
	if err.Error() != "error" {
		t.Error("results don't match")
	}

	pres.AssertExpectations(t)
}
