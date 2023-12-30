package repository

import "fidibo-search/model"

type BooksRepository interface {
	SearchBooks(keyword string) (books []model.Books, err error)
}
