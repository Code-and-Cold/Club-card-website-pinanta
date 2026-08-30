package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// ---------- Простые блоки сайта (hero, warmup, why_us, team, benefits, quote, news_block, apply_block) ----------

// GET /api/admin/content/:key
func getSiteBlock(c echo.Context) error {
	key := c.Param("key")
	if !isValidSiteBlockKey(key) {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "неизвестный блок"})
	}
	var content JSONMap
	err := db.QueryRow(`SELECT content FROM site_blocks WHERE key = $1`, key).Scan(&content)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "ошибка чтения блока"})
	}
	return c.JSON(http.StatusOK, echo.Map{"key": key, "content": content})
}

// PUT /api/admin/content/:key
// body: {"content": {произвольные поля блока}}
func putSiteBlock(c echo.Context) error {
	key := c.Param("key")
	if !isValidSiteBlockKey(key) {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "неизвестный блок"})
	}
	var body struct {
		Content JSONMap `json:"content"`
	}
	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "некорректные данные"})
	}
	normalizeSiteBlockContent(body.Content)
	if err := updateSiteBlock(key, body.Content); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "не удалось обновить блок"})
	}
	return c.JSON(http.StatusOK, echo.Map{"status": "ok"})
}

// ---------- Команда ----------

// GET /api/admin/team - включая скрытых участников
func adminListTeam(c echo.Context) error {
	items, err := listTeamMembers(true)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "ошибка чтения команды"})
	}
	return c.JSON(http.StatusOK, items)
}

// POST /api/admin/team
func createTeamMember(c echo.Context) error {
	var m TeamMember
	if err := c.Bind(&m); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "некорректные данные"})
	}
	m.PhotoURL = normalizeAssetURL(m.PhotoURL)
	err := db.QueryRow(
		`INSERT INTO team_members (name, role, photo_url, hover_info, description, contacts, hidden, sort_order)
		 VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id, created_at`,
		m.Name, m.Role, m.PhotoURL, m.HoverInfo, m.Description, m.Contacts, m.Hidden, m.SortOrder,
	).Scan(&m.ID, &m.CreatedAt)
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
	m.PhotoURL = normalizeAssetURL(m.PhotoURL)
	res, err := db.Exec(
		`UPDATE team_members SET name=$1, role=$2, photo_url=$3, hover_info=$4, description=$5, contacts=$6, sort_order=$7 WHERE id=$8`,
		m.Name, m.Role, m.PhotoURL, m.HoverInfo, m.Description, m.Contacts, m.SortOrder, id,
	)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "не удалось обновить участника"})
	}
	if n, _ := res.RowsAffected(); n == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "участник не найден"})
	}
	return c.JSON(http.StatusOK, echo.Map{"status": "ok"})
}

// PATCH /api/admin/team/:id/hidden  body: {"hidden": true|false}
func setTeamMemberHidden(c echo.Context) error {
	return setHiddenGeneric(c, "team_members")
}

// DELETE /api/admin/team/:id
func deleteTeamMember(c echo.Context) error {
	return deleteGeneric(c, "team_members", "участник не найден")
}

// ---------- Карточки "почему у нас классно" ----------

// GET /api/admin/why-us-cards
func adminListWhyUsCards(c echo.Context) error {
	items, err := listWhyUsCards(true)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "ошибка чтения карточек"})
	}
	return c.JSON(http.StatusOK, items)
}

// POST /api/admin/why-us-cards
func createWhyUsCard(c echo.Context) error {
	var card WhyUsCard
	if err := c.Bind(&card); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "некорректные данные"})
	}
	card.IconURL = normalizeAssetURL(card.IconURL)
	err := db.QueryRow(
		`INSERT INTO why_us_cards (icon_url, title, description, hidden, sort_order)
		 VALUES ($1, $2, $3, $4, $5) RETURNING id`,
		card.IconURL, card.Title, card.Description, card.Hidden, card.SortOrder,
	).Scan(&card.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "не удалось создать карточку"})
	}
	return c.JSON(http.StatusCreated, card)
}

// PUT /api/admin/why-us-cards/:id
func updateWhyUsCard(c echo.Context) error {
	id := c.Param("id")
	var card WhyUsCard
	if err := c.Bind(&card); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "некорректные данные"})
	}
	card.IconURL = normalizeAssetURL(card.IconURL)
	res, err := db.Exec(
		`UPDATE why_us_cards SET icon_url=$1, title=$2, description=$3, sort_order=$4 WHERE id=$5`,
		card.IconURL, card.Title, card.Description, card.SortOrder, id,
	)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "не удалось обновить карточку"})
	}
	if n, _ := res.RowsAffected(); n == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "карточка не найдена"})
	}
	return c.JSON(http.StatusOK, echo.Map{"status": "ok"})
}

// PATCH /api/admin/why-us-cards/:id/hidden
func setWhyUsCardHidden(c echo.Context) error {
	return setHiddenGeneric(c, "why_us_cards")
}

