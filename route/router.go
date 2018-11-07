package route

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func router() {
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
