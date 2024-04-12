package admin_presenter

import (
	"io"
)

func (a *adminPresenter) VerifyPage(w io.Writer, id int64) error {
	err := a.tmpl.ExecuteTemplate(w, "verify_email", id)
	if err != nil {
		return err
	}

	return nil
}
