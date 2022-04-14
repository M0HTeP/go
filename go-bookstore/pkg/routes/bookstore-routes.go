package routes

import (
	"github.com/M0HTeP/go-bookstore/pkg/controllers"
	"github.com/gorilla/mux"
)

// Создаем необходимые нам хендлеры(не хватает вокабуляра для того чтобы написать это слово на русский манер) для функций
var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")           // функция создания книги
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")               // функция получения информации о книгах
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")   // функция получения информации о книге по ее Id
	router.HandleFunc("/book{bookId}/", controllers.UpdateBook).Methods("PUT")    // функция обновления информации о книге под определенным ID
	router.HandleFunc("/book{bookId}/", controllers.DeleteBook).Methods("DELETE") // функция удаления книги
}
