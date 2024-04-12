package admin_presenter

import (
	"gitlab.com/maometusu/qr_menu/internal/entity/models"
	"io"
)

func (a *adminPresenter) AdminPage(w io.Writer, profile *models.Profile) error {
	err := a.tmpl.ExecuteTemplate(w, "admin_page", profile)
	if err != nil {
		return err
	}

	return nil
}
