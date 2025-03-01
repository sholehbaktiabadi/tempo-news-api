package dto

type ArticleCreateRequestDto struct {
	Title   string `json:"title" validate:"required"`
	Content string `json:"content" validate:"required"`
	Author  string `json:"author" validate:"required"`
}

type ArticleUpdateRequestDto struct {
	Title   string `json:"title" validate:"omitempty"`
	Content string `json:"content" validate:"omitempty"`
	Author  string `json:"author" validate:"omitempty"`
}

type ArticleGetAllQueryRequest struct {
	Search string `query:"search"`
	Author string `query:"author"`
}
