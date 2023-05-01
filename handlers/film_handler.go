package handlers

import (
	"context"
	filmdto "dumbflix-api/dto/film"
	dto "dumbflix-api/dto/result"
	"dumbflix-api/models"
	repository "dumbflix-api/repositories"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

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
		films[i].Film_Thumbnail = p.Film_Thumbnail
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: films})
}

func (h *handlerFilm) GetFilm(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	film, err := h.FilmRepository.GetFilm(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

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

	var ctx = context.Background()
	var CLOUD_NAME = os.Getenv("CLOUD_NAME")
	var API_KEY = os.Getenv("API_KEY")
	var API_SECRET = os.Getenv("API_SECRET")

	// Add your Cloudinary credentials ...
	cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)

	// Upload file to Cloudinary ...
	resp, err := cld.Upload.Upload(ctx, imageFile, uploader.UploadParams{Folder: "dumbflix"})
	if err != nil {
		fmt.Println(err.Error())
	}
	
	validation := validator.New()
	err = validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	film := models.Film{
		Title:          request.Title,
		Film_Thumbnail: resp.SecureURL,
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

	var ctx = context.Background()
	var CLOUD_NAME = os.Getenv("CLOUD_NAME")
	var API_KEY = os.Getenv("API_KEY")
	var API_SECRET = os.Getenv("API_SECRET")

	// Add your Cloudinary credentials ...
	cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)

	// Upload file to Cloudinary ...
	resp, err := cld.Upload.Upload(ctx, imageFile, uploader.UploadParams{Folder: "dumbflix"})

	if err != nil {
		fmt.Println(err.Error())
	}

	film, _ := h.FilmRepository.GetFilm(id)

	if request.Title != "" {
		film.Title = request.Title
	}
	if request.Film_Thumbnail != "" {
		film.Film_Thumbnail = resp.SecureURL
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

	editedFilm, err = h.FilmRepository.EditFilm(film)

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
