package repository

import (
	"tempo-news-api/entity"

	"gorm.io/gorm"
)

type ArticleRepository interface {
	GetOne(id int) (entity.Article, error)
	Create(e entity.Article) error
	Update(e entity.Article, id int) error
	GetAll() ([]entity.Article, error)
	Delete(id int) error
}

type articleRepo struct {
	db *gorm.DB
}

func NewArticleRepository(db *gorm.DB) ArticleRepository {
	return &articleRepo{db}
}

func (r *articleRepo) GetOne(id int) (entity.Article, error) {
	var article entity.Article
	err := r.db.First(&article, id).Error
	return article, err
}

func (r *articleRepo) GetAll() ([]entity.Article, error) {
	var articles []entity.Article
	err := r.db.Order("created_at desc").Find(&articles)
	return articles, err.Error
}

func (r *articleRepo) Create(e entity.Article) error {
	return r.db.Create(&e).Error
}

func (r *articleRepo) Update(e entity.Article, id int) error {
	return r.db.Where("id = ?", id).Updates(e).Error
}

func (r *articleRepo) Delete(id int) error {
	return r.db.Delete(entity.Article{}, id).Error
}
