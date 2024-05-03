package admin_presenter

import "io"

func (a *adminPresenter) RestorePasswordPage(w io.Writer, noUser, passwordMismatch, unexpected bool) error {
	err := a.tmpl.ExecuteTemplate(w, "restore_password_page", map[string]any{
		"NoUser":           noUser,
		"PasswordMismatch": passwordMismatch,
		"Unexpected":       unexpected,
	})
	if err != nil {
		return err
	}

	return nil
}
