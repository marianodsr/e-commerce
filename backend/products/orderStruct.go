package products

import (
	"github.com/jinzhu/gorm"
)

//Order Struct
type Order struct {
	gorm.Model
	ClientName   string
	ClientEmail  string
	ClientAdress string
	Products     []Product `gorm:"many2many:orders_products"`
}
