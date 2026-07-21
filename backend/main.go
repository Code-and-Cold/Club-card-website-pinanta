package main

import (
	"log"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	db = connectDB()
	defer db.Close()

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS()) // на случай, если фронт будет открыт отдельно от сервера

	// Статика: index.html - публичная страница, admin.html - админка
	e.Static("/", "static")

	api := e.Group("/api")

	// ---- Публичные маршруты ----
	api.GET("/team", getTeam)
	api.GET("/news", getNews)
	api.POST("/apply", postApply)

	// ---- Админские маршруты, защищены Basic Auth ----
	adminLogin := envOr("ADMIN_LOGIN", "admin")
	adminPassword := envOr("ADMIN_PASSWORD", "admin123")

	admin := api.Group("/admin")
	admin.Use(middleware.BasicAuth(func(login, password string, c echo.Context) (bool, error) {
		return login == adminLogin && password == adminPassword, nil
	}))

	admin.POST("/team", createTeamMember)
	admin.PUT("/team/:id", updateTeamMember)
	admin.DELETE("/team/:id", deleteTeamMember)

	admin.POST("/news", createNews)
	admin.PUT("/news/:id", updateNews)
	admin.DELETE("/news/:id", deleteNews)

	admin.GET("/applications", getApplications)
	admin.DELETE("/applications/:id", deleteApplication)

	port := envOr("PORT", "8080")
	log.Printf("сервер запущен на порту %s", port)
	e.Logger.Fatal(e.Start(":" + port))
}

func envOr(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
