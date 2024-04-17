package admin_presenter

import "io"

func (a *adminPresenter) RegisterPage(w io.Writer, emailExists, passwordsNotMatch, unexpected bool) error {
	err := a.tmpl.ExecuteTemplate(w, "register_page", map[string]any{
		"EmailExists":       emailExists,
		"PasswordsNotMatch": passwordsNotMatch,
		"Unexpected":        unexpected,
	})
	if err != nil {
		return err
	}

	return nil
}
