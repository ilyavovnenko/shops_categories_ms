package usecase

import (
	"context"

	"github.com/ilyavovnenko/shops_categories_ms/internal/category"
)

type CategoryUsecase struct {
	catRepo category.CategoryRepository
}

func New(catRepo category.CategoryRepository) *CategoryUsecase {
	return &CategoryUsecase{
		catRepo: catRepo,
	}
}

func (cu *CategoryUsecase) GetAll(cTimeout context.Context, page int, perPage uint16) ([]category.Category, error) {
	return cu.catRepo.GetAll(cTimeout, perPage, page)
}

func (cu *CategoryUsecase) GetByID(cTimeout context.Context, id int64) (category.Category, error) {
	res, err := cu.catRepo.GetByID(cTimeout, id)
	if err != nil {
		return category.Category{}, err
	}

	return res, nil
}
