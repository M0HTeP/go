package main

import (
	"fmt"
	"github.com/M0HTeP/go-fiber-crm-basic/database"
	"github.com/M0HTeP/go-fiber-crm-basic/lead"
	"github.com/gofiber/fiber"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/lead", lead.GetLeads)
	app.Get("/api/v1/lead:id", lead.GetLead)
	app.Post("/api/v1/lead", lead.NewLead)
	app.Delete("/api/v1/lead:id", lead.DeleteLead)
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "leads.db") // открываем БД leads.db
	if err != nil {
		panic("faild to connect database")
	}
	fmt.Println("Connection opened to database")
	database.DBConn.AutiMigrate(&lead.Lead{})
	ftm.Println("Database Migrated")
}

func main() {
	app := fiber.New()
	initDatabase()
	setupRoutes(app)
	app.Listen(3000)
	defer database.DBConn.Close() // после выполнения функции закрыть БД

}
