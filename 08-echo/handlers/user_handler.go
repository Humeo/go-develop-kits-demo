package handlers

import (
	"08-echo/models"
	"08-echo/services"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type UserHandler struct {
	userService *services.UserService
}

func NewUserHandler(us *services.UserService) *UserHandler {
	return &UserHandler{userService: us}
}

// GetUsers 获取所有用户
func (h *UserHandler) GetUsers(c echo.Context) error {
	users, err := h.userService.GetAll()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, users)
}

// GetUser 获取单个用户
func (h *UserHandler) GetUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := h.userService.GetByID(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "User not found")
	}
	return c.JSON(http.StatusOK, user)
}

// CreateUser 创建用户
func (h *UserHandler) CreateUser(c echo.Context) error {
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	createdUser, err := h.userService.Create(user)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, createdUser)
}
