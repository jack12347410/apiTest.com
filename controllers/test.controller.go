package controllers

import (
	"apiTest.com/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type testController struct {
	testService models.TestService
}

func (t *testController) CreateTest(ctx *gin.Context) {
	var input *models.CreateTestInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"error": err.Error()})
		return
	}

	result, err := t.testService.Create(input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest,
			gin.H{"error": err.Error()})
	}

	ctx.JSON(http.StatusOK,
		gin.H{"data": result})
}

func (t *testController) FindOneTest(ctx *gin.Context) {
	test, err := t.testService.FindById(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound,
			gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK,
		gin.H{"data": test})
}

func (t *testController) FindTests(ctx *gin.Context) {
	tests, err := t.testService.FindAll()

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound,
			gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK,
		gin.H{"data": tests})
}

func (t *testController) UpdateTest(ctx *gin.Context) {
	var update *models.UpdateTestInput
	if err := ctx.ShouldBindJSON(&update); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"error": err.Error()})
		return
	}

	post, err := t.testService.Update(ctx.Param("id"), update)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound,
			gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK,
		gin.H{"data": post})
}

func (t *testController) DeleteTest(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func NewTestController(s models.TestService) models.TestController {
	return &testController{
		testService: s,
	}
}

//func CreateTest(ctx *gin.Context) {
//	var input *models.CreateTestInput
//	if err := ctx.ShouldBindJSON(&input); err != nil {
//		ctx.AbortWithStatusJSON(http.StatusBadRequest,
//			gin.H{"error": err.Error()})
//		return
//	}
//
//	test := models.Test{
//		Title:    input.Title,
//		Content:  input.Content,
//		CreateAt: time.Now(),
//		UpdateAt: time.Now(),
//	}
//	initializers.DB.Create(&test)
//
//	ctx.JSON(http.StatusOK,
//		gin.H{"data": test})
//}
//
//func FindTests(ctx *gin.Context) {
//	var tests *[]models.Test
//	var err error
//
//	if id := ctx.Param("id"); id != "" {
//		err = initializers.DB.Where("id = ?", id).Find(&tests).Error
//	} else {
//		err = initializers.DB.Find(&tests).Error
//	}
//
//	if err != nil {
//		ctx.AbortWithStatusJSON(http.StatusNotFound,
//			gin.H{"error": err.Error()})
//		return
//	}
//
//	ctx.JSON(http.StatusOK,
//		gin.H{"data": tests})
//}
//
//func UpdateTest(ctx *gin.Context) {
//	var post *models.Test
//	if err := initializers.DB.Where("id = ?", ctx.Param("id")).First(&post).Error; err != nil {
//		ctx.AbortWithStatusJSON(http.StatusNotFound,
//			gin.H{"error": err.Error()})
//		return
//	}
//
//	var update *models.UpdateTestInput
//	if err := ctx.ShouldBindJSON(&update); err != nil {
//		ctx.AbortWithStatusJSON(http.StatusBadRequest,
//			gin.H{"error": err.Error()})
//		return
//	}
//
//	updatePost := models.Test{
//		Title:    update.Title,
//		Content:  update.Content,
//		UpdateAt: time.Now(),
//	}
//	initializers.DB.Model(&post).Updates(&updatePost)
//	ctx.JSON(http.StatusOK,
//		gin.H{"data": post})
//}
//
//func DeleteTest(ctx *gin.Context) {
//	var post *models.Test
//	if err := initializers.DB.Where("id = ?", ctx.Param("id")).First(&post).Error; err != nil {
//		ctx.AbortWithStatusJSON(http.StatusNotFound,
//			gin.H{"error": err.Error()})
//		return
//	}
//
//	initializers.DB.Delete(&post)
//	ctx.JSON(http.StatusOK,
//		gin.H{"data": post})
//}
