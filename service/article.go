package service

import (
	"tempo-news-api/dto"
	"tempo-news-api/entity"
	"tempo-news-api/repository"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type articleService struct {
	articleRepo repository.ArticleRepository
}

type ArticleService interface {
	GetOne(id int) (entity.Article, error)
	Create(dto dto.ArticleCreateRequestDto) error
	Update(dto dto.ArticleUpdateRequestDto, id int) error
	GetAll() ([]entity.Article, error)
	Delete(id int) error
}

func NewArticleService(db *gorm.DB) ArticleService {
	return &articleService{articleRepo: repository.NewArticleRepository(db)}
}

func (s *articleService) GetOne(id int) (entity.Article, error) {
	article, err := s.articleRepo.GetOne(id)
	if err != nil {
		return article, err
	}
	return article, nil
}

func (s *articleService) GetAll() ([]entity.Article, error) {
	article, err := s.articleRepo.GetAll()
	if err != nil {
		return article, err
	}
	return article, nil
}

func (s *articleService) Create(dto dto.ArticleCreateRequestDto) error {
	article := entity.Article{}
	copier.Copy(&article, &dto)
	err := s.articleRepo.Create(article)
	if err != nil {
		return err
	}
	return nil
}

func (s *articleService) Update(dto dto.ArticleUpdateRequestDto, id int) error {
	article := entity.Article{}
	copier.Copy(&article, &dto)
	err := s.articleRepo.Update(article, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *articleService) Delete(id int) error {
	err := s.articleRepo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
