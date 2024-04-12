package menu_repository

import (
	"gitlab.com/maometusu/qr_menu/internal/use_case/repository"
)

type menuRepository struct {
}

func NewMenuRepository() repository.MenuRepository {
	return &menuRepository{}
}