// DELETE /api/admin/why-us-cards/:id
func deleteWhyUsCard(c echo.Context) error {
	return deleteGeneric(c, "why_us_cards", "карточка не найдена")
}

// ---------- Преимущества ("что ты получишь") ----------

// GET /api/admin/benefits
func adminListBenefits(c echo.Context) error {
	items, err := listBenefits(true)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "ошибка чтения преимуществ"})
	}
	return c.JSON(http.StatusOK, items)
}

// POST /api/admin/benefits
func createBenefit(c echo.Context) error {
	var b Benefit
	if err := c.Bind(&b); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "некорректные данные"})
	}
	b.IconURL = normalizeAssetURL(b.IconURL)
	err := db.QueryRow(
		`INSERT INTO benefits (icon_url, title, description, hidden, sort_order)
		 VALUES ($1, $2, $3, $4, $5) RETURNING id`,
		b.IconURL, b.Title, b.Description, b.Hidden, b.SortOrder,
	).Scan(&b.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "не удалось создать преимущество"})
	}
	return c.JSON(http.StatusCreated, b)
}

// PUT /api/admin/benefits/:id
func updateBenefit(c echo.Context) error {
	id := c.Param("id")
	var b Benefit
	if err := c.Bind(&b); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "некорректные данные"})
	}
	b.IconURL = normalizeAssetURL(b.IconURL)
	res, err := db.Exec(
		`UPDATE benefits SET icon_url=$1, title=$2, description=$3, sort_order=$4 WHERE id=$5`,
		b.IconURL, b.Title, b.Description, b.SortOrder, id,
	)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "не удалось обновить преимущество"})
	}
	if n, _ := res.RowsAffected(); n == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "преимущество не найдено"})
	}
	return c.JSON(http.StatusOK, echo.Map{"status": "ok"})
}

// PATCH /api/admin/benefits/:id/hidden
func setBenefitHidden(c echo.Context) error {
	return setHiddenGeneric(c, "benefits")
}

// DELETE /api/admin/benefits/:id
func deleteBenefit(c echo.Context) error {
	return deleteGeneric(c, "benefits", "преимущество не найдено")
}

// ---------- Новости ----------

// GET /api/admin/news - включая скрытые и отложенные
func adminListNews(c echo.Context) error {
	items, err := listNews(true)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "ошибка чтения новостей"})
	}
	return c.JSON(http.StatusOK, items)
}

// newsInput - тело запроса для создания/обновления новости.
// publish_at передаётся строкой в формате RFC3339 (например, из <input type="datetime-local">
// на фронте нужно привести к ISO 8601). Пустая строка/отсутствие поля = опубликовать сразу.
type newsInput struct {
	ImageURL  string `json:"image_url"`
	Title     string `json:"title"`
	ShortDesc string `json:"short_desc"`
	FullDesc  string `json:"full_desc"`
	Hidden    bool   `json:"hidden"`
	PublishAt string `json:"publish_at"`
}

// POST /api/admin/news
func createNews(c echo.Context) error {
	var in newsInput
	if err := c.Bind(&in); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "некорректные данные"})
	}
	in.ImageURL = normalizeAssetURL(in.ImageURL)
	publishAt, err := parsePublishAt(in.PublishAt)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "некорректная дата публикации (ожидается формат RFC3339)"})
	}

	var n News
	err = db.QueryRow(
		`INSERT INTO news (image_url, title, short_desc, full_desc, hidden, publish_at)
		 VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, created_at`,
		in.ImageURL, in.Title, in.ShortDesc, in.FullDesc, in.Hidden, publishAt,
	).Scan(&n.ID, &n.CreatedAt)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "не удалось создать новость"})
	}
	n.ImageURL, n.Title, n.ShortDesc, n.FullDesc, n.Hidden, n.PublishAt = in.ImageURL, in.Title, in.ShortDesc, in.FullDesc, in.Hidden, publishAt
	return c.JSON(http.StatusCreated, n)
}

// PUT /api/admin/news/:id
func updateNews(c echo.Context) error {
	id := c.Param("id")
	var in newsInput
	if err := c.Bind(&in); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "некорректные данные"})
	}
	in.ImageURL = normalizeAssetURL(in.ImageURL)
	publishAt, err := parsePublishAt(in.PublishAt)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "некорректная дата публикации (ожидается формат RFC3339)"})
	}
	res, err := db.Exec(
		`UPDATE news SET image_url=$1, title=$2, short_desc=$3, full_desc=$4, publish_at=$5 WHERE id=$6`,
		in.ImageURL, in.Title, in.ShortDesc, in.FullDesc, publishAt, id,
	)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "не удалось обновить новость"})
	}
	if rows, _ := res.RowsAffected(); rows == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "новость не найдена"})
	}
	return c.JSON(http.StatusOK, echo.Map{"status": "ok"})
}

// PATCH /api/admin/news/:id/hidden
func setNewsHidden(c echo.Context) error {
	return setHiddenGeneric(c, "news")
}

