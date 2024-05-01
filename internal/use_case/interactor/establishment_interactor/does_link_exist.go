package establishment_interactor

import (
	"context"
	"strings"
)

func (e *establishmentInteractor) DoesLinkExist(ctx context.Context, link string) (bool, error) {
	return e.repository.DoesLinkExist(ctx, strings.ToLower(link))
}
