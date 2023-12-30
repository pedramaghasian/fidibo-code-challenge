package request

type SearchBooksRequest struct {
	Search string `form:"search" binding:"omitempty"`
}
