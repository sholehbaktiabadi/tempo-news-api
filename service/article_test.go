package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"testing"

	"tempo-news-api/dto"
	"tempo-news-api/entity"
	"tempo-news-api/repository"
	"tempo-news-api/variable"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestArticleService_GetOne(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockArticleRepo := repository.NewMockArticleRepository(ctrl)
	mockCacheService := NewMockCacheService(ctrl)

	articleService := &articleService{
		articleRepo:  mockArticleRepo,
		cacheService: mockCacheService,
	}

	// Article found in cache
	t.Run("Article found in cache", func(t *testing.T) {
		article := entity.Article{Id: 1, Title: "Test Article", Content: "Test Content", Author: "Test Author"}
		articleJSON, _ := json.Marshal(article)
		key := fmt.Sprintf("%s-%d", variable.ArticlePrefixKey, article.Id)

		mockCacheService.EXPECT().Exists(key).Return(true, nil)
		mockCacheService.EXPECT().Get(key).Return(string(articleJSON), nil)

		result, err := articleService.GetOne(article.Id)
		assert.NoError(t, err)
		assert.Equal(t, article, result)
	})

	// Article not found in cache, fetch from DB
	t.Run("Article not found in cache, fetch from DB", func(t *testing.T) {
		article := entity.Article{Id: 1, Title: "Test Article", Content: "Test Content", Author: "Test Author"}
		key := fmt.Sprintf("%s-%d", variable.ArticlePrefixKey, article.Id)

		mockCacheService.EXPECT().Exists(key).Return(false, nil)
		mockArticleRepo.EXPECT().GetOne(article.Id).Return(article, nil)

		result, err := articleService.GetOne(article.Id)
		assert.NoError(t, err)
		assert.Equal(t, article, result)
	})

	// Error fetching from cache
	t.Run("Error fetching from cache", func(t *testing.T) {
		article := entity.Article{Id: 1, Title: "Test Article", Content: "Test Content", Author: "Test Author"}
		key := fmt.Sprintf("%s-%d", variable.ArticlePrefixKey, article.Id)

		mockCacheService.EXPECT().Exists(key).Return(false, errors.New("cache error"))

		_, err := articleService.GetOne(article.Id)
		assert.Error(t, err)
	})

	// Error fetching from DB
	t.Run("Error fetching from DB", func(t *testing.T) {
		article := entity.Article{Id: 1, Title: "Test Article", Content: "Test Content", Author: "Test Author"}
		key := fmt.Sprintf("%s-%d", variable.ArticlePrefixKey, article.Id)

		mockCacheService.EXPECT().Exists(key).Return(false, nil)
		mockArticleRepo.EXPECT().GetOne(article.Id).Return(entity.Article{}, errors.New("DB error"))

		_, err := articleService.GetOne(article.Id)
		assert.Error(t, err)
	})
}

func TestArticleService_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockArticleRepo := repository.NewMockArticleRepository(ctrl)
	mockCacheService := NewMockCacheService(ctrl)

	articleService := &articleService{
		articleRepo:  mockArticleRepo,
		cacheService: mockCacheService,
	}

	// Successfully create article
	t.Run("Successfully create article", func(t *testing.T) {
		articleDto := dto.ArticleCreateRequestDto{Title: "Test Article", Content: "Test Content", Author: "Test Author"}
		article := entity.Article{Title: "Test Article", Content: "Test Content", Author: "Test Author"}
		createdArticle := entity.Article{Id: 1, Title: "Test Article", Content: "Test Content", Author: "Test Author"}
		key := fmt.Sprintf("%s-%d", variable.ArticlePrefixKey, createdArticle.Id)

		mockArticleRepo.EXPECT().Create(article).Return(createdArticle, nil)
		mockCacheService.EXPECT().Set(key, gomock.Any()).Return(nil)

		err := articleService.Create(articleDto)
		assert.NoError(t, err)
	})

	// Error creating article
	t.Run("Error creating article", func(t *testing.T) {
		articleDto := dto.ArticleCreateRequestDto{Title: "Test Article", Content: "Test Content", Author: "Test Author"}
		article := entity.Article{Title: "Test Article", Content: "Test Content", Author: "Test Author"}

		mockArticleRepo.EXPECT().Create(article).Return(entity.Article{}, errors.New("DB error"))

		err := articleService.Create(articleDto)
		assert.Error(t, err)
	})
}

