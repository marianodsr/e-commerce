package products

import "github.com/jinzhu/gorm"

type orderRepository struct {
	db *gorm.DB
}

func (repo *orderRepository) saveOrder(order *Order) {

	repo.db.Save(order)
}

func (repo *orderRepository) migrate() {

	repo.db.LogMode(true)
	repo.db.AutoMigrate(&Order{}, &Product{})

}
