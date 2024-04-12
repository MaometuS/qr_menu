package admin_interactor

import "context"

func (a *adminInteractor) EditName(ctx context.Context, email, name string, id int64) (bool, error) {
	profile, err := a.profileRepository.GetOne(ctx, id)
	if err != nil {
		return false, err
	}

	if name != "" {
		profile.Name = name
	}

	err = a.profileRepository.Update(ctx, profile)
	if err != nil {
		return false, err
	}

	return email != "" && profile.Email != email, nil
}
