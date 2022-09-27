package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

var (
	port string = "8080"

	db []Pizza
)

type Pizza struct {
	Id       int     `json:"id"`
	Diameter int     `json:"diameter"`
	Price    float64 `json:"price"`
	Title    string  `json:"title"`
}

func init() {
	pizza1 := Pizza{
		Id:       1,
		Diameter: 22,
		Price:    800.80,
		Title:    "Peppironi",
	}

	pizza2 := Pizza{
		Id:       2,
		Diameter: 22,
		Price:    800.80,
		Title:    "Peppironi2",
	}

	pizza3 := Pizza{
		Id:       3,
		Diameter: 22,
		Price:    800.80,
		Title:    "Peppironi3",
	}

	db = append(db, pizza1, pizza2, pizza3)
}

func FindPizzaById(id int) (Pizza, bool) {
	var pizza Pizza
	var found bool
	for _, p := range db {
		if p.Id == id {
			pizza = p
			found = true
			break
		}
	}
	return pizza, found
}

func GetAllPizzas(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK) //statys code request
	json.NewEncoder(writer).Encode(db)

}

func GetPizzaById(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Fatal(err)
	}
	pizza, ok := FindPizzaById(id)
	if ok {
		writer.WriteHeader(http.StatusOK)
		json.NewEncoder(writer).Encode(pizza)
	} else {
		log.Fatal("Error")
	}
}

func main() {
	log.Println("Trying to start REST api pizza")
	router := mux.NewRouter()

	router.HandleFunc("/pizzas", GetAllPizzas).Methods("GET")
	router.HandleFunc("/pizza/{id}", GetPizzaById).Methods("GET")

	log.Println("Router configurated successfully")

	http.ListenAndServe(":"+port, router)

	//if err := http.ListenAndServe(":"+port, nil); err != nil {
	//	log.Fatal(err)
	//}
}
