package models

import (
	"github.com/M0HTeP/go/go-bookstore/pkg/config"
	"github.com/jinzhu/gorm"
)

// создаем переменную для хранения БД
var db *gorm.DB

// создаем структуру хранения данных о книгах
type Book struct {
	gorm.Model
	Name        string `gorm:""json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

// функция начала работы с БД
func init() {
	config.Connect()        // соединяемся с БД
	db = config.GetDB()     // используем БД из конфига
	db.AutoMigrate(&Book{}) // функция добавления недостающих полей (?)
}

// функция создания новой книги
func (b *Book) CreateBook() *Book { // передаем указатель на информацию о книге
	db.NewRecord(b) // проверяем наличие инфы в БД со значением b
	db.Create(&b)   // делаем новую запись в БД со значением b
	return b
}

// функция получения списка всех книг
func GetAllBooks() []Book {
	var Books []Book // создаем переменную с пустым списком
	db.Find(&Books)  // ищем указатели на книги
	return Books     // возвращаем полученную информацию
}

// функция получения информации о книге с заданным ID
func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book                          // создаем переменную
	db := db.Where("ID=?", Id).Find(&getBook) // в БД в графе ID ищем значение переменной getBook(интовое значение ID)
	return &getBook, db                       // возвращаем полученную информацию
}

// функция удаления книги
func DeleteBook(ID int64) Book {
	var book Book                     // создаем переменную
	db.Where("ID=?", ID).Delete(book) // в БД в графе ID ищем значение переменной Book(интовое значение ID) и удаляем эту книгу
	return book                       // возвращаем переменную (?)
}
