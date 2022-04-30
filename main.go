package main

import (
	"github.com/gofiber/fiber"
	"github.com/maximof/go-fibre-crm-basic/database"
)

func setepRoutes(app *fiber.App)  {
	app.Get(GetLeads)
	app.Get(GetLead)
	app.Post(NewLead)
	app.Delete(DeleteLead)
}

func initDatabase()  {
	database
}

func main()  {
	app := fiber.New()

	setepRoutes(app)

	app.Listen(3001)

}