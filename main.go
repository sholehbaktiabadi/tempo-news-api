package main

import (
	"database/sql"
	"fmt"
	"log"
	"tempo-news-api/config"
	"tempo-news-api/controller"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
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
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get *sql.DB: %v", err)
	}
	if err := runMigrations(sqlDB); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}
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

func runMigrations(db *sql.DB) error {
	goose.SetDialect("postgres")

	err := goose.Up(db, "migration")
	if err != nil {
		return fmt.Errorf("goose migration failed: %w", err)
	}
	fmt.Println("Migrations applied successfully")
	return nil
}
