package controller

import (
	"fidibo-search/dto/request"
	"fidibo-search/dto/response"
	"fidibo-search/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type BooksController struct {
	booksService service.BooksService
}

func NewBooksController(service service.BooksService) *BooksController {
	return &BooksController{
		booksService: service,
	}
}

// SearchBooks 			godoc
// @Summary				Search For Books By Name.
// @Param       		search query string false "book search"
// @Description			Return the books.
// @Produce				application/json
// @Tags				books
// @Success				200 {object} response.Response{}
// @Router				/books [get]
func (controller *BooksController) SearchBooks(ctx *gin.Context) {
	log.Info().Msg("searchBooks")

	req := request.SearchBooksRequest{}
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, response.Response{
			Code:   http.StatusBadRequest,
			Status: err.Error(),
		})
		return
	}

	searchResponse, err := controller.booksService.SearchBooks(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.Response{
			Code:   http.StatusBadRequest,
			Status: err.Error(),
		})
		return
	}

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   searchResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}
