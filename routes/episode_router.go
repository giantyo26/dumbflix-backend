package routes

import (
	"dumbflix-api/handlers"
	"dumbflix-api/pkg/middleware"
	"dumbflix-api/pkg/mysql"
	"dumbflix-api/repositories"

	"github.com/labstack/echo/v4"
)

func EpisodeRoutes(e *echo.Group) {
	EpisodeRepository := repositories.RepositoryEpisode(mysql.DB)

	h := handlers.HandlerEpisode(EpisodeRepository)

	e.GET("/films/:filmID/episodes", h.FindEpisodesByFilm)
	e.GET("/films/:filmID/episodes/:id", h.GetEpisodeByFilm)
	e.POST("/episodes", middleware.UploadImage(h.AddEpisode))
	e.PATCH("/episodes/:id", middleware.UploadImage(h.EditEpisode))
	e.DELETE("/episodes/:id", h.DeleteEpisode)
}
