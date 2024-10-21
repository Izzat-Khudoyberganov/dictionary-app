package main

import (
	"github.com/Izzat-Khudoyberganov/dictionary-app/db"
	"github.com/Izzat-Khudoyberganov/dictionary-app/route"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()
	route.RegisterRoutes(server)
	server.Run(":8000")
}
