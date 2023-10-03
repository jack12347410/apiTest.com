package models

import (
	"github.com/gin-gonic/gin"
	"time"
)

type Test struct {
	ID       uint      `json:"id" gorm:"primaryKey"`
	Title    string    `json:"title"`
	Content  string    `json:"content"`
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
}

type CreateTestInput struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required""`
}

type UpdateTestInput struct {
	Title    string    `json:"title"`
	Content  string    `json:"content"`
	UpdateAt time.Time `json:"update_at"`
}

type TestService interface {
	Create(input *CreateTestInput) (*Test, error)
	FindAll() (*[]Test, error)
	FindById(id string) (*Test, error)
	Update(id string, input *UpdateTestInput) (*Test, error)
	Delete(id int32) (*Test, error)
}

type TestRepository interface {
	Create(create *Test) (*Test, error)
	FindAll() (*[]Test, error)
	FindById(id string) (*Test, error)
	Update(post *Test, update *Test) (*Test, error)
	Delete(delete *Test) (*Test, error)
	Migrate() error
}

type TestController interface {
	CreateTest(ctx *gin.Context)
	FindTests(ctx *gin.Context)
	FindOneTest(ctx *gin.Context)
	UpdateTest(ctx *gin.Context)
	DeleteTest(ctx *gin.Context)
}
