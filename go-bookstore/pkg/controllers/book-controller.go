package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/M0HTeP/go/go-bookstore/pkg/models"
	"github.com/M0HTeP/go/go-bookstore/pkg/utils"
	"github.com/gorilla/mux"
)

var NewBook models.Book // переменная для создания новой книги по шаблону(заранее заданным графам в структуре)

// функция получения списка всех книг
func GetBook(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()
	res, _ := json.Marshal(newBooks) // декодим сообщение из json`a
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK) // записываем статус http (?)
	w.Write(res)                 // выводим результат
}

// функция получения книги с определенным ID
func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)                       // получаем данные из запроса
	bookId := vars["bookId"]                  // заносим в переменную bookId из нашего запроса
	ID, err := strconv.ParseInt(bookId, 0, 0) //  парсим ID
	if err != nil {
		fmt.Println("error while parsing") //  выводим в случае ошибки
	}
	bookDetails, _ := models.GetBookById(ID) //  получаем данные книги с определенным ID
	res, _ := json.Marshal(bookDetails)      // декодим сообщение из json`a
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK) // записываем статус http (?)
	w.Write(res)                 // выводим результат
}

// функция создания новой книги
func CreateBook(w http.ResponseWriter, r *http.Request) {
	CreateBook := &models.Book{}   // получаем информацию из структуры о том, какие нужны графы
	utils.ParseBody(r, CreateBook) // парсим тело запроса как задали в Utils
	b := CreateBook.CreateBook()   // загоняем в переменную созданную ранее функцию
	res, _ := json.Marshal(b)      // декодим сообщение из json`a
	w.WriteHeader(http.StatusOK)   // записываем статус http (?)
	w.Write(res)                   // выводим результат
}

// функция удаления книги по ID
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)                       // получаем данные из запроса
	bookId := vars["bookId"]                  // заносим в переменную bookId из нашего запроса
	ID, err := strconv.ParseInt(bookId, 0, 0) //  парсим ID
	if err != nil {
		fmt.Println("error while parsing") //  выводим в случае ошибки
	}
	book := models.DeleteBook(ID) //  удаляем книгу с определенным ID
	res, _ := json.Marshal(book)  // декодим сообщение из json`a
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK) // записываем статус http (?)
	w.Write(res)                 // выводим результат
}

// функция обновления информации о книге
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var updateBook = &models.Book{}           // получаем информацию из структуры о том, какие нужны графы
	utils.ParseBody(r, updateBook)            // парсим тело запроса как задали в Utils
	vars := mux.Vars(r)                       // получаем данные из запроса
	bookId := vars["bookId"]                  // заносим в переменную bookId из нашего запроса
	ID, err := strconv.ParseInt(bookId, 0, 0) //  парсим ID
	if err != nil {
		fmt.Println("error while parsing") //  выводим в случае ошибки
	}
	bookDetails, db := models.GetBookById(ID) // находим книгу с нужным ID
	if updateBook.Name != "" {                // если название книги не пустое
		bookDetails.Name = updateBook.Name // заполняем названием из запроса
	}
	if updateBook.Author != "" { // если автор книги отсутствует
		bookDetails.Author = updateBook.Author // заполняем автором из запроса
	}
	if updateBook.Publication != "" { // если издатель книги отсутствует
		bookDetails.Publication = updateBook.Publication // заполняем издателем из запроса
	}
	db.Save(&bookDetails)               // сохраняем изменения в БД
	res, _ := json.Marshal(bookDetails) // декодим сообщение из json`a
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK) // записываем статус http (?)
	w.Write(res)                 // выводим результат
}
