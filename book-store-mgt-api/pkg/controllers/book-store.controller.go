package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/dkrest1/book-store-api/pkg/models"
	"github.com/dkrest1/book-store-api/pkg/utils"
	"github.com/gorilla/mux"
)



var  NewBook models.Book

func GetBooks (w http.ResponseWriter, r * http.Request) {
	NewBook := models.GetAllBooks()
	res, _ := json.Marshal(NewBook)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBook(w http.ResponseWriter, r * http.Request) {
	vars  := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 10, 64)
	if err != nil {
		http.Error(w, "Error while parsing bookId", http.StatusBadRequest)
        return
	}
	bookDetails, _ := models.GetBookById(ID)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "appplication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	CreateBook := &models.Book{}
	utils.ParseBody(r, CreateBook)
	b := CreateBook.CreateBook()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	println(res)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r * http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 10, 64)
	if err != nil {
		http.Error(w, "Error while parsing bookId", http.StatusBadRequest)
        return
	}
	book := models.DeleteBook(ID)
	res, _ := json.Marshal((book))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var updateBook = &models.Book{}
	utils.ParseBody(r, updateBook)
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 10, 64)
	if err != nil {
		http.Error(w, "Error while parsing bookId", http.StatusBadRequest)
        return
	}
	bookDetails, db := models.GetBookById(ID) 
	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}
	db.Save(&bookDetails)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}