package routes

import (
	"dumbflix-api/handlers"
	"dumbflix-api/pkg/mysql"
	"dumbflix-api/repositories"

	"github.com/labstack/echo/v4"
)

func CategoryRoutes(e *echo.Group) {
	categoryRepository := repositories.RepositoryCategory(mysql.DB)

	h := handlers.HandlerCategory(categoryRepository)

	e.GET("/categories", h.FindCategories)
	e.GET("/categories/:id", h.GetCategory)
	e.POST("/categories", h.AddCategory)
	e.PATCH("/categories/:id", h.EditCategory)
	e.DELETE("/categories/:id", h.DeleteCategory)
}
