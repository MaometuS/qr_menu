package admin_presenter

import "io"

func (a *adminPresenter) LoginPage(w io.Writer) error {
	err := a.tmpl.ExecuteTemplate(w, "login_page", nil)
	if err != nil {
		return err
	}

	return nil
}
