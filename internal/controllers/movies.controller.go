package controllers

import (
	"net/http"

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
