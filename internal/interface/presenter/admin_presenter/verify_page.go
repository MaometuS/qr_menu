package admin_presenter

import (
	"io"
)

func (a *adminPresenter) VerifyPage(w io.Writer, id int64, mismatch, hasError bool) error {
	err := a.tmpl.ExecuteTemplate(w, "verify_email", map[string]any{
		"ID":       id,
		"Error":    hasError,
		"Mismatch": mismatch,
	})
	if err != nil {
		return err
	}

	return nil
}
