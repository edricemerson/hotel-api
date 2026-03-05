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

func (h *RoomHandler) GetRooms(c echo.Context) error {

	rooms, err := h.service.GetRooms()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, rooms)
}

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

func (h *RoomHandler) CreateRoom(c echo.Context) error {

	var room entity.Room

	if err := c.Bind(&room); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid request body",
		})
	}

	err := h.service.Create(room)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "room created successfully",
		"room":    room,
	})
}

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

	return c.JSON(http.StatusOK, map[string]string{
		"message": "room updated successfully",
	})
}

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
