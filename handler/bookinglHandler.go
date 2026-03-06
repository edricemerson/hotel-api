package handler

import (
	"net/http"
	"strconv"
	"time"

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

// @Summary Create a new booking
// @Description Book a room for specified check-in and check-out dates (requires authentication)
// @Tags bookings
// @Accept json
// @Produce json
// @Security Bearer
// @Param request body entity.BookingRequest true "Booking details with RoomID, CheckIn (YYYY-MM-DD), CheckOut (YYYY-MM-DD)"
// @Success 201 {object} map[string]interface{} "Booking created successfully"
// @Failure 400 {object} map[string]string "Invalid request body or date format"
// @Failure 401 {object} map[string]string "Unauthorized - invalid or missing token"
// @Router /bookings [post]
func (h *BookingHandler) CreateBooking(c echo.Context) error {

	var req entity.BookingRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid request body",
		})
	}

	checkIn, err := time.Parse("2006-01-02", req.CheckIn)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid check_in format (use YYYY-MM-DD)",
		})
	}

	checkOut, err := time.Parse("2006-01-02", req.CheckOut)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid check_out format (use YYYY-MM-DD)",
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

	idValue, ok := claims["user_id"]
	if !ok {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"error": "user id missing in token",
		})
	}

	userID := int(idValue.(float64))

	booking := entity.Booking{
		UserID:   userID,
		RoomID:   req.RoomID,
		CheckIn:  checkIn,
		CheckOut: checkOut,
	}

	err = h.service.CreateBooking(&booking)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	bookingWithRoom, err := h.service.GetBookingByID(strconv.Itoa(booking.ID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "failed to load booking data",
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "booking created successfully",
		"data":    bookingWithRoom,
	})
}

// @Summary Get all user bookings
// @Description Retrieve all bookings for the authenticated user
// @Tags bookings
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} map[string]interface{} "List of user bookings"
// @Failure 401 {object} map[string]string "Unauthorized - invalid or missing token"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /bookings [get]
func (h *BookingHandler) GetMyBookings(c echo.Context) error {

	claims := c.Get("user").(jwt.MapClaims)

	idValue, ok := claims["user_id"]
	if !ok {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"error": "user id missing in token",
		})
	}

	userIDFloat, ok := idValue.(float64)
	if !ok {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"error": "invalid user id format",
		})
	}

	userID := int(userIDFloat)

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

// @Summary Get booking by ID
// @Description Retrieve details of a specific booking
// @Tags bookings
// @Accept json
// @Produce json
// @Param id path string true "Booking ID"
// @Success 200 {object} map[string]interface{} "Booking details retrieved successfully"
// @Failure 404 {object} map[string]string "Booking not found"
// @Router /bookings/{id} [get]
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

// @Summary Update a booking
// @Description Update booking details (check-in, check-out dates)
// @Tags bookings
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "Booking ID"
// @Param request body entity.Booking true "Updated booking details"
// @Success 200 {object} map[string]string "Booking updated successfully"
// @Failure 400 {object} map[string]string "Invalid request body"
// @Failure 401 {object} map[string]string "Unauthorized - invalid or missing token"
// @Router /bookings/{id} [put]
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

// @Summary Delete a booking
// @Description Cancel an existing booking
// @Tags bookings
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "Booking ID"
// @Success 200 {object} map[string]string "Booking deleted successfully"
// @Failure 400 {object} map[string]string "Failed to delete booking"
// @Failure 401 {object} map[string]string "Unauthorized - invalid or missing token"
// @Router /bookings/{id} [delete]
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
