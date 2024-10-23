package route

import (
	"net/http"
	"strconv"

	"github.com/Izzat-Khudoyberganov/dictionary-app/models"
	"github.com/gin-gonic/gin"
)

func createPhrasa(context *gin.Context) {
	var phrasa models.Phrasa

	err := context.ShouldBindJSON(&phrasa)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not create phrasa", "error": err})
		return
	}

	err = phrasa.SavePhrasa()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save phrasa"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Phrasa created."})

}

func getAllPhrasa(context *gin.Context) {
	phrasa, err := models.GetAllPhrasa()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not parse phrasa"})
		return
	}
	context.JSON(http.StatusOK, phrasa)
}

func updatePhrasa(context *gin.Context) {
	phrasaId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid phrasa id format"})
		return
	}

	var phrasa models.Phrasa

	err = context.ShouldBindJSON(&phrasa)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
	}

	phrasa.ID = phrasaId

	err = phrasa.UpdatePhrasa()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update phrasa", "error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Test updated successfully", "test": phrasa})
}

func getPhrasaById(context *gin.Context) {
	phrasaId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse phrasa id"})
		return
	}

	phrasa, err := models.GetPhrasaById(phrasaId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch phrasa"})
		return
	}

	context.JSON(http.StatusOK, phrasa)
}

func deletePhrasa(context *gin.Context) {
	phrasaId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse phrasa id"})
		return
	}

	phrasa, err := models.GetPhrasaById(phrasaId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch the phrasa"})
		return
	}

	err = phrasa.DeletePhrasa()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch the phrasa"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Phrasa successfully deleted"})

}
