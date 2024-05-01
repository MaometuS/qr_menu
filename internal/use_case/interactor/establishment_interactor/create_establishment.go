package establishment_interactor

import (
	"context"
	"gitlab.com/maometusu/qr_menu/internal/entity/models"
	"strings"
)

func (e *establishmentInteractor) CreateEstablishment(context context.Context, establishment *models.Establishment) error {
	establishment.Link = strings.ToLower(establishment.Link)
	return e.repository.CreateEstablishment(context, establishment)
}
