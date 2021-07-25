package usecase

import (
	"context"

	"github.com/ilyavovnenko/shops_categories_ms/internal/attribute"
)

type AttributeUsecase struct {
	attrRepo attribute.AttributeRepository
}

func New(attrRepo attribute.AttributeRepository) *AttributeUsecase {
	return &AttributeUsecase{
		attrRepo: attrRepo,
	}
}

func (au *AttributeUsecase) GetAllAttributes(cTimeout context.Context, page int, perPage uint16) ([]attribute.Attribute, error) {
	return au.attrRepo.GetAllAttributes(cTimeout, perPage, page)
}

func (au *AttributeUsecase) GetAttributeByID(cTimeout context.Context, id int64) (attribute.Attribute, error) {
	res, err := au.attrRepo.GetAttributeByID(cTimeout, id)
	if err != nil {
		return attribute.Attribute{}, err
	}

	return res, nil
}
