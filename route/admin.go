package route

import (
	"net/http"

	"github.com/Izzat-Khudoyberganov/dictionary-app/models"
	"github.com/Izzat-Khudoyberganov/dictionary-app/utils"
	"github.com/gin-gonic/gin"
)

func createAdmin(context *gin.Context) {
	var admin models.Admin

	err := context.ShouldBindJSON(&admin)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not create admin"})
		return
	}

	err = admin.SaveAdmin()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save admin"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Admin created."})
}
func loginAdmin(context *gin.Context) {
	var admin models.Admin

	err := context.ShouldBindJSON(&admin)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
	}

	err = admin.ValidateAdmin()

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	token, err := utils.GenerateTokenForAdmin(admin.Login, admin.ID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not autenticate admin"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Successfully logged in.", "token": token})

}