// PATCH /api/admin/news/:id/postpone  body: {"publish_at": "2026-08-01T10:00:00+03:00"}
// Отложить публикацию новости на указанное время. Пустая строка publish_at - опубликовать сразу.
func postponeNews(c echo.Context) error {
	id := c.Param("id")
	var body struct {
		PublishAt string `json:"publish_at"`
	}
	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "некорректные данные"})
	}
	publishAt, err := parsePublishAt(body.PublishAt)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "некорректная дата (ожидается формат RFC3339)"})
	}
	res, err := db.Exec(`UPDATE news SET publish_at=$1 WHERE id=$2`, publishAt, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "не удалось отложить новость"})
	}
	if rows, _ := res.RowsAffected(); rows == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "новость не найдена"})
	}
	return c.JSON(http.StatusOK, echo.Map{"status": "ok"})
}

// DELETE /api/admin/news/:id
func deleteNews(c echo.Context) error {
	return deleteGeneric(c, "news", "новость не найдена")
}

// ---------- Выпадающие списки (высшая школа / курс) ----------

// GET /api/admin/dropdown-options?kind=school|course
func adminListDropdownOptions(c echo.Context) error {
	kind := c.QueryParam("kind")
	if kind != "school" && kind != "course" {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "параметр kind должен быть 'school' или 'course'"})
	}
	items, err := listDropdownOptions(kind, true)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "ошибка чтения списка"})
	}
	return c.JSON(http.StatusOK, items)
}

// POST /api/admin/dropdown-options
func createDropdownOption(c echo.Context) error {
	var o DropdownOption
	if err := c.Bind(&o); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "некорректные данные"})
	}
	if o.Kind != "school" && o.Kind != "course" {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "поле kind должно быть 'school' или 'course'"})
	}
	err := db.QueryRow(
		`INSERT INTO dropdown_options (kind, value, sort_order, active) VALUES ($1, $2, $3, $4) RETURNING id`,
		o.Kind, o.Value, o.SortOrder, o.Active,
	).Scan(&o.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "не удалось создать пункт списка"})
	}
	return c.JSON(http.StatusCreated, o)
}

// PUT /api/admin/dropdown-options/:id
func updateDropdownOption(c echo.Context) error {
	id := c.Param("id")
	var o DropdownOption
	if err := c.Bind(&o); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "некорректные данные"})
	}
	res, err := db.Exec(
		`UPDATE dropdown_options SET value=$1, sort_order=$2, active=$3 WHERE id=$4`,
		o.Value, o.SortOrder, o.Active, id,
	)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "не удалось обновить пункт списка"})
	}
	if n, _ := res.RowsAffected(); n == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "пункт списка не найден"})
	}
	return c.JSON(http.StatusOK, echo.Map{"status": "ok"})
}

// DELETE /api/admin/dropdown-options/:id
func deleteDropdownOption(c echo.Context) error {
	return deleteGeneric(c, "dropdown_options", "пункт списка не найден")
}

// ---------- Заявки ----------

// GET /api/admin/applications
func getApplications(c echo.Context) error {
	rows, err := db.Query(`SELECT id, full_name, school, course, vk_link, agreement, created_at
		FROM applications ORDER BY created_at DESC`)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "ошибка чтения заявок"})
	}
	defer rows.Close()

	items := []Application{}
	for rows.Next() {
		var a Application
		if err := rows.Scan(&a.ID, &a.FullName, &a.School, &a.Course, &a.VKLink, &a.Agreement, &a.CreatedAt); err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": "ошибка чтения строки"})
		}
		items = append(items, a)
	}
	return c.JSON(http.StatusOK, items)
}

// DELETE /api/admin/applications/:id
func deleteApplication(c echo.Context) error {
	return deleteGeneric(c, "applications", "заявка не найдена")
}

// ---------- Общие вспомогательные функции ----------

// setHiddenGeneric - PATCH .../:id/hidden {"hidden": true|false} для таблиц с колонкой hidden
func setHiddenGeneric(c echo.Context, table string) error {
	id := c.Param("id")
	if _, err := strconv.Atoi(id); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "некорректный id"})
	}
	var body struct {
		Hidden bool `json:"hidden"`
	}
	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "некорректные данные"})
	}
	res, err := db.Exec(`UPDATE `+table+` SET hidden=$1 WHERE id=$2`, body.Hidden, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "не удалось изменить видимость"})
	}
	if n, _ := res.RowsAffected(); n == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "запись не найдена"})
	}
	return c.JSON(http.StatusOK, echo.Map{"status": "ok"})
}

// deleteGeneric - DELETE .../:id для простых таблиц
func deleteGeneric(c echo.Context, table, notFoundMsg string) error {
	id := c.Param("id")
	if _, err := strconv.Atoi(id); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "некорректный id"})
	}
	res, err := db.Exec(`DELETE FROM `+table+` WHERE id=$1`, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "не удалось удалить запись"})
	}
	if n, _ := res.RowsAffected(); n == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": notFoundMsg})
	}
	return c.JSON(http.StatusOK, echo.Map{"status": "ok"})
}
