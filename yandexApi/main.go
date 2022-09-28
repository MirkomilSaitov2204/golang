package main

import (
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"yandexApi/utils"
)

const (
	apiPrefix string = "/api/v1"
)

var (
	port                    string
	bookResourcePrefix      string = apiPrefix + "/book"
	manyBooksResourcePrefix string = apiPrefix + "/books"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Could not find env file")
	}

	port = os.Getenv("app_port")
}

func main() {
	log.Println("Starting Rest api server", port)

	router := mux.NewRouter()
	utils.BuildBookResourse(router, bookResourcePrefix)
	utils.BuildManyBooksResourse(router, manyBooksResourcePrefix)

	log.Fatal(http.ListenAndServe(":"+port, router))
}
