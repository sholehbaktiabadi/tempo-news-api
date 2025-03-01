package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"tempo-news-api/dto"
	"tempo-news-api/entity"
	"tempo-news-api/repository"
	"tempo-news-api/variable"

	"github.com/jinzhu/copier"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type articleService struct {
	articleRepo  repository.ArticleRepository
	cacheService CacheService
}

type ArticleService interface {
	GetOne(id int) (entity.Article, error)
	Create(dto dto.ArticleCreateRequestDto) error
	Update(dto dto.ArticleUpdateRequestDto, id int) error
	GetAll(dto dto.ArticleGetAllQueryRequest) ([]entity.Article, error)
	Delete(id int) error
}

func NewArticleService(db *gorm.DB, redisClient *redis.Client) ArticleService {
	return &articleService{
		articleRepo:  repository.NewArticleRepository(db),
		cacheService: NewCacheService(redisClient),
	}
}

func (s *articleService) GetOne(id int) (entity.Article, error) {
	var article entity.Article
	key := fmt.Sprintf("%s-%d", variable.ArticlePrefixKey, id)
	isCacheExist, err := s.cacheService.Exists(key)
	if err != nil {
		return article, err
	}
	if isCacheExist {
		fmt.Println("from cache")
		fetchedData, err := s.cacheService.Get(key)
		if err != nil {
			return article, err
		}
		err = json.Unmarshal([]byte(fetchedData.(string)), &article)
		if err != nil {
			return article, errors.New("failed to parse cached article")
		}
		return article, nil
	}
	fmt.Println("from db")
	article, err = s.articleRepo.GetOne(id)
	if err != nil {
		return article, err
	}
	articleJSON, _ := json.Marshal(article)
	_ = s.cacheService.Set(key, string(articleJSON))
	return article, nil
}

func (s *articleService) Create(dto dto.ArticleCreateRequestDto) error {
	var article entity.Article
	copier.Copy(&article, &dto)
	data, err := s.articleRepo.Create(article)
	if err != nil {
		return err
	}
	key := fmt.Sprintf("%s-%d", variable.ArticlePrefixKey, data.Id)
	articleJSON, _ := json.Marshal(data)
	err = s.cacheService.Set(key, string(articleJSON))
	if err != nil {
		fmt.Println("Failed to store in cache:", err)
	}
	return nil
}

func (s *articleService) GetAll(dto dto.ArticleGetAllQueryRequest) ([]entity.Article, error) {
	article, err := s.articleRepo.GetAll(dto)
	if err != nil {
		return article, err
	}
	return article, nil
}

func (s *articleService) Update(dto dto.ArticleUpdateRequestDto, id int) error {
	article := entity.Article{Id: id}
	copier.Copy(&article, &dto)
	data, err := s.articleRepo.Update(article, id)
	if err != nil {
		return err
	}
	fmt.Println(data)
	key := fmt.Sprintf("%s-%d", variable.ArticlePrefixKey, data.Id)
	articleJSON, _ := json.Marshal(data)
	err = s.cacheService.Set(key, string(articleJSON))
	if err != nil {
		fmt.Println("Failed to store in cache:", err)
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
