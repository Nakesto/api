package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/join-api/model"
)

func UserLogin(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var user []model.User
	result := db.Find(&user)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "error", "message": result.Error})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "success", "data": user})
		return
	}
}
