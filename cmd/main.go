package main

import (
	"fmt"
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
	messageRepo := database.NewGormMessageRepository(db)
	roomRepo := database.NewRoomRepository(db)

	// Servicios
	messageService := message.NewService(messageRepo)
	roomService := room.NewService(roomRepo)

	// Controladores
	messageHandler := httpHandlers.NewMessageHandler(messageService)
	roomHandler := httpHandlers.NewRoomHandler(roomService)

	// Configurar Gin
	r := gin.Default()

	// Middleware de autenticación
	r.Use(auth.AuthMiddleware())

	// Rutas
	r.POST("/messages", messageHandler.CreateMessage)
	r.GET("/messages/:room_id", messageHandler.GetMessagesByRoom)
	r.POST("/rooms", roomHandler.CreateRoom) // Ruta para crear una sala

	token, err := auth.GenerateJWT(1)
	if err != nil {
		log.Fatalf("Error generando el token: %v", err)
	}
	fmt.Printf("Token generado: %s\n", token)
	r.Run(":8081")

}
