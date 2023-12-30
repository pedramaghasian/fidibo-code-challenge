package service

import (
	"fidibo-search/dto/request"
	"fidibo-search/dto/response"
	"fidibo-search/repository"
	"strings"

	"github.com/go-playground/validator/v10"
)

type BooksServiceImpl struct {
	BooksRepository repository.BooksRepository
	Validate        *validator.Validate
}

func NewBooksServiceImpl(tagRepository repository.BooksRepository, validate *validator.Validate) BooksService {
	return &BooksServiceImpl{
		BooksRepository: tagRepository,
		Validate:        validate,
	}
}

func (t *BooksServiceImpl) SearchBooks(keyword request.SearchBooksRequest) ([]response.SearchBooksResponse, error) {
	if err := t.Validate.Struct(keyword); err != nil {
		return nil, err
	}

	searchKeyword := strings.TrimSpace(strings.ToLower(keyword.Search))

	books, err := t.BooksRepository.SearchBooks(searchKeyword)
	if err != nil {
		return nil, err
	}

	res := []response.SearchBooksResponse{}
	for _, book := range books {
		res = append(res, response.SearchBooksResponse{
			Title: book.Title,
			Cover: book.Cover,
			Type:  book.Type,
		})
	}

	return res, nil
}
