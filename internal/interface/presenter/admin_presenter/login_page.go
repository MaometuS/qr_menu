package admin_presenter

import "io"

func (a *adminPresenter) LoginPage(w io.Writer, noMatch, unexpected bool) error {
	err := a.tmpl.ExecuteTemplate(w, "login_page", map[string]any{
		"NoMatch":    noMatch,
		"Unexpected": unexpected,
	})
	if err != nil {
		return err
	}

	return nil
}
