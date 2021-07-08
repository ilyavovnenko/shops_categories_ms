package attribute

import "time"

type Attribute struct {
	ID           uint64    `json:"id"`
	CategoryID   uint64    `json:"category_id" validate:"required"`
	Type         string    `json:"type" validate:"required,oneof=string text integer float boolean"`
	Level        string    `json:"level" validate:"required,oneof=product variation"`
	Active       int8      `json:"active" validate:"min=0,max=1"`
	Mandatory    int8      `json:"mandatory" validate:"min=0,max=1"`
	Multivalue   int8      `json:"multivalue" validate:"min=0,max=1"`
	Priority     int8      `json:"priority" validate:"min=0,max=1"`
	TechName     string    `json:"tech_name" validate:"required,min=3,max=255"`
	Name         string    `json:"name" validate:"required,min=3,max=255"`
	DefaultValue string    `json:"default_value"`
	Validation   string    `json:"validation"`
	UpdatedAt    time.Time `json:"updated_at"`
	CreatedAt    time.Time `json:"created_at"`
}

type AttributeValue struct {
	ID          uint64    `json:"id"`
	AttributeID uint64    `json:"attribute_id" validate:"required"`
	TechName    string    `json:"tech_name" validate:"required,min=3,max=255"`
	Name        string    `json:"name" validate:"required,min=3,max=255"`
	UpdatedAt   time.Time `json:"updated_at"`
	CreatedAt   time.Time `json:"created_at"`
}

type AttributeRepository interface {
	DeactivateShopAttributes(ShopID uint)
	StoreAttribute(attr Attribute) (Attribute, error)
	StoreAttributeValue(attrVal AttributeValue) (AttributeValue, error)
}
