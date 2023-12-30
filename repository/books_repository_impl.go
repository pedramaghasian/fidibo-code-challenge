package repository

import (
	"context"
	"encoding/json"
	"errors"
	"fidibo-search/model"
	"fmt"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type BooksRepositoryImpl struct {
	Db    *gorm.DB
	Cache *redis.Client
}

func NewBooksRepositoryImpl(Db *gorm.DB, Cache *redis.Client) BooksRepository {
	return &BooksRepositoryImpl{Db: Db, Cache: Cache}
}

func (t *BooksRepositoryImpl) SearchBooks(keyword string) (books []model.Books, err error) {
	cacheResult, err := t.Cache.Get(context.Background(), keyword).Result()
	if err == nil {
		err := json.Unmarshal([]byte(cacheResult), &books)
		if err != nil {
			return nil, fmt.Errorf("error unmarshalling data from Redis: %w", err)
		}
		return books, nil
	}

	result := t.Db.Where("title LIKE ?", "%"+keyword+"%").Find(&books)
	if result.Error != nil {
		return nil, fmt.Errorf("error querying database: %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return nil, errors.New("no book found")
	}

	jsonData, err := json.Marshal(books)
	if err != nil {
		return nil, fmt.Errorf("error marshalling data for Redis: %w", err)
	}

	if err := t.Cache.Set(context.Background(), keyword, jsonData, 0).Err(); err != nil {
		fmt.Println("Failed to set data in Redis:", err)
	}

	return books, nil
}
