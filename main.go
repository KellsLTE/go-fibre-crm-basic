package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	"github.com/maximof/go-fibre-crm-basic/database"
	"github.com/maximof/go-fibre-crm-basic/lead"
)

func setupRoutes(app *fiber.App) {
	app.Get("api/v1/leads", lead.GetLeads)
	app.Get("api/v1/lead/:id", lead.GetLead)
	app.Post("api/v1/lead", lead.NewLead)
	app.Delete("api/v1/lead/:id", lead.DeleteLead)
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "leads.db")
	if err != nil {
		panic("failed to connect to database")
	}
	fmt.Println("Connection opened to database")
	database.DBConn.AutoMigrate(&lead.Lead{})
	fmt.Println("Database Migrated")
}

func main() {
	app := fiber.New()

	initDatabase()

	setupRoutes(app)

	app.Listen(":3001")

	defer func(DBConn *gorm.DB) {
		err := DBConn.Close()
		if err != nil {

		}
	}(database.DBConn)
}
