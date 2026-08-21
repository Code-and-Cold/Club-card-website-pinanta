package main

import (
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
)

var (
	fullNameRe = regexp.MustCompile(`^[A-Za-zА-Яа-яЁё]+([ \-][A-Za-zА-Яа-яЁё]+)*$`)
	vkLinkRe   = regexp.MustCompile(`(?i)https?://(?:www\.)?(?:vk\.com|vk\.ru)/[^\s]+`)
)

// GET /api/site - агрегированные данные для главной страницы (все блоки одним запросом)
func getSite(c echo.Context) error {
	resp := echo.Map{}

	blocks, err := loadAllSiteBlocks()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "ошибка чтения блоков сайта"})
	}
	for key, content := range blocks {
		resp[key] = content
	}

	whyUsCards, err := listWhyUsCards(false)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "ошибка чтения карточек"})
	}
	resp["why_us_cards"] = whyUsCards

	team, err := listTeamMembers(false)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "ошибка чтения команды"})
	}
	resp["team_members"] = team

	benefits, err := listBenefits(false)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "ошибка чтения преимуществ"})
	}
	resp["benefits_items"] = benefits

	news, err := listNews(false)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "ошибка чтения новостей"})
	}
	resp["news"] = news

	schools, err := listDropdownOptions("school", false)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "ошибка чтения списка школ"})
	}
	resp["schools"] = schools

	courses, err := listDropdownOptions("course", false)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "ошибка чтения списка курсов"})
	}
	resp["courses"] = courses

	return c.JSON(http.StatusOK, resp)
}

// POST /api/apply - заявка на вступление в клуб
func postApply(c echo.Context) error {
	var in ApplyRequest
	if err := c.Bind(&in); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "формат данных не подходит", "message": "Не удалось прочитать данные формы. Проверьте, что все поля заполнены корректно."})
	}

	in.FullName = strings.TrimSpace(in.FullName)
	in.School = strings.TrimSpace(in.School)
	in.Course = strings.TrimSpace(in.Course)
	in.VKLink = strings.TrimSpace(in.VKLink)

	if in.FullName == "" || !fullNameRe.MatchString(in.FullName) {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error":   "invalid_full_name",
			"message": "ФИО должно содержать только буквы и пробелы (без цифр и спецсимволов).",
		})
	}
	if len([]rune(in.FullName)) < 5 {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error":   "invalid_full_name",
			"message": "Укажите, пожалуйста, ФИО полностью.",
		})
	}

	if in.School == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error":   "invalid_school",
			"message": "Выберите высшую школу из списка.",
		})
	}
	schoolOK, err := dropdownOptionExists("school", in.School)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "внутренняя ошибка", "message": "Форма не отправлена, попробуйте ещё раз позже."})
	}
	if !schoolOK {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error":   "invalid_school",
			"message": "Выбранная высшая школа недоступна. Обновите страницу и выберите значение из списка.",
		})
	}

	if in.Course == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error":   "invalid_course",
			"message": "Выберите курс из списка.",
		})
	}
	courseOK, err := dropdownOptionExists("course", in.Course)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "внутренняя ошибка", "message": "Форма не отправлена, попробуйте ещё раз позже."})
	}
	if !courseOK {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error":   "invalid_course",
			"message": "Выбранный курс недоступен. Обновите страницу и выберите значение из списка.",
		})
	}

	if in.VKLink == "" || !vkLinkRe.MatchString(in.VKLink) {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error":   "invalid_vk_link",
			"message": "Укажите корректную ссылку на страницу ВКонтакте (https://vk.com/...).",
		})
	}

	if !in.Agreement {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error":   "agreement_required",
			"message": "Необходимо подтвердить согласие для отправки формы.",
		})
	}

	app := Application{
		FullName:  in.FullName,
		School:    in.School,
		Course:    in.Course,
		VKLink:    in.VKLink,
		Agreement: true,
		CreatedAt: time.Now(),
	}

	err = db.QueryRow(
		`INSERT INTO applications (full_name, school, course, vk_link, agreement)
		 VALUES ($1, $2, $3, $4, $5) RETURNING id, created_at`,
		app.FullName, app.School, app.Course, app.VKLink, app.Agreement,
	).Scan(&app.ID, &app.CreatedAt)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error":   "save_failed",
			"message": "Не удалось сохранить заявку. Форма не отправлена, попробуйте ещё раз позже.",
		})
	}

	// Синхронизация с таблицей на Яндекс Диске - в фоне, чтобы не задерживать ответ пользователю.
	// Заявка уже надёжно сохранена в БД, даже если Яндекс Диск временно недоступен.
	syncApplicationToYandexAsync(app)

	return c.JSON(http.StatusCreated, echo.Map{
		"id":      app.ID,
		"status":  "ok",
		"message": "Заявка успешно отправлена!",
	})
}
