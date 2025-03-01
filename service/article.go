package service

import (
	"tempo-news-api/entity"
	"tempo-news-api/repository"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type articleService struct {
	articleRepo repository.ArticleRepository
}

type ArticleService interface {
	GetByID(c echo.Context, id int) (entity.Article, error)
}

func NewArticleService(db *gorm.DB) ArticleService {
	return &articleService{articleRepo: repository.NewArticleRepository(db)}
}

func (s *articleService) GetByID(c echo.Context, id int) (entity.Article, error) {
	article, err := s.articleRepo.GetByID(id)
	if err != nil {
		return article, err
	}
	return article, nil
}
