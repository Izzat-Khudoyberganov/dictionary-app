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
	dictionary, err := models.GetAllDictionary()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not parse dictionary"})
		return
	}
	context.JSON(http.StatusOK, dictionary)
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

	test, err := models.GetAllDictionaryById(dictionaryId)

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

	dictionary, err := models.GetAllDictionaryById(dictionaryId)

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
