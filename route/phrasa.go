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
	// Get query parameters for pagination, with default values if not provided
	pageStr := context.DefaultQuery("page", "1")
	limitStr := context.DefaultQuery("limit", "10")

	// Convert query parameters to integers
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		limit = 10
	}

	// Call the model function to fetch the data with pagination
	phrasa, totalPhrases, err := models.GetPhrasaWithPagination(page, limit)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch data"})
		return
	}

	// Calculate total pages
	totalPages := (totalPhrases + limit - 1) / limit

	// Return the structured response
	context.JSON(http.StatusOK, gin.H{
		"currentPage": page,
		"totalPage":   totalPages,
		"data":        phrasa,
	})
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
