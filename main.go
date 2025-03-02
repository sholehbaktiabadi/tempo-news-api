package main

import (
	"fmt"
	"tempo-news-api/config"
	"tempo-news-api/controller"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/redis/go-redis/v9"
)

func main() {
	env := config.Loadenv()
	db := config.InitDB()
	route := echo.New()
	route.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PATCH", "DELETE"},
	}))
	redisAddr := fmt.Sprintf("%s:%s", env.Redis.Host, env.Redis.Port)
	apiV1 := route.Group("api/v1")
	articleRoute := apiV1.Group("/article")
	fmt.Println(redisAddr)
	redisClient := redis.NewClient(&redis.Options{Addr: redisAddr})
	articleController := controller.NewArticleController(db, redisClient)
	articleRoute.GET("", articleController.GetAll)
	articleRoute.POST("", articleController.Create)
	articleRoute.GET("/:id", articleController.GetOne)
	articleRoute.PATCH("/:id", articleController.Update)
	articleRoute.DELETE("/:id", articleController.Delete)
	route.Start(":" + env.App.Port)
}
