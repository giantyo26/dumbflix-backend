package main

import (
	"fmt"
	"os"

	"dumbflix-api/database"
	"dumbflix-api/pkg/mysql"
	"dumbflix-api/routes"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load env file")
	}
	// var port = os.Getenv("PORT")
	var port = os.Getenv("PORT")

	e.Static("/uploads", "./uploads")

	mysql.ConnectToDatabase()
	database.RunMigration()
	routes.RouteInit(e.Group("/api/v1"))

	fmt.Println("server running localhost: " + port)
	e.Logger.Fatal(e.Start("localhost:" + port))
}
