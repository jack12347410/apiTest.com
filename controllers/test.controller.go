package controllers

import (
	"apiTest.com/initializers"
	"apiTest.com/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func CreateTest(ctx *gin.Context) {
	var input *models.CreateTestInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"error": err.Error()})
		return
	}

	test := models.Test{
		Title:    input.Title,
		Content:  input.Content,
		CreateAt: time.Now(),
		UpdateAt: time.Now(),
	}
	initializers.DB.Create(&test)

	ctx.JSON(http.StatusOK,
		gin.H{"data": test})
}

func FindTests(ctx *gin.Context) {
	var tests *[]models.Test
	var err error

	if id := ctx.Param("id"); id != "" {
		err = initializers.DB.Where("id = ?", id).Find(&tests).Error
	} else {
		err = initializers.DB.Find(&tests).Error
	}

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound,
			gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK,
		gin.H{"data": tests})
}

func UpdateTest(ctx *gin.Context) {
	var post *models.Test
	if err := initializers.DB.Where("id = ?", ctx.Param("id")).First(&post).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound,
			gin.H{"error": err.Error()})
		return
	}

	var update *models.UpdateTestInput
	if err := ctx.ShouldBindJSON(&update); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"error": err.Error()})
		return
	}

	updatePost := models.Test{
		Title:    update.Title,
		Content:  update.Content,
		UpdateAt: time.Now(),
	}
	initializers.DB.Model(&post).Updates(&updatePost)
	ctx.JSON(http.StatusOK,
		gin.H{"data": post})
}

func DeleteTest(ctx *gin.Context) {
	var post *models.Test
	if err := initializers.DB.Where("id = ?", ctx.Param("id")).First(&post).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound,
			gin.H{"error": err.Error()})
		return
	}

	initializers.DB.Delete(&post)
	ctx.JSON(http.StatusOK,
		gin.H{"data": post})
}
