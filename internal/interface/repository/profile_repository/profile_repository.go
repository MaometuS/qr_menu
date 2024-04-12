package profile_repository

import (
	"gitlab.com/maometusu/qr_menu/internal/use_case/repository"
)

type profileRepository struct{}

func NewProfileRepository() repository.ProfileRepository {
	return &profileRepository{}
}
