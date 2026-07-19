package main

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

// GET /api/team
func getTeam(c echo.Context) error {
	rows, err := db.Query(`SELECT id, name, role, photo_url, sort_order, created_at
		FROM team_members ORDER BY sort_order ASC, id ASC`)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "ошибка чтения команды"})
	}
	defer rows.Close()

	members := []TeamMember{}
	for rows.Next() {
		var m TeamMember
		if err := rows.Scan(&m.ID, &m.Name, &m.Role, &m.PhotoURL, &m.SortOrder, &m.CreatedAt); err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": "ошибка чтения строки"})
		}
		members = append(members, m)
	}
	return c.JSON(http.StatusOK, members)
}

// GET /api/news
func getNews(c echo.Context) error {
	rows, err := db.Query(`SELECT id, title, content, created_at
		FROM news ORDER BY created_at DESC`)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "ошибка чтения новостей"})
	}
	defer rows.Close()

	items := []News{}
	for rows.Next() {
		var n News
		if err := rows.Scan(&n.ID, &n.Title, &n.Content, &n.CreatedAt); err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": "ошибка чтения строки"})
		}
		items = append(items, n)
	}
	return c.JSON(http.StatusOK, items)
}

// POST /api/apply - форма заявки на вступление в клуб (без верификации)
func postApply(c echo.Context) error {
	var in struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}
	if err := c.Bind(&in); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "некорректные данные"})
	}

	in.Name = strings.TrimSpace(in.Name)
	in.Email = strings.TrimSpace(in.Email)

	if in.Name == "" || in.Email == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "имя и почта обязательны"})
	}

	var id int
	err := db.QueryRow(`INSERT INTO applications (name, email) VALUES ($1, $2) RETURNING id`,
		in.Name, in.Email).Scan(&id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "не удалось сохранить заявку"})
	}

	return c.JSON(http.StatusCreated, echo.Map{"id": id, "status": "ok"})
}
