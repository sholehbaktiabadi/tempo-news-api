package dto

type ArticleCreateRequestDto struct {
	Title   string `json:"title" validate:"required"`
	Content string `json:"content" validate:"required"`
}

type ArticleUpdateRequestDto struct {
	Title   string `json:"title" validate:"omitempty"`
	Content string `json:"content" validate:"omitempty"`
}
