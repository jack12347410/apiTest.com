package controllers

import (
	"apiTest.com/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type UserController struct {
	DB *gorm.DB
}

func CreateUser(ctx *gin.Context) {
	var create *models.CreateUserInput
	if err := ctx.ShouldBindJSON(&create); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

}
