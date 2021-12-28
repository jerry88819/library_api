package main

import (
	"library_api/database"
	"library_api/model"
	"library_api/routes"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {

	db, err := database.ConnectBun() // 先連接 Library Database
	if err != nil {
		fmt.Printf("database error =>%v\n", err)
		return
	} // if()
	fmt.Printf("database init success\n")

	defer db.Close()

	model.Init(db)

	// Start a new fiber app
	app := fiber.New()

	// Setup the router
	api := app.Group("/library", logger.New())

	// Setup the Node Routes
	routes.SetupRoutes(api)

	// Listen on PORT 3000
	app.Listen(":3001")

} // main()
