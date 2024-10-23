package main

import (
	"encoding/json"
	"go-crud/models"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func main() {
	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database")
	}

	db.AutoMigrate(&models.Product{})

	r := mux.NewRouter()

	r.HandleFunc("/products", GetProducts).Methods("GET")
	r.HandleFunc("/products/{id}", GetProduct).Methods("GET")
	r.HandleFunc("/products", CreateProduct).Methods("POST")
	r.HandleFunc("/products/{id}", UpdateProduct).Methods("PUT")
	r.HandleFunc("/products/{id}", DeleteProduct).Methods("DELETE")

	log.Println("Server running on port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}

func GetProducts(w http.ResponseWriter, r *http.Request) {
	var products []models.Product

	db.Find(&products)

	json.NewEncoder(w).Encode(products)
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product

	json.NewDecoder(r.Body).Decode(&product)

	db.Create(&product)

	json.NewEncoder(w).Encode(product)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product

	json.NewDecoder(r.Body).Decode(&product)

	db.Create(&product)

	json.NewEncoder(w).Encode(product)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	var product models.Product

	db.First(&product, params["id"])

	json.NewDecoder(r.Body).Decode(&product)

	db.Save(&product)

	json.NewEncoder(w).Encode(product)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	var product models.Product

	db.Delete(&product, params["id"])

	json.NewEncoder(w).Encode("Product deleted")
}
