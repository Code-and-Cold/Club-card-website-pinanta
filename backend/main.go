package main

import (
	"log"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println(".env не найден, используются переменные окружения")
	}
	
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
	// Фронт при монтировании страницы дёргает GET /api/site и получает разом все блоки
	api.GET("/site", getSite)
	api.POST("/apply", postApply)

	// ---- Админские маршруты, защищены Basic Auth ----
	adminLogin := envOr("ADMIN_LOGIN", "admin")
	adminPassword := envOr("ADMIN_PASSWORD", "admin123")

	admin := api.Group("/admin")
	admin.Use(middleware.BasicAuth(func(login, password string, c echo.Context) (bool, error) {
		return login == adminLogin && password == adminPassword, nil
	}))

	// Простые текстовые/JSON блоки: hero, warmup, why_us, team, benefits, quote, news_block, apply_block
	admin.GET("/content/:key", getSiteBlock)
	admin.PUT("/content/:key", putSiteBlock)

	// Команда
	admin.GET("/team", adminListTeam)
	admin.POST("/team", createTeamMember)
	admin.PUT("/team/:id", updateTeamMember)
	admin.PATCH("/team/:id/hidden", setTeamMemberHidden)
	admin.DELETE("/team/:id", deleteTeamMember)

	// Карточки "почему у нас классно"
	admin.GET("/why-us-cards", adminListWhyUsCards)
	admin.POST("/why-us-cards", createWhyUsCard)
	admin.PUT("/why-us-cards/:id", updateWhyUsCard)
	admin.PATCH("/why-us-cards/:id/hidden", setWhyUsCardHidden)
	admin.DELETE("/why-us-cards/:id", deleteWhyUsCard)

	// Преимущества "что ты получишь"
	admin.GET("/benefits", adminListBenefits)
	admin.POST("/benefits", createBenefit)
	admin.PUT("/benefits/:id", updateBenefit)
	admin.PATCH("/benefits/:id/hidden", setBenefitHidden)
	admin.DELETE("/benefits/:id", deleteBenefit)

	// Новости
	admin.GET("/news", adminListNews)
	admin.POST("/news", createNews)
	admin.PUT("/news/:id", updateNews)
	admin.PATCH("/news/:id/hidden", setNewsHidden)
	admin.PATCH("/news/:id/postpone", postponeNews)
	admin.DELETE("/news/:id", deleteNews)

	// Выпадающие списки формы заявки (высшая школа / курс)
	admin.GET("/dropdown-options", adminListDropdownOptions)
	admin.POST("/dropdown-options", createDropdownOption)
	admin.PUT("/dropdown-options/:id", updateDropdownOption)
	admin.DELETE("/dropdown-options/:id", deleteDropdownOption)

	// Заявки
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
