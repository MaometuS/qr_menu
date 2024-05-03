package admin_interactor

import (
	"context"
	"errors"
	"gitlab.com/maometusu/qr_menu/internal/use_case/presenter"
	"testing"
)

func TestAdminInteractor_VerifyRestorePasswordPage(t *testing.T) {
	pres := presenter.NewAdminPresenterMock()
	pres.On("VerifyRestorePasswordPage", nil, int64(1), false, false).Return(errors.New("error"))

	inter := NewAdminInteractor(pres, nil, nil, nil)
	err := inter.VerifyRestorePasswordPage(context.Background(), nil, 1)
	if err == nil {
		t.Error("must have error")
	}
	if err != nil && err.Error() != "error" {
		t.Error("error mismatch")
	}

	pres = presenter.NewAdminPresenterMock()
	pres.On("VerifyRestorePasswordPage", nil, int64(1), false, false).Return(nil)

	inter = NewAdminInteractor(pres, nil, nil, nil)
	err = inter.VerifyRestorePasswordPage(context.Background(), nil, 1)
	if err != nil {
		t.Error("has error")
	}

	pres.AssertExpectations(t)
}
