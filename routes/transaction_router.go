package routes

import (
	"dumbflix-api/handlers"
	"dumbflix-api/pkg/mysql"
	"dumbflix-api/repositories"

	"github.com/labstack/echo/v4"
)

func TransactionRoutes(e *echo.Group) {
	TransactionRepository := repositories.RepositoryTransaction(mysql.DB)

	h := handlers.HandlerTransaction(TransactionRepository)

	e.GET("/transactions", h.FindTransactions)
	e.GET("/transactions/:id", h.GetTransaction)
	e.POST("/transactions", h.AddTransaction)
	e.PATCH("/transactions/:id", h.EditTransaction)
	e.DELETE("/transactions/:id", h.DeleteTransaction)
}
