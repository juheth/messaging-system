package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/juheth/messaging-system/internal/domain"
	"github.com/juheth/messaging-system/internal/usecases/message"
)

type MessageHandler struct {
	messageService *message.Service
}

func NewMessageHandler(service *message.Service) *MessageHandler {
	return &MessageHandler{
		messageService: service,
	}
}

func (h *MessageHandler) CreateMessage(c *gin.Context) {
	var request struct {
		Content string `json:"content"`
		RoomID  int    `json:"room_id"`
	}

	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	user := domain.User{ID: userID.(int)}

	message, err := h.messageService.CreateMessage(request.Content, user, request.RoomID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, message)
}

func (h *MessageHandler) GetMessagesByRoom(c *gin.Context) {
	roomIDStr := c.Param("room_id")
	roomID, err := strconv.Atoi(roomIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid room ID"})
		return
	}

	messages, err := h.messageService.GetMessagesByRoom(roomID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, messages)
}
