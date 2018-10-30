package main

import (
	"encoding/json"
	"fmt"
	"github.com/ahnsv/snappay-server/models"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}
func ProductsHandler(w http.ResponseWriter, r *http.Request) {
	products := models.Products{
		models.Product{
			ID: 1,
		},
		models.Product{
			ID: 2,
		},
	}
	json.NewEncoder(w).Encode(products)
}
func UrlsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Urls!")
}
func ProductByIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productID := vars["productID"]
	fmt.Fprintln(w, "Product show:", productID)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", IndexHandler)
	r.HandleFunc("/products", ProductsHandler)
	r.HandleFunc("/products/{productID}", ProductByIDHandler)
	r.HandleFunc("/urls", UrlsHandler)

	http.Handle("/", r)
	headersOk := handlers.AllowedHeaders([]string{"Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"})

	fmt.Println("Running server!")
	log.Fatal(http.ListenAndServe(":9000", handlers.CORS(originsOk, headersOk, methodsOk)(r)))
}
