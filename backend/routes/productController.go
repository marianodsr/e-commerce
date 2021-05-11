package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi"

	"github.com/marianodsr/ecommerce/payments"
	"github.com/marianodsr/ecommerce/products"
)

type mix struct {
	products *[]products.Product
	image    *os.File
}

//ProductListener func
func ProductListener(r chi.Router) {

	r.Post("/create", createProduct)
	r.Post("/getAllByName", getProducts)
	r.Post("/changePage", changePage)
	r.Post("/attemptPayment", attemptPayment)

}

func createProduct(w http.ResponseWriter, r *http.Request) {

	product := &products.ProductMediator{}

	err := json.NewDecoder(r.Body).Decode(product)
	if err != nil {
		fmt.Printf("Error en la decodificación: %v", err)
		w.WriteHeader(500)
		return
	}

	products.CreateProduct(product)
	fmt.Println("Se creó el producto exitosamente!")
	w.WriteHeader(200)

}

//QueryStruy
type queryStruct struct {
	Query   string `json:"query"`
	IDParam uint   `json:"id_param"`
	Type    string `json:"type"`
}

type productLister struct {
	Products products.ProductList `json:"products"`
	Ids      map[string]uint      `json:"id_map"`
	Type     string               `json:"type"`
}

func getProducts(w http.ResponseWriter, r *http.Request) {

	query := &queryStruct{}

	err := json.NewDecoder(r.Body).Decode(query)
	if err != nil {
		fmt.Printf("error en la decodificación: %v", err)
		return
	}

	newQuery := "name LIKE ?"
	queryParams := "%" + query.Query + "%"

	products, ids := products.GetProductsByName(newQuery, queryParams, "id asc")

	finalProducts := &productLister{
		Products: products,
		Ids:      ids,
	}

	encoder := json.NewEncoder(w)

	encoder.Encode(finalProducts)

}

func changePage(w http.ResponseWriter, r *http.Request) {

	query := &queryStruct{}
	order := "id asc"
	newQuery := ""

	err := json.NewDecoder(r.Body).Decode(query)
	if err != nil {
		fmt.Printf("Error line 95: %v", err)
	}

	if query.Type == "next" {
		newQuery = fmt.Sprintf("name LIKE ? AND id > %v", query.IDParam)
		order = "id asc"

	} else if query.Type == "previous" {

		newQuery = fmt.Sprintf("name LIKE ? AND id < %v", query.IDParam)
		order = "id desc"
	}

	queryParams := "%" + query.Query + "%"

	newProducts, ids := products.ChangePage(newQuery, queryParams, order)

	finalProducts := &productLister{
		Products: newProducts,
		Ids:      ids,
	}

	json.NewEncoder(w).Encode(finalProducts)

}

func attemptPayment(w http.ResponseWriter, r *http.Request) {

	items := products.ItemList{}

	err := json.NewDecoder(r.Body).Decode(&items)

	if err != nil {
		fmt.Printf("Error en la decodificación, linea 128 :%v", err)
		return
	}

	clientSecret, err := payments.AttemptPayment(items)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(clientSecret)

	w.WriteHeader(http.StatusOK)

}
