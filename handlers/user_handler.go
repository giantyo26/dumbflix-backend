package handlers

import (
	"context"
	dto "dumbflix-api/dto/result"
	userdto "dumbflix-api/dto/user"
	"dumbflix-api/models"
	"dumbflix-api/repositories"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type handlerUser struct {
	UserRepository repositories.UserRepository
}

func HandlerUser(UserRepository repositories.UserRepository) *handlerUser {
	return &handlerUser{UserRepository}
}

func (h *handlerUser) FindUsers(c echo.Context) error {
	users, err := h.UserRepository.FindUsers()
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: users})
}

func (h *handlerUser) GetUser(c echo.Context) error {

	id, _ := strconv.Atoi(c.Param("id")) // Get the "id" parameter from the URL path, and store it in the "id" variable

	user, err := h.UserRepository.GetUser(id) // / Calling the "GetUser" method on the "UserRepository" field of "handler" struct to get a user with the provided "id
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	// Return a JSON response with a HTTP status code of 200 (OK) and the user information converted to a custom response struct in the response body
	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: user})
}

func (h *handlerUser) EditUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	imageFile := c.Get("imageFile").(string)

	request := userdto.UserRequest{
		Name:          c.FormValue("fullname"),
		AvatarProfile: imageFile,
		Email:         c.FormValue("email"),
		Gender:        c.FormValue("gender"),
		Phone:         c.FormValue("phone"),
		Address:       c.FormValue("address"),
		Subscribe:     c.FormValue("subscribe"),
	}
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
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

	user, err := h.UserRepository.GetUser(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}
	if request.Name != "" {
		user.Name = request.Name
	}
	if request.AvatarProfile != "" {
		user.AvatarProfile = resp.SecureURL
	}
	if request.Email != "" {
		user.Email = request.Email
	}
	if request.Gender != "" {
		user.Gender = request.Gender
	}
	if request.Name != "" {
		user.Name = request.Name
	}
	if request.Phone != "" {
		user.Phone = request.Phone
	}
	if request.Address != "" {
		user.Address = request.Address
	}

	updatedUser, err := h.UserRepository.EditUser(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: updatedUser})
}

func (h *handlerUser) DeleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	user, err := h.UserRepository.GetUser(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	data, err := h.UserRepository.DeleteUser(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: deleteResponse(data)})

}

func deleteResponse(u models.User) userdto.DeleteUserResponse {
	return userdto.DeleteUserResponse{
		ID: u.ID,
	}
}
