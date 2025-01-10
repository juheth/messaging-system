package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/juheth/messaging-system/internal/usecases/room"
)

type RoomHandler struct {
	service *room.Service
}

func NewRoomHandler(service *room.Service) *RoomHandler {
	return &RoomHandler{service: service}
}

func (h *RoomHandler) CreateRoom(c *gin.Context) {
	var request struct {
		Name string `json:"name"`
	}
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	room, err := h.service.CreateRoom(request.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, room)
}
