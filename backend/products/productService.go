package products

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
)

//CreateProduct func
func CreateProduct(product *ProductMediator) *Product {

	newPrice, _ := strconv.Atoi(product.Price)

	bytes, err := base64.StdEncoding.DecodeString(product.Picture)
	if err != nil {
		fmt.Println("error decoding b64 line 16:", err)
	}

	finalProduct := &Product{
		Name:        product.Name,
		Price:       newPrice,
		Description: product.Description,
	}

	repo.saveProduct(finalProduct)
	path, _ := filepath.Abs("./../products/images")
	fmt.Println(path)

	pID := finalProduct.ID

	stringID := strconv.FormatUint(uint64(pID), 10)

	finalPath := filepath.Join(path, "producto"+stringID+".jpg")

	f, err := os.Create(finalPath)
	if err != nil {
		fmt.Printf("err line 22: %v", err)
	}
	defer f.Close()

	_, err = f.Write(bytes)
	if err != nil {
		fmt.Printf("err line 28: %v", err)
	}

	f.Sync()

	finalProduct.Picture = finalPath
	repo.db.Save(finalProduct)

	fmt.Printf("\nSE CREO EL PRODUCTO\n")

	return finalProduct

}

//GetAllProducts func
func GetAllProducts() *[]Product {

	products := repo.getAll()

	return products

}

//GetProductsByName func
func GetProductsByName(query string, queryParams string, order string) (ProductList, map[string]uint) {

	products, ids := repo.getByQuery(query, queryParams, order)

	for i := range products {

		encodedPicture := base64encode(products[i].Picture)
		products[i].Picture = encodedPicture

	}

	return products, ids
}

//GetProductByID func
func GetProductByID(id uint) *Product {
	product := repo.getProductByID(id)
	return product
}

//ChangePage func
func ChangePage(query string, queryParams string, order string) (ProductList, map[string]uint) {

	products, ids := repo.getByQuery(query, queryParams, order)

	for i := range products {

		encodedPicture := base64encode(products[i].Picture)
		products[i].Picture = encodedPicture

	}

	return products, ids

}

func base64encode(imgPath string) string {

	f, err := os.Open(imgPath)
	if err != nil {
		fmt.Printf("Error abriendo archivo, linea 88. Err: %v", err)
	}

	reader := bufio.NewReader(f)
	bytes, _ := ioutil.ReadAll(reader)

	encodedPicture := base64.StdEncoding.EncodeToString(bytes)

	return encodedPicture

}
