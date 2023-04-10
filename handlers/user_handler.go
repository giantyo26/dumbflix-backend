package handlers

import (
	dto "dumbflix-api/dto/result"
	userdto "dumbflix-api/dto/user"
	"dumbflix-api/models"
	"dumbflix-api/repositories"
	"net/http"
	"strconv"

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

	request := new(userdto.UserRequest)
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	user, err := h.UserRepository.GetUser(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	if request.Name != "" {
		user.Name = request.Name
	}
	if request.Email != "" {
		user.Email = request.Email
	}
	if request.Password != "" {
		user.Password = request.Password
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
	if request.Subscribe != "" {
		user.Subscribe = request.Subscribe
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
