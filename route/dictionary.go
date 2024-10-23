package route

import (
	"net/http"
	"strconv"

	"github.com/Izzat-Khudoyberganov/dictionary-app/models"
	"github.com/gin-gonic/gin"
)

func createDictionary(context *gin.Context) {
	var dictionary models.Dictionary

	err := context.ShouldBindJSON(&dictionary)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not create dictionary", "error": err})
		return
	}

	err = dictionary.SaveDictionary()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save dictionary"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Dictionary created."})

}

func getAllDictionary(context *gin.Context) {
	pageStr := context.DefaultQuery("page", "1")
	limitStr := context.DefaultQuery("limiit", "10")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		limit = 10
	}

	dictionary, totalDictionary, err := models.GetDictionary(page, limit)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch data"})
		return
	}

	// Calculate total pages
	totalPages := (totalDictionary + limit - 1) / limit

	// Return the structured response
	context.JSON(http.StatusOK, gin.H{
		"currentPage": page,
		"totalPage":   totalPages,
		"data":        dictionary,
	})
}

func updateDictionary(context *gin.Context) {
	dictionaryId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid dictionary id format"})
		return
	}

	var dictionary models.Dictionary

	err = context.ShouldBindJSON(&dictionary)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
	}

	dictionary.ID = dictionaryId

	err = dictionary.UpdateDictionary()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update dictionary", "error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Test updated successfully", "test": dictionary})
}

func getDictionaryById(context *gin.Context) {
	dictionaryId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse dictionary id"})
		return
	}

	test, err := models.GetDictionaryById(dictionaryId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch dictionary"})
		return
	}

	context.JSON(http.StatusOK, test)
}

func deleteDictionary(context *gin.Context) {
	dictionaryId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse dictionary id"})
		return
	}

	dictionary, err := models.GetDictionaryById(dictionaryId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch the dictionary"})
		return
	}

	err = dictionary.DeleteDictionary()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch the dictionary"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Dictionary successfully deleted"})

}
