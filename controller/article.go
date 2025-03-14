package controller

import (
	"strconv"
	"tempo-news-api/dto"
	"tempo-news-api/helper"
	"tempo-news-api/service"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type ArticleController struct {
	articleService service.ArticleService
}

func NewArticleController(db *gorm.DB, redisClient *redis.Client) ArticleController {
	service := service.NewArticleService(db, redisClient)
	controller := ArticleController{
		articleService: service,
	}
	return controller
}

func (r ArticleController) GetOne(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return helper.ResErrHandler(c, err)
	}
	data, err := r.articleService.GetOne(id)
	if err != nil {
		return helper.ResErrHandler(c, err)
	}
	return helper.ResOK(c, "ok", data)
}

func (r ArticleController) GetAll(c echo.Context) error {
	params := new(dto.ArticleGetAllQueryRequest)
	if err := c.Bind(params); err != nil {
		return helper.ResErrHandler(c, err)
	}
	data, err := r.articleService.GetAll(*params)
	if err != nil {
		return helper.ResErrHandler(c, err)
	}
	return helper.ResOK(c, "ok", data)
}

func (r ArticleController) Create(c echo.Context) error {
	var payload dto.ArticleCreateRequestDto

	if err := c.Bind(&payload); err != nil {
		return helper.ResErrHandler(c, err)
	}

	if err := validator.New(validator.WithRequiredStructEnabled()).Struct(payload); err != nil {
		return helper.ResErrHandler(c, err)
	}

	err := r.articleService.Create(payload)
	if err != nil {
		return helper.ResErrHandler(c, err)
	}

	return helper.ResOK(c, "ok", nil)
}

func (r ArticleController) Update(c echo.Context) error {
	var payload dto.ArticleUpdateRequestDto
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return helper.ResErrHandler(c, err)
	}

	if err := c.Bind(&payload); err != nil {
		return helper.ResErrHandler(c, err)
	}

	if err := validator.New(validator.WithRequiredStructEnabled()).Struct(payload); err != nil {
		return helper.ResErrHandler(c, err)
	}

	if err := r.articleService.Update(payload, id); err != nil {
		return helper.ResErrHandler(c, err)
	}

	return helper.ResOK(c, "ok", nil)
}

func (r ArticleController) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return helper.ResErrHandler(c, err)
	}
	err = r.articleService.Delete(id)
	if err != nil {
		return helper.ResErrHandler(c, err)
	}
	return helper.ResOK(c, "ok", nil)
}
