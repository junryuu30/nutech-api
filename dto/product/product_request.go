package productdto

type ProductRequest struct {
	Name         string `json:"name" form:"name" gorm:"type: varchar(255)"`
	PriceGoods   int    `json:"pricegoods" form:"pricegoods" gorm:"type: int"`
	SellingPrice int    `json:"sellingprice" form:"sellingprice" gorm:"type: int"`
	Image        string `json:"image" form:"image" gorm:"type: varchar(255)"`
	Stock        int    `json:"stock" form:"stock" gorm:"type: int"`
	UserID       int    `json:"user_id" gorm:"type: int"`
}

type UpdateProductRequest struct {
	Name         string `json:"name" form:"name" gorm:"type: varchar(255)"`
	PriceGoods   int    `json:"pricegoods" form:"pricegoods" gorm:"type: int"`
	SellingPrice int    `json:"sellingprice" form:"sellingprice" gorm:"type: int"`
	Image        string `json:"image" form:"image" gorm:"type: varchar(255)"`
	Stock        int    `json:"stock" form:"stock" gorm:"type: int"`
}
