package route

import (
	"github.com/Izzat-Khudoyberganov/dictionary-app/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {

	// Protected routes
	authenticateAdmin := server.Group("/")
	authenticateAdmin.Use(middleware.AuthenticateAdmin)
	authenticateAdmin.POST("/dictionary", createDictionary)
	authenticateAdmin.DELETE("/dictionary/:id", deleteDictionary)
	authenticateAdmin.PATCH("/dictionary/:id", updateDictionary)

	server.GET("/dictionary", getAllDictionary)
	server.GET("/dictionary/:id", getDictionaryById)

	server.POST("/admin", createAdmin)
	server.POST("/login-admin", loginAdmin)

	authenticateAdmin.POST("/phrasal-verbs", createPhrasa)
	authenticateAdmin.PATCH("/phrasal-verbs/:id", updatePhrasa)
	authenticateAdmin.DELETE("/phrasal-verbs/:id", deletePhrasa)

	server.GET("/phrasal-verbs", getAllPhrasa)
	server.GET("/phrasal-verbs/:id", getPhrasaById)
}
