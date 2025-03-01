package entity

import "time"

type Article struct {
	Id        int       `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	Title     string    `json:"title" gorm:"column:title"`
	Content   string    `json:"content" gorm:"column:content"`
	Author    string    `json:"author" gorm:"column:author"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (Article) TableName() string {
	return "article"
}
