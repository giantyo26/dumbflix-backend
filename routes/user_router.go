package routes

import (
	"dumbflix-api/handlers"
	"dumbflix-api/pkg/mysql"
	"dumbflix-api/repositories"

	"github.com/labstack/echo/v4"
)

func UserRoutes(e *echo.Group) {
	userRepository := repositories.RepositoryUser(mysql.DB)
	h := handlers.HandlerUser(userRepository)

	e.GET("/users", h.FindUsers)
	e.GET("/users/:id", h.GetUser)
	e.PATCH("/users/:id", h.EditUser)
	e.DELETE("/users/:id", h.DeleteUser)
}
