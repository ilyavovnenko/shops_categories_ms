package shop

import "time"

const AmazonCOM = "amazon.com"
const AmazonDE = "amazon.de"
const AmazonNL = "amazon.nl"
const EbayCOM = "ebay.com"
const EbayDE = "ebay.de"
const EbayNL = "ebay.nl"
const BolCom = "bol.com"

type Shop struct {
	ID         uint      `json:"id"`
	ShopTypeID uint      `json:"shop_type_id" validate:"required"`
	Name       string    `json:"name" validate:"required,min=3,max=255"`
	UpdatedAt  time.Time `json:"updated_at"`
	CreatedAt  time.Time `json:"created_at"`
}

type ShopRepository interface {
	GetIdByName(name string) (uint, error)
}
