package repository

import (
	"github.com/ilyavovnenko/shops_categories_ms/internal/attribute"
	"github.com/ilyavovnenko/shops_categories_ms/internal/category"
	"github.com/ilyavovnenko/shops_categories_ms/internal/shop"
)

type RepoCollection struct {
	CategoryRepo  category.CategoryRepository
	AttributeRepo attribute.AttributeRepository
	ShopRepo      shop.ShopRepository
}
