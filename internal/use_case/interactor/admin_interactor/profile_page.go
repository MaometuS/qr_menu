package admin_interactor

import (
	"context"
	"io"
)

func (a *adminInteractor) ProfilePage(context context.Context, w io.Writer, id int64) error {
	profile, err := a.profileRepository.GetOne(context, id)
	if err != nil {
		return err
	}

	return a.presenter.ProfilePage(w, profile)
}
