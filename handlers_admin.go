package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// ---------- Команда ----------

// POST /api/admin/team
func createTeamMember(c echo.Context) error {
	var m TeamMember
	if err := c.Bind(&m); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "некорректные данные"})
	}
	err := db.QueryRow(`INSERT INTO team_members (name, role, photo_url, sort_order)
		VALUES ($1, $2, $3, $4) RETURNING id, created_at`,
		m.Name, m.Role, m.PhotoURL, m.SortOrder).Scan(&m.ID, &m.CreatedAt)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "не удалось создать участника"})
	}
	return c.JSON(http.StatusCreated, m)
}

// PUT /api/admin/team/:id
func updateTeamMember(c echo.Context) error {
	id := c.Param("id")
	var m TeamMember
	if err := c.Bind(&m); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "некорректные данные"})
	}
	res, err := db.Exec(`UPDATE team_members SET name=$1, role=$2, photo_url=$3, sort_order=$4 WHERE id=$5`,
		m.Name, m.Role, m.PhotoURL, m.SortOrder, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "не удалось обновить участника"})
	}
	if n, _ := res.RowsAffected(); n == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "участник не найден"})
	}
	return c.JSON(http.StatusOK, echo.Map{"status": "ok"})
}

// DELETE /api/admin/team/:id
func deleteTeamMember(c echo.Context) error {
	id := c.Param("id")
	res, err := db.Exec(`DELETE FROM team_members WHERE id=$1`, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "не удалось удалить участника"})
	}
	if n, _ := res.RowsAffected(); n == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "участник не найден"})
	}
	return c.JSON(http.StatusOK, echo.Map{"status": "ok"})
}

// ---------- Новости ----------

// POST /api/admin/news
func createNews(c echo.Context) error {
	var n News
	if err := c.Bind(&n); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "некорректные данные"})
	}
	err := db.QueryRow(`INSERT INTO news (title, content) VALUES ($1, $2) RETURNING id, created_at`,
		n.Title, n.Content).Scan(&n.ID, &n.CreatedAt)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "не удалось создать новость"})
	}
	return c.JSON(http.StatusCreated, n)
}

// PUT /api/admin/news/:id
func updateNews(c echo.Context) error {
	id := c.Param("id")
	var n News
	if err := c.Bind(&n); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "некорректные данные"})
	}
	res, err := db.Exec(`UPDATE news SET title=$1, content=$2 WHERE id=$3`, n.Title, n.Content, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "не удалось обновить новость"})
	}
	if rows, _ := res.RowsAffected(); rows == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "новость не найдена"})
	}
	return c.JSON(http.StatusOK, echo.Map{"status": "ok"})
}

// DELETE /api/admin/news/:id
func deleteNews(c echo.Context) error {
	id := c.Param("id")
	res, err := db.Exec(`DELETE FROM news WHERE id=$1`, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "не удалось удалить новость"})
	}
	if rows, _ := res.RowsAffected(); rows == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "новость не найдена"})
	}
	return c.JSON(http.StatusOK, echo.Map{"status": "ok"})
}

// ---------- Заявки ----------

// GET /api/admin/applications
func getApplications(c echo.Context) error {
	rows, err := db.Query(`SELECT id, name, email, created_at FROM applications ORDER BY created_at DESC`)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "ошибка чтения заявок"})
	}
	defer rows.Close()

	items := []Application{}
	for rows.Next() {
		var a Application
		if err := rows.Scan(&a.ID, &a.Name, &a.Email, &a.CreatedAt); err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": "ошибка чтения строки"})
		}
		items = append(items, a)
	}
	return c.JSON(http.StatusOK, items)
}

// DELETE /api/admin/applications/:id
func deleteApplication(c echo.Context) error {
	id := c.Param("id")
	res, err := db.Exec(`DELETE FROM applications WHERE id=$1`, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "не удалось удалить заявку"})
	}
	if rows, _ := res.RowsAffected(); rows == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "заявка не найдена"})
	}
	return c.JSON(http.StatusOK, echo.Map{"status": "ok"})
}
