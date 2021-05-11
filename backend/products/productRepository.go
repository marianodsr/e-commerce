package products

import (
	"github.com/jinzhu/gorm"
	//ASasd
	_ "github.com/lib/pq"
)

//Repository struct
type Repository struct {
	db *gorm.DB
}

var repo Repository

//ProductRepository interace
type ProductRepository interface {
	saveProduct()
	deleteProduct()
	getAll()
	getProductById()
}

//SetDatabase func
func SetDatabase(db *gorm.DB) {

	repo.db = db
}

func (repo *Repository) saveProduct(p *Product) {
	repo.db.Save(p)
}

func (repo *Repository) getAll() *[]Product {

	productList := []Product{}
	repo.db.Find(&productList)
	return &productList
}

func (repo *Repository) getByQuery(query string, queryParams string, order string) (ProductList, map[string]uint) {

	productList := []Product{}
	var count uint

	repo.db.Model(&Product{}).Where(query, queryParams).Count(&count)
	repo.db.Where(query, queryParams).Limit(12).Order(order).Find(&productList)

	filteredCount := len(productList) - 1

	ids := map[string]uint{
		"count":   count,
		"firstID": productList[0].ID,
		"lastID":  productList[filteredCount].ID,
	}

	return productList, ids
}

func (repo *Repository) getProductByID(id uint) *Product {

	product := &Product{}
	repo.db.First(product, id)

	return product

}
