package controller

import (
	"strconv"
	"tempo-news-api/helper"
	"tempo-news-api/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type ArticleController struct {
	articleService service.ArticleService
}

func NewArticleController(db *gorm.DB) ArticleController {
	service := service.NewArticleService(db)
	controller := ArticleController{
		articleService: service,
	}
	return controller
}

func (r ArticleController) GetByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return helper.ResErrHandler(c, err)
	}
	data, err := r.articleService.GetByID(c, id)
	if err != nil {
		return helper.ResErrHandler(c, err)
	}
	return helper.ResOK(c, "ok", data)
}
