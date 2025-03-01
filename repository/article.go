package repository

import (
	"tempo-news-api/dto"
	"tempo-news-api/entity"

	"gorm.io/gorm"
)

type ArticleRepository interface {
	GetOne(id int) (entity.Article, error)
	Create(e entity.Article) (entity.Article, error)
	Update(e entity.Article, id int) (entity.Article, error)
	GetAll(dto dto.ArticleGetAllQueryRequest) ([]entity.Article, error)
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

func (r *articleRepo) GetAll(dto dto.ArticleGetAllQueryRequest) ([]entity.Article, error) {
	var articles []entity.Article
	query := r.db.Model(&entity.Article{}).Order("created_at desc")
	if dto.Search != "" {
		searchPattern := "%" + dto.Search + "%"
		query = query.Where("title ILIKE ? OR content ILIKE ?", searchPattern, searchPattern)
	}
	if dto.Author != "" {
		query = query.Where("LOWER(author) = LOWER(?)", dto.Author)
	}
	err := query.Find(&articles).Error
	return articles, err
}

func (r *articleRepo) Create(e entity.Article) (entity.Article, error) {
	err := r.db.Create(&e).Error
	return e, err
}

func (r *articleRepo) Update(e entity.Article, id int) (entity.Article, error) {
	err := r.db.Where("id = ?", id).Updates(&e).Error
	return e, err
}

func (r *articleRepo) Delete(id int) error {
	return r.db.Delete(entity.Article{}, id).Error
}
