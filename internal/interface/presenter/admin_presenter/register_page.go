package admin_presenter

import "io"

func (a *adminPresenter) RegisterPage(w io.Writer) error {
	err := a.tmpl.ExecuteTemplate(w, "register_page", nil)
	if err != nil {
		return err
	}

	return nil
}
