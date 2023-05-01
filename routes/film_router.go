package routes

import (
	"dumbflix-api/handlers"
	"dumbflix-api/pkg/middleware"
	"dumbflix-api/pkg/mysql"
	"dumbflix-api/repositories"

	"github.com/labstack/echo/v4"
)

func FilmRoutes(e *echo.Group) {
	FilmRepository := repositories.RepositoryFilm(mysql.DB)

	h := handlers.HandlerFilm(FilmRepository)

	e.GET("/films", h.FindFilms)
	e.GET("/films/:id", middleware.Auth(h.GetFilm))
	e.POST("/films", middleware.Auth(middleware.UploadImage(h.AddFilm)))
	e.PATCH("/films/:id", middleware.Auth(middleware.UploadImage(h.EditFilm)))
	e.DELETE("/films/:id", middleware.Auth(h.DeleteFilm))

}
