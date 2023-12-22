package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zakihaha/gin-funding/helper"
	"github.com/zakihaha/gin-funding/user"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Register failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newUser, err := h.userService.RegisterUser(input)
	if err != nil {
		response := helper.APIResponse("Register failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// jwt

	formatter := user.FormatUser(newUser, "token")

	response := helper.APIResponse("Account has been registered", http.StatusCreated, "success", formatter)

	c.JSON(http.StatusCreated, response)
}