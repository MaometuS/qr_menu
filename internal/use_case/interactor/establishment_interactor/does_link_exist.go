package establishment_interactor

import "context"

func (e *establishmentInteractor) DoesLinkExist(ctx context.Context, link string) (bool, error) {
	return e.repository.DoesLinkExist(ctx, link)
}
