package handlers

import (
	filmdto "dumbflix-api/dto/film"
	dto "dumbflix-api/dto/result"
	"dumbflix-api/models"
	repository "dumbflix-api/repositories"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

var path_file = "http://localhost:5500/uploads/"

type handlerFilm struct {
	FilmRepository repository.FilmRepository
}

func HandlerFilm(FilmRepository repository.FilmRepository) *handlerFilm {
	return &handlerFilm{FilmRepository}
}

func (h *handlerFilm) FindFilms(c echo.Context) error {
	films, err := h.FilmRepository.FindFilms()
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	for i, p := range films {
		films[i].Film_Thumbnail = path_file + p.Film_Thumbnail
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: films})
}

func (h *handlerFilm) GetFilm(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	film, err := h.FilmRepository.GetFilm(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	film.Film_Thumbnail = path_file + film.Film_Thumbnail

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: film})
}

func (h *handlerFilm) AddFilm(c echo.Context) error {
	imageFile := c.Get("imageFile").(string)
	year, _ := strconv.Atoi(c.FormValue("year"))
	category_id, _ := strconv.Atoi(c.FormValue("category_id"))

	request := filmdto.FilmRequest{
		Title:          c.FormValue("title"),
		Film_Thumbnail: imageFile,
		Year:           year,
		CategoryID:     category_id,
		Description:    c.FormValue("description"),
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	film := models.Film{
		Title:          request.Title,
		Film_Thumbnail: request.Film_Thumbnail,
		Year:           request.Year,
		CategoryID:     request.CategoryID,
		Description:    request.Description,
	}

	film, err = h.FilmRepository.AddFilm(film)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	film, _ = h.FilmRepository.GetFilm(film.ID)
	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: convertResponseFilm(film)})
}

func (h *handlerFilm) EditFilm(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	imageFile := c.Get("imageFile").(string)
	year, _ := strconv.Atoi(c.FormValue("year"))
	category_id, _ := strconv.Atoi(c.FormValue("category_id"))

	request := filmdto.FilmRequest{
		Title:          c.FormValue("title"),
		Film_Thumbnail: imageFile,
		Year:           year,
		CategoryID:     category_id,
		Description:    c.FormValue("description"),
	}

	film, _ := h.FilmRepository.GetFilm(id)

	if request.Title != "" {
		film.Title = request.Title
	}
	if request.Film_Thumbnail != "" {
		film.Film_Thumbnail = request.Film_Thumbnail
	}
	if request.Year != 0 {
		film.Year = request.Year
	}
	if request.CategoryID != 0 {
		film.CategoryID = request.CategoryID
	}
	if request.Description != "" {
		film.Description = request.Description
	}

	editedFilm := models.Film{
		Title:          request.Title,
		Film_Thumbnail: request.Film_Thumbnail,
		Year:           request.Year,
		CategoryID:     request.CategoryID,
		Description:    request.Description,
	}
	editedFilm, err := h.FilmRepository.EditFilm(film)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: convertResponseFilm(editedFilm)})
}

func (h *handlerFilm) DeleteFilm(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param(("id")))

	film, err := h.FilmRepository.GetFilm(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}
	data, err := h.FilmRepository.DeleteFilm(film)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}
	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: data})
}

func convertResponseFilm(u models.Film) filmdto.FilmResponse {
	return filmdto.FilmResponse{
		Title:          u.Title,
		Film_Thumbnail: u.Film_Thumbnail,
		Year:           u.Year,
		CategoryID:     u.CategoryID,
		Category:       u.Category,
		Description:    u.Description,
	}
}
