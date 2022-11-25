package models

type Product struct {
	ID           int    `json:"id" gorm:"primary_key:auto_increment"`
	Name         string `json:"name" gorm:"unique"`
	PriceGoods   int    `json:"pricegoods" form:"pricegoods" gorm:"type: int"`
	SellingPrice int    `json:"sellingprice" form:"sellingprice" gorm:"type: int"`
	Image        string `json:"image" form:"image" gorm:"type: varchar(255)"`
	Stock        int    `json:"stock" form:"stock" gorm:"type: int"`
}
