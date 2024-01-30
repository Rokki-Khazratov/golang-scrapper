// main.go
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"api/data-scrappers/scrapper"
)

func HandleRequests() {
	r := mux.NewRouter()
	c := cors.AllowAll()
	r.HandleFunc("/api/laptops/", GetAllLaptops).Methods("GET")
	r.HandleFunc("/api/laptop/{id}/", GetProductDetail).Methods("GET")

	http.Handle("/", c.Handler(r))
	http.ListenAndServe(":8080", nil)
}

func main() {
	scrapper.ScrapeBrostore()
	HandleRequests()
}

func GetAllLaptops(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	jsonData, err := json.Marshal(scrapper.Products)
	if err != nil {
		http.Error(w, "Error converting to JSON", http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)
}

func GetProductDetail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id := params["id"]

	productID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	if productID < 1 || productID > len(scrapper.Products) {
		http.Error(w, "Product ID out of range", http.StatusNotFound)
		return
	}

	link := scrapper.Products[productID-1].Link
	detailedProduct, err := scrapper.ScrapeBrostoreDetail(productID, link)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error getting product detail: %v", err), http.StatusInternalServerError)
		return
	}

	jsonData, err := json.Marshal(detailedProduct)
	if err != nil {
		http.Error(w, "Error converting to JSON", http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)
}
