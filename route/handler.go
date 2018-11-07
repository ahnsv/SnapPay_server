package route

import (
	"encoding/json"
	"fmt"
	"github.com/ahnsv/snappay-server/models"
	"github.com/gorilla/mux"
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
