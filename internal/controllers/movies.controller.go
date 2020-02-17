package controllers

import (
	"net/http"

	"github.com/Api-Marvel/internal/models"
	"github.com/Api-Marvel/internal/services"
	"github.com/gin-gonic/gin"
)

//GetIndex controller...
func GetIndex(c *gin.Context) {
	services.GetIndexString()
	c.JSON(http.StatusOK, "Get index correcto")
}

//GetComics ...
func GetComics(c *gin.Context) {
	DataComics := services.GetAllComics()
	c.JSON(http.StatusOK, DataComics)
}

//CreateComic ...
func CreateComic(c *gin.Context) {
	var comic models.ResultsComics
	c.BindJSON(&comic)
	result := services.CreateComic(&comic)
	if result != 1 {
		c.JSON(http.StatusInternalServerError, comic)
	}
	c.JSON(http.StatusOK, comic)
}
