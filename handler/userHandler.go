package handler

import (
	"net/http"

	"hotel-api/service/user"
	"hotel-api/util"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	service user.Service
}

func NewUserHandler(s user.Service) *UserHandler {
	return &UserHandler{s}
}

type RegisterRequest struct {
	Name     string `json:"name" example:"John Doe"`
	Email    string `json:"email" example:"john@example.com"`
	Password string `json:"password" example:"password123"`
	Phone    string `json:"phone" example:"+1234567890"`
}

type LoginRequest struct {
	Email    string `json:"email" example:"john@example.com"`
	Password string `json:"password" example:"password123"`
}

// @Summary Register a new user
// @Description Create a new user account with name, email, password and phone
// @Tags users
// @Accept json
// @Produce json
// @Param request body RegisterRequest true "User registration details"
// @Success 201 {object} map[string]interface{} "User registered successfully"
// @Failure 400 {object} map[string]string "Invalid request body or user already exists"
// @Router /users/register [post]
func (h *UserHandler) Register(c echo.Context) error {

	var req RegisterRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid request body",
		})
	}

	user, err := h.service.Register(
		req.Name,
		req.Email,
		req.Password,
		req.Phone,
	)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "Registered successfully",
		"user":    user,
	})
}

// @Summary User login
// @Description Authenticate user with email and password, returns JWT token
// @Tags users
// @Accept json
// @Produce json
// @Param request body LoginRequest true "Login credentials"
// @Success 200 {object} map[string]interface{} "Login successful with JWT token"
// @Failure 400 {object} map[string]string "Invalid request body"
// @Failure 401 {object} map[string]string "Invalid email or password"
// @Failure 500 {object} map[string]string "Failed to generate token"
// @Router /users/login [post]
func (h *UserHandler) Login(c echo.Context) error {

	var req LoginRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid request body",
		})
	}

	user, err := h.service.Login(req.Email, req.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"error": err.Error(),
		})
	}

	token, err := util.GenerateJWT(user.ID, user.Email, user.Role)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "failed to generate token",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "login successful",
		"token":   token,
	})
}
