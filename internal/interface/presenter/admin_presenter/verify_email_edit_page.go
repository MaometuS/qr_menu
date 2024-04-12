package admin_presenter

import "io"

func (a *adminPresenter) VerifyEmailEditPage(w io.Writer, email string) error {
	return a.tmpl.ExecuteTemplate(w, "verify_email_edit_page", email)
}
