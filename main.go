package main

import (
	"tempo-news-api/config"
	"tempo-news-api/controller"

	"github.com/labstack/echo/v4"
)

func main() {
	db := config.InitDB()
	route := echo.New()
	apiV1 := route.Group("api/v1")
	articleController := controller.NewArticleController(db)
	apiV1.GET("/article/:id", articleController.GetOne)
	apiV1.POST("/article", articleController.Create)
	apiV1.PATCH("/article/:id", articleController.Update)
	apiV1.GET("/article", articleController.GetAll)
	apiV1.DELETE("/article/:id", articleController.Delete)
	route.Start(":8000")
}
