package category

import "time"

type Category struct {
	ID             uint64    `json:"id"`
	ShopID         uint      `json:"shop_id" validate:"required"`
	ShopExternalId string    `json:"shop_external_id" validate:"required"`
	Active         int8      `json:"active" validate:"min=0,max=1"`
	Name           string    `json:"name" validate:"required,min=3,max=255"`
	UpdatedAt      time.Time `json:"updated_at"`
	CreatedAt      time.Time `json:"created_at"`
}

type CategoryRepository interface {
	DeactivateShopCategories(ShopID uint)
	StoreCategory(cat Category) (Category, error)
	StoreParentCategoryConnection(catId uint64, parentCatId uint64) error
}
