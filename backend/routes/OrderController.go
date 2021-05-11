package routes

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/marianodsr/ecommerce/products"
)

func OrderListener(r chi.Router) {

	r.Get("/test", testGorm)

}

func testGorm(w http.ResponseWriter, r *http.Request) {

	products.Migrate()

}
