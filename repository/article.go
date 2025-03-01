package repository

import (
	"tempo-news-api/entity"

	"gorm.io/gorm"
)

type DB struct {
	Conn *gorm.DB
}

type ArticleRepository interface {
	GetByID(id int) (entity.Article, error)
}

func NewArticleRepository(db *gorm.DB) ArticleRepository {
	return &DB{Conn: db}
}

func (db *DB) GetByID(id int) (entity.Article, error) {
	article := entity.Article{}
	result := db.Conn.Where("id = ?", id).First(&article)
	return article, result.Error
}