func TestArticleService_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockArticleRepo := repository.NewMockArticleRepository(ctrl)
	mockCacheService := NewMockCacheService(ctrl)

	articleService := &articleService{
		articleRepo:  mockArticleRepo,
		cacheService: mockCacheService,
	}

	// Successfully update article
	t.Run("Successfully update article", func(t *testing.T) {
		articleDto := dto.ArticleUpdateRequestDto{Title: "Updated Article", Content: "Updated Content", Author: "Updated Author"}
		article := entity.Article{Id: 1, Title: "Updated Article", Content: "Updated Content", Author: "Updated Author"}
		key := fmt.Sprintf("%s-%d", variable.ArticlePrefixKey, article.Id)

		mockArticleRepo.EXPECT().Update(article, article.Id).Return(article, nil)
		mockCacheService.EXPECT().Set(key, gomock.Any()).Return(nil)

		err := articleService.Update(articleDto, article.Id)
		assert.NoError(t, err)
	})

	// Error updating article
	t.Run("Error updating article", func(t *testing.T) {
		articleDto := dto.ArticleUpdateRequestDto{Title: "Updated Article", Content: "Updated Content", Author: "Updated Author"}
		article := entity.Article{Id: 1, Title: "Updated Article", Content: "Updated Content", Author: "Updated Author"}

		mockArticleRepo.EXPECT().Update(article, article.Id).Return(entity.Article{}, errors.New("DB error"))

		err := articleService.Update(articleDto, article.Id)
		assert.Error(t, err)
	})
}

func TestArticleService_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockArticleRepo := repository.NewMockArticleRepository(ctrl)
	mockCacheService := NewMockCacheService(ctrl)

	articleService := &articleService{
		articleRepo:  mockArticleRepo,
		cacheService: mockCacheService,
	}

	// Successfully delete article
	t.Run("Successfully delete article", func(t *testing.T) {
		articleId := 1

		mockArticleRepo.EXPECT().Delete(articleId).Return(nil)

		err := articleService.Delete(articleId)
		assert.NoError(t, err)
	})

	// Error deleting article
	t.Run("Error deleting article", func(t *testing.T) {
		articleId := 1

		mockArticleRepo.EXPECT().Delete(articleId).Return(errors.New("DB error"))

		err := articleService.Delete(articleId)
		assert.Error(t, err)
	})
}

func TestArticleService_GetAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockArticleRepo := repository.NewMockArticleRepository(ctrl)
	mockCacheService := NewMockCacheService(ctrl)

	articleService := &articleService{
		articleRepo:  mockArticleRepo,
		cacheService: mockCacheService,
	}

	// Successfully get all articles
	t.Run("Successfully get all articles", func(t *testing.T) {
		query := dto.ArticleGetAllQueryRequest{Search: "test", Author: "test author"}
		articles := []entity.Article{
			{Id: 1, Title: "Test Article 1", Content: "Test Content 1", Author: "Test Author 1"},
			{Id: 2, Title: "Test Article 2", Content: "Test Content 2", Author: "Test Author 2"},
		}

		mockArticleRepo.EXPECT().GetAll(query).Return(articles, nil)

		result, err := articleService.GetAll(query)
		assert.NoError(t, err)
		assert.Equal(t, articles, result)
	})

	// Error getting all articles
	t.Run("Error getting all articles", func(t *testing.T) {
		query := dto.ArticleGetAllQueryRequest{Search: "test", Author: "test author"}

		mockArticleRepo.EXPECT().GetAll(query).Return(nil, errors.New("DB error"))

		_, err := articleService.GetAll(query)
		assert.Error(t, err)
	})
}
