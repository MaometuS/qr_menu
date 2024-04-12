package item_repository

import (
	"gitlab.com/maometusu/qr_menu/internal/use_case/repository"
)

type itemRepository struct{}

func NewItemRepository() repository.ItemRepository {
	return &itemRepository{}
}
