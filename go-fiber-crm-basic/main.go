package main

import (
	"fmt"
	"github.com/M0HTeP/go-fiber-crm-basic/database"
	"github.com/M0HTeP/go-fiber-crm-basic/lead"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/lead", lead.GetLeads)
	app.Get("/api/v1/lead/:id", lead.GetLead)
	app.Post("/api/v1/lead", lead.NewLead)
	app.Delete("/api/v1/lead/:id", lead.DeleteLead)
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "leads.db") // открываем БД leads.db
	if err != nil {
		panic("faild to connect database")	//    если подключение не удалось говорим об ошибке
	}
	fmt.Println("Connection opened to database")	//    если подключились к БД - сообщаем об этом
	database.DBConn.AutoMigrate(&lead.Lead{})	//    автомиграция БД (добавление недостающих полей для заполнения БД, например: дата создания поля)
	fmt.Println("Database Migrated")	//    говорим что миграция успешно завершена
}

func main() {
	app := fiber.New()	//    создаем новый инстанс
	initDatabase()	//    инициализируем БД
	setupRoutes(app)	//    настраиваем рауты (ссылки)
	app.Listen(3000)	//    прослушиваем порт(на этом порту открываем возможность работы с БД)
	defer database.DBConn.Close() // после выполнения функции закрыть БД

}
