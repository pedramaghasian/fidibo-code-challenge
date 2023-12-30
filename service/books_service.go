package service

import (
	"fidibo-search/dto/request"
	"fidibo-search/dto/response"
)

type BooksService interface {
	SearchBooks(keyword request.SearchBooksRequest) (books []response.SearchBooksResponse, err error)
}
