package establishment_interactor

import (
	"context"
	"gitlab.com/maometusu/qr_menu/internal/entity/models"
)

func (e *establishmentInteractor) CreateEstablishment(context context.Context, establishment *models.Establishment) error {
	return e.repository.CreateEstablishment(context, establishment)
}
