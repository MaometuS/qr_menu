package admin_presenter

import "io"

func (a *adminPresenter) VerifyRestorePasswordPage(w io.Writer, id int64, mismatch, hasError bool) error {
	err := a.tmpl.ExecuteTemplate(w, "verify_restore_password_page", map[string]any{
		"ID":       id,
		"Mismatch": mismatch,
		"HasError": hasError,
	})
	if err != nil {
		return err
	}

	return nil
}
