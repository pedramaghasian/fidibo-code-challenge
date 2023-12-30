package response

type SearchBooksResponse struct {
	Title string `json:"title"`
	Cover string `json:"cover"`
	Type  uint8  `json:"type"`
}
