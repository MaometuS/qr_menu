package category_repository

import (
	"gitlab.com/maometusu/qr_menu/internal/use_case/repository"
)

type categoryRepository struct{}

func NewCategoryRepository() repository.CategoryRepository {
	return &categoryRepository{}
}
