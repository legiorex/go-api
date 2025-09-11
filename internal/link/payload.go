package link

type LinkRequest struct {
	Url string `json:"url" validate:"required,url"`
}

type LinkUpdateRequest struct {
	Url  string `json:"url" validate:"required,url"`
	Hash string `json:"hash"`
}

type LinksPaginationPayload struct {
	Links []Link `json:"links"`
	Total int64  `json:"total"`
}
