package main

import (
	"fmt"
	"os"

	"dumbflix-api/database"
	"dumbflix-api/pkg/mysql"
	"dumbflix-api/routes"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PATCH, echo.DELETE},
		AllowHeaders: []string{"X-Requested-With", "Content-Type", "Authorization"},
	}))

	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load env file")
	}
	var port = os.Getenv("PORT")

	e.Static("/uploads", "./uploads")

	mysql.ConnectToDatabase()
	database.RunMigration()
	routes.RouteInit(e.Group("/api/v1"))

	fmt.Println("server running:" + port)
	e.Logger.Fatal(e.Start(":" + port))
}
