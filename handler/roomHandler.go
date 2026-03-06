package handler

import (
	"net/http"

	"hotel-api/entity"
	"hotel-api/service/room"

	"github.com/labstack/echo/v4"
)

type RoomHandler struct {
	service room.Service
}

func NewRoomHandler(s room.Service) *RoomHandler {
	return &RoomHandler{s}
}

// @Summary Get all rooms
// @Description Retrieve a list of all available rooms
// @Tags rooms
// @Accept json
// @Produce json
// @Success 200 {array} entity.Room "List of rooms retrieved successfully"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /rooms [get]
func (h *RoomHandler) GetRooms(c echo.Context) error {

	rooms, err := h.service.GetRooms()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, rooms)
}

// @Summary Get room by ID
// @Description Retrieve details of a specific room using its ID
// @Tags rooms
// @Accept json
// @Produce json
// @Param id path string true "Room ID"
// @Success 200 {object} entity.Room "Room details retrieved successfully"
// @Failure 404 {object} map[string]string "Room not found"
// @Router /rooms/{id} [get]
func (h *RoomHandler) GetRoomByID(c echo.Context) error {

	id := c.Param("id")

	room, err := h.service.GetRoomByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "room not found",
		})
	}

	return c.JSON(http.StatusOK, room)
}

// @Summary Create a new room
// @Description Add a new room to the hotel system
// @Tags rooms
// @Accept json
// @Produce json
// @Param room body entity.Room true "Room details"
// @Success 201 {object} map[string]interface{} "Room created successfully"
// @Failure 400 {object} map[string]string "Invalid request body"
// @Router /rooms [post]
func (h *RoomHandler) CreateRoom(c echo.Context) error {

	var room entity.Room

	if err := c.Bind(&room); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid request body",
		})
	}

	err := h.service.Create(&room)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "room created successfully",
		"data":    room,
	})
}

// @Summary Update a room
// @Description Update details of an existing room
// @Tags rooms
// @Accept json
// @Produce json
// @Param id path string true "Room ID"
// @Param room body entity.Room true "Updated room details"
// @Success 200 {object} map[string]interface{} "Room updated successfully"
// @Failure 400 {object} map[string]string "Invalid request body"
// @Failure 500 {object} map[string]string "Failed to update room"
// @Router /rooms/{id} [put]
func (h *RoomHandler) UpdateRoom(c echo.Context) error {

	id := c.Param("id")

	var room entity.Room

	if err := c.Bind(&room); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid request body",
		})
	}

	err := h.service.Update(id, room)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	updatedRoom, err := h.service.GetRoomByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "failed to fetch updated room",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "room updated successfully",
		"data":    updatedRoom,
	})
}

// @Summary Delete a room
// @Description Remove a room from the hotel system
// @Tags rooms
// @Accept json
// @Produce json
// @Param id path string true "Room ID"
// @Success 200 {object} map[string]string "Room deleted successfully"
// @Failure 500 {object} map[string]string "Failed to delete room"
// @Router /rooms/{id} [delete]
func (h *RoomHandler) DeleteRoom(c echo.Context) error {

	id := c.Param("id")

	err := h.service.Delete(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "room deleted successfully",
	})
}
