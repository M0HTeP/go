package main

import (
	"fmt"
	"github.com/go-fiber-crm-basic/database"
	"github.com/gofiber/fiber"
)

func setupRoutes(app *fiber.App) {
	app.Get(GetLeads)
	app.Get(GetLead)
	app.Post(NewLead)
	app.Delete(Delete)
}

func initDatabase() {

}

func main() {
	app := fiber.New()
	setupRoutes(app)
	app.Listen(3000)

}
