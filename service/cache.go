package service

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type CacheService interface {
	Set(key string, value any) error
	Get(key string) (any, error)
	Exists(key string) (bool, error)
}

type cacheService struct {
	client *redis.Client
	ctx    context.Context
}

func NewCacheService(redisClient *redis.Client) CacheService {
	return &cacheService{
		client: redisClient,
		ctx:    context.Background(),
	}
}

func (s *cacheService) Set(key string, value any) error {
	return s.client.Set(s.ctx, key, value, time.Hour).Err()
}

func (s *cacheService) Get(key string) (any, error) {
	return s.client.Get(s.ctx, key).Result()
}

func (s *cacheService) Exists(key string) (bool, error) {
	exists, err := s.client.Exists(s.ctx, key).Result()
	return exists > 0, err
}
