package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PostController struct {
	DB *gorm.DB
}

func NewPostController(DB *gorm.DB) PostController {
	return PostController{DB}
}

func (pc *PostController) CreatePost(ctx *gin.Context) {
	//currentUser := ctx.MustGet("currentUser").(models.User)
	//var payload *models.CreatePostRequest
	//
	//if err := ctx.ShouldBindJSON(&payload); err != nil {
	//	ctx.JSON(http.StatusBadRequest, err.Error())
	//	return
	//}
	//
	//nowTime := time.Now()
	//newPost := models.Post{
	//	Title:    payload.Title,
	//	Content:  payload.Content,
	//	Image:    payload.Image,
	//	User:     currentUser.ID,
	//	CreateAt: nowTime,
	//	UpdateAt: nowTime,
	//}
}
