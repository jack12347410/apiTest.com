package models

import "time"

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
