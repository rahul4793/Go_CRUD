package controller

import (
	"net/http"

	"crud-app/manager"
	"crud-app/request"
	"crud-app/response"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type UserController struct {
	manager   *manager.UserManager
	validator *validator.Validate
}

func NewUserController(m *manager.UserManager, v *validator.Validate) *UserController {
	return &UserController{
		manager:   m,
		validator: v,
	}
}

func (c *UserController) CreateUser(ctx echo.Context) error {
	var req request.CreateUserRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.UserResponse{Message: "Invalid request"})
	}

	if err := c.validator.Struct(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.UserResponse{Message: err.Error()})
	}

	user, err := c.manager.CreateUser(ctx.Request().Context(), req)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.UserResponse{Message: err.Error()})
	}

	return ctx.JSON(http.StatusCreated, response.UserResponse{Message: "User created", Data: user})
}

func (c *UserController) GetAllUsers(ctx echo.Context) error {
	var req request.PaginationRequest

	if err := ctx.Bind(&req); err != nil {
		req.Page = 1
		req.Limit = 10
	}

	// users, err := c.manager.GetAllUsers(ctx.Request().Context(), req.Page, req.Limit)
	users, err := c.manager.GetAllUsers(req.Page, req.Limit)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.UserResponse{Message: err.Error()})
	}

	return ctx.JSON(http.StatusOK, response.UserResponse{
		Message: "Users fetched",
		Data: map[string]interface{}{
			"page":  req.Page,
			"limit": req.Limit,
			"users": users,
		},
	})
}

func (c *UserController) GetUserByID(ctx echo.Context) error {
	id := ctx.Param("id")
	user, err := c.manager.GetUserByID(ctx.Request().Context(), id)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, response.UserResponse{Message: err.Error()})
	}

	return ctx.JSON(http.StatusOK, response.UserResponse{Message: "User fetched", Data: user})
}

func (c *UserController) UpdateUser(ctx echo.Context) error {
	id := ctx.Param("id")
	var req request.UpdateUserRequest

	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.UserResponse{Message: "Invalid request"})
	}

	if err := c.validator.Struct(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.UserResponse{Message: "Bad request"})
	}

	_, err := c.manager.GetUserByID(ctx.Request().Context(), id)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, response.UserResponse{Message: "User not found or is deleted"})
	}

	if err := c.manager.UpdateUser(ctx.Request().Context(), id, req); err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.UserResponse{Message: err.Error()})
	}

	return ctx.JSON(http.StatusOK, response.UserResponse{Message: "User updated"})
}

func (c *UserController) DeleteUser(ctx echo.Context) error {
	id := ctx.Param("id")

	_, err := c.manager.GetUserByID(ctx.Request().Context(), id)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, response.UserResponse{Message: "User already deleted"})
	}

	if err := c.manager.DeleteUser(ctx.Request().Context(), id); err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.UserResponse{Message: err.Error()})
	}

	return ctx.JSON(http.StatusOK, response.UserResponse{Message: "User soft-deleted"})
}
