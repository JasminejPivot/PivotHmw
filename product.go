package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

type products struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/product/{id}", getproducts).Methods("GET", "POST")
	log.Fatal(http.ListenAndServe(":8080", r))

}

func getproducts(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	ID := vars["id"]

	idInt, err := strconv.Atoi(ID)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	product, err := os.ReadFile("./products.json")
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var items []products

	err = json.Unmarshal(product, &items)
	if err != nil {
		log.Println(err)

	}

	for _, product := range items {
		if product.ID == idInt {
			response, err := json.Marshal(product)
			if err != nil {
				log.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusOK)
			_, err = w.Write(response)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		return
	}
}