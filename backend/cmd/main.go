package main

import (
	"log"
	"net/http"

	"github.com/marianodsr/ecommerce/products"
	"github.com/marianodsr/ecommerce/routes"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/jinzhu/gorm"
)

func main() {
	db, err := gorm.Open("postgres", "user=postgres password=root dbname=ecommerce port=5432 sslmode=disable")
	if err != nil {
		log.Fatal("DATABASE CONNECTION FAILED!")
	}

	defer db.Close()
	products.SetDatabase(db)
	handleRoutes()
}

func handleRoutes() {
	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins: []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type"},
		ExposedHeaders: []string{"Content-Type"},
	}))

	r.Route("/products", routes.ProductListener)
	r.Route("/orders", routes.OrderListener)

	log.Fatal(http.ListenAndServe(":8000", r))
}
