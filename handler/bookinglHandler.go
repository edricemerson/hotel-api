package handler

import (
	"net/http"

	"hotel-api/entity"
	"hotel-api/service/booking"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type BookingHandler struct {
	service booking.Service
}

func NewBookingHandler(s booking.Service) *BookingHandler {
	return &BookingHandler{s}
}

func (h *BookingHandler) CreateBooking(c echo.Context) error {

	var booking entity.Booking

	if err := c.Bind(&booking); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid request body",
		})
	}

	user := c.Get("user")
	if user == nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"error": "unauthorized",
		})
	}

	claims, ok := user.(jwt.MapClaims)
	if !ok {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"error": "invalid token claims",
		})
	}

	idValue, ok := claims["user_id"] // FIXED HERE
	if !ok {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"error": "user id missing in token",
		})
	}

	userID := int(idValue.(float64))

	booking.UserID = userID

	err := h.service.CreateBooking(&booking)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "booking created successfully",
		"data":    booking,
	})
}

func (h *BookingHandler) GetMyBookings(c echo.Context) error {

	claims := c.Get("user").(jwt.MapClaims)
	userID := int(claims["id"].(float64))

	bookings, err := h.service.GetMyBookings(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": bookings,
	})
}

func (h *BookingHandler) GetBookingByID(c echo.Context) error {

	id := c.Param("id")

	booking, err := h.service.GetBookingByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "booking not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": booking,
	})
}

func (h *BookingHandler) UpdateBooking(c echo.Context) error {

	id := c.Param("id")

	var booking entity.Booking

	if err := c.Bind(&booking); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid request body",
		})
	}

	err := h.service.UpdateBooking(id, &booking)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "booking updated successfully",
	})
}

func (h *BookingHandler) DeleteBooking(c echo.Context) error {

	id := c.Param("id")

	err := h.service.DeleteBooking(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "booking deleted successfully",
	})
}
