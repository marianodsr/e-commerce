package products

import (
	"github.com/jinzhu/gorm"
)

//Product struct
type Product struct {
	gorm.Model
	Name        string  `json:"name"`
	Price       int     `json:"price"`
	Description string  `json:"description"`
	Picture     string  `json:"picture"`
	Order       []Order `gorm:"many2many:orders_products"`
}

//ProductMediator struct
type ProductMediator struct {
	Name        string `json:"name"`
	Price       string `json:"price"`
	Description string `json:"description"`
	Picture     string `json:"picture"`
}

//ProductList type
type ProductList []Product
