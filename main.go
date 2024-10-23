package main

import (
	"log"
	"os"
	"time"

	"github.com/Izzat-Khudoyberganov/dictionary-app/db"
	"github.com/Izzat-Khudoyberganov/dictionary-app/route"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()
	server.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	route.RegisterRoutes(server)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	err := server.Run(":" + port)
	if err != nil {
		log.Panicf("error: %s", err)
	}
}
