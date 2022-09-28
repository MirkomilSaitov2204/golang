package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"yandexApi/models"
)

func GetBookById(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	id, err := strconv.Atoi(mux.Vars(request)["id"])
	if err != nil {
		log.Println("Error wthile parsing happend", err)
		writer.WriteHeader(http.StatusBadRequest)
		msg := models.Message{Message: "Do not use parametr id as uncasted to int type"}
		json.NewEncoder(writer).Encode(msg)
	}
	book, ok := models.FindBookById(id)
	log.Println("get book with id: ", id)
	if !ok {
		writer.WriteHeader(http.StatusNotFound)
		msg := models.Message{Message: "message not found"}
		json.NewEncoder(writer).Encode(msg)
	} else {
		writer.WriteHeader(http.StatusOK)
		json.NewEncoder(writer).Encode(book)
	}
}

func CreateBook(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	log.Println("createing new books ...")
	var book models.Book

	err := json.NewDecoder(request.Body).Decode(&book)

	if err != nil {
		msg := models.Message{Message: "provided json file is invalid"}
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	var newBookId int = len(models.DB) + 1
	book.Id = newBookId
	models.DB = append(models.DB, book)
	writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer).Encode(book)

}
func UpdateBookById(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	log.Println("updating new books ...")

	id, err := strconv.Atoi(mux.Vars(request)["id"])
	if err != nil {
		log.Println("Error wthile parsing happend", err)
		writer.WriteHeader(http.StatusBadRequest)
		msg := models.Message{Message: "Do not use parametr id as uncasted to int type"}
		json.NewEncoder(writer).Encode(msg)
	}

	oldBook, ok := models.FindBookById(id)
	var newBook models.Book

	if !ok {
		writer.WriteHeader(http.StatusNotFound)
		msg := models.Message{Message: "message not found"}
		json.NewEncoder(writer).Encode(msg)
	}
	err = json.NewDecoder(request.Body).Decode(&newBook)
	if err != nil {
		msg := models.Message{Message: "provided json file is invalid"}
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(msg)
		return
	}

}
func DeleteBookById(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	initHeaders(writer)
	log.Println("deleting new books ...")

	id, err := strconv.Atoi(mux.Vars(request)["id"])
	if err != nil {
		log.Println("Error wthile parsing happend", err)
		writer.WriteHeader(http.StatusBadRequest)
		msg := models.Message{Message: "Do not use parametr id as uncasted to int type"}
		json.NewEncoder(writer).Encode(msg)
	}

	book, ok := models.FindBookById(id)
	if !ok {
		writer.WriteHeader(http.StatusNotFound)
		msg := models.Message{Message: "message not found"}
		json.NewEncoder(writer).Encode(msg)
	}
	msg := models.Message{Message: "message successfully deleted"}
	json.NewEncoder(writer).Encode(msg)

}
