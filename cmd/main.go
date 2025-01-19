package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/juheth/messaging-system/internal/auth"
	"github.com/juheth/messaging-system/internal/infrastructure/database"
	httpHandlers "github.com/juheth/messaging-system/internal/interfaces/http"
	"github.com/juheth/messaging-system/internal/usecases/message"
	"github.com/juheth/messaging-system/internal/usecases/room"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "user:password@tcp(localhost:3306)/chat_db?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	// Repositorios
	messageRepo := database.NewMessageRepository(db)
	roomRepo := database.NewRoomRepository(db)

	// Servicios
	messageService := message.NewService(messageRepo)
	roomService := room.NewService(roomRepo)

	// Controladores
	messageHandler := httpHandlers.NewMessageHandler(messageService)
	roomHandler := httpHandlers.NewRoomHandler(roomService)

	// Configurar Gin
	r := gin.Default()

	// Ruta para generar el token
	r.POST("/token", auth.GenerateTokenHandler)

	// Middleware de autenticación
	r.Use(auth.AuthMiddleware())

	// Rutas de mensajes
	r.POST("/messages", messageHandler.CreateMessage)
	r.GET("/messages/:room_id", messageHandler.GetMessagesByRoom)
	r.PUT("/messages/:id", messageHandler.UpdateMessage)
	r.DELETE("/messages/:id", messageHandler.DeleteMessage)

	// Rutas de salas
	r.POST("/rooms", roomHandler.CreateRoom)
	r.PUT("/rooms/:id", roomHandler.UpdateRoom)
	r.DELETE("/rooms/:id", roomHandler.DeleteRoom)

	if err := r.Run(":8081"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
