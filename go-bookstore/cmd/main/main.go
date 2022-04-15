package main

import (
	"log"
	"net/http"

	"github.com/M0HTeP/go/go-bookstore/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// главная функция
func main() {
	r := mux.NewRouter()                                // создаем новый инстанс
	routes.RegisterBookStoreRoutes(r)                   // Для нашего инстанса cоздаем необходимые нам хендлеры(не хватает вокабуляра для того чтобы написать это слово на русский манер) для функций
	http.Handle("/", r)                                 // регистрируем хендлер (?)
	log.Fatal(http.ListenAndServe("localhost:9010", r)) // в слечае ошибке на назначенном порте - заносим ее в лог
}
