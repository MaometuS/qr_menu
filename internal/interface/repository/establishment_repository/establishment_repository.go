package establishment_repository

import (
	"gitlab.com/maometusu/qr_menu/internal/use_case/repository"
)

type establishmentRepository struct{}

func NewEstablishmentRepository() repository.EstablishmentRepository {
	return &establishmentRepository{}
}
