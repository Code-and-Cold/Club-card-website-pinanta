# 🔐 Админский API

Базовый URL: `http://your-domain.com`

## Авторизация

Все админские эндпоинты требуют **HTTP Basic Auth**.

**Заголовок:**
    Authorization: Basic base64(login:password)

**Примеры запросов:**

    # Вариант 1: явный заголовок (base64 от "admin:admin123" = "YWRtaW46YWRtaW4xMjM=")
    curl -H "Authorization: Basic YWRtaW46YWRtaW4xMjM=" \
      http://localhost:8080/api/admin/content/hero

    # Вариант 2: флаг -u (проще)
    curl -u admin:admin123 http://localhost:8080/api/admin/content/hero

---

## 1. Управление текстовыми блоками сайта

Доступные ключи блоков: `hero`, `warmup`, `why_us`, `team`, `benefits`, `quote`, `news_block`, `apply_block`

### GET /api/admin/content/:key

Получить содержимое блока.

    curl -u admin:admin123 http://localhost:8080/api/admin/content/hero

**Ответ (200 OK):**

    {
      "key": "hero",
      "content": {
        "title": "Код и Холод",
        "subtitle": "Клуб разработчиков",
        "photo_url": "https://example.com/hero.jpg"
      },
      "updated_at": "2026-08-01T10:00:00+03:00"
    }

### PUT /api/admin/content/:key

Обновить содержимое блока.

    curl -X PUT -u admin:admin123 \
      -H "Content-Type: application/json" \
      -d '{"content": {"title": "Новый заголовок", "subtitle": "Новый подзаголовок", "photo_url": "https://example.com/new.jpg"}}' \
      http://localhost:8080/api/admin/content/hero

**Ответ (200 OK):**

    {
      "key": "hero",
      "content": {
        "title": "Новый заголовок",
        "subtitle": "Новый подзаголовок",
        "photo_url": "https://example.com/new.jpg"
      },
      "updated_at": "2026-08-03T15:30:00+03:00"
    }

**Структура content для разных блоков:**

**Hero** (1 блок):
    { "title": "Заголовок", "subtitle": "Подзаголовок", "photo_url": "https://..." }

**Warmup** (2 блок — разогрев):
    { "text": "Текст лозунга" }

**Why Us** (3 блок — заголовок "почему у нас классно"):
    { "title": "Почему у нас классно" }

**Team** (4 блок — заголовок "наша команда"):
    { "title": "Наша команда" }

**Benefits** (5 блок — заголовок "что ты получишь"):
    { "title": "Что ты получишь" }

**Quote** (6 блок — цитата):
    { "text": "Вдохновляющая цитата" }

**News Block** (7 блок — заголовок "новости"):
    { "title": "Новости" }

**Apply Block** (8 блок — форма заявки):
    { "title": "Подать заявку", "description": "Описание блока", "photo_url": "https://..." }

---

## 2. Управление командой (блок 4 — карточки участников)

### GET /api/admin/team

Получить список участников.

    curl -u admin:admin123 http://localhost:8080/api/admin/team

**Ответ (200 OK):**

    [
      {
        "id": 1,
        "name": "Иван Иванов",
        "role": "Руководитель клуба",
        "photo_url": "https://example.com/ivan.jpg",
        "hover_info": "Люблю программировать",
        "description": "Руководитель клуба с 5-летним опытом",
        "contacts": "@ivanov | ivan@example.com",
        "hidden": false,
        "sort_order": 1,
        "created_at": "2026-08-01T10:00:00+03:00"
      }
    ]

### POST /api/admin/team

Добавить нового участника.

    curl -X POST -u admin:admin123 \
      -H "Content-Type: application/json" \
      -d '{"name": "Пётр Петров", "role": "Разработчик", "photo_url": "https://example.com/petr.jpg", "hover_info": "Инфо при наведении", "description": "Описание", "contacts": "@petrov", "sort_order": 2}' \
      http://localhost:8080/api/admin/team

**Ответ (201 Created):**

    {
      "id": 2,
      "name": "Пётр Петров",
      "role": "Разработчик",
      "photo_url": "https://example.com/petr.jpg",
      "hover_info": "Инфо при наведении",
      "description": "Описание",
      "contacts": "@petrov",
      "hidden": false,
      "sort_order": 2,
      "created_at": "2026-08-03T15:30:00+03:00"
    }

### PUT /api/admin/team/:id

Обновить участника.

    curl -X PUT -u admin:admin123 \
      -H "Content-Type: application/json" \
      -d '{"name": "Пётр Петров (обновлённый)", "role": "Старший разработчик", "photo_url": "https://example.com/petr-new.jpg", "hover_info": "Новая инфа", "description": "Новое описание", "contacts": "@petrov_new", "sort_order": 1}' \
      http://localhost:8080/api/admin/team/2

**Ответ (200 OK):** обновлённый объект участника.

### DELETE /api/admin/team/:id

Удалить участника.

    curl -X DELETE -u admin:admin123 http://localhost:8080/api/admin/team/2

**Ответ:** 204 No Content

### PATCH /api/admin/team/:id/hidden

Скрыть / показать участника.

    curl -X PATCH -u admin:admin123 \
      -H "Content-Type: application/json" \
      -d '{"hidden": true}' \
      http://localhost:8080/api/admin/team/1/hidden

**Ответ (200 OK):** объект участника с обновлённым флагом `hidden`.

> Скрытые участники не отдаются в публичном `GET /api/site`, но остаются в БД.

---

## 3. Управление карточками "Почему у нас классно" (блок 3)

### GET /api/admin/why-us-cards

    curl -u admin:admin123 http://localhost:8080/api/admin/why-us-cards

**Ответ (200 OK):**

    [
      {
        "id": 1,
        "icon_url": "https://example.com/mentors.png",
        "title": "Опытные менторы",
        "description": "Наши менторы помогут вам в развитии",
        "hidden": false,
        "sort_order": 1
      }
    ]

### POST /api/admin/why-us-cards

    curl -X POST -u admin:admin123 \
      -H "Content-Type: application/json" \
      -d '{"icon_url": "https://example.com/projects.png", "title": "Реальные проекты", "description": "Работа над настоящими проектами", "sort_order": 2}' \
      http://localhost:8080/api/admin/why-us-cards

**Ответ (201 Created):** объект созданной карточки.

### PUT /api/admin/why-us-cards/:id

Обновить карточку (поля те же, что и при создании).

    curl -X PUT -u admin:admin123 \
      -H "Content-Type: application/json" \
      -d '{"icon_url": "https://example.com/new.png", "title": "Обновлённый заголовок", "description": "Обновлённое описание", "sort_order": 1}' \
      http://localhost:8080/api/admin/why-us-cards/1

### DELETE /api/admin/why-us-cards/:id

    curl -X DELETE -u admin:admin123 http://localhost:8080/api/admin/why-us-cards/1

**Ответ:** 204 No Content

### PATCH /api/admin/why-us-cards/:id/hidden

    curl -X PATCH -u admin:admin123 \
      -H "Content-Type: application/json" \
      -d '{"hidden": true}' \
      http://localhost:8080/api/admin/why-us-cards/1/hidden

---

## 4. Управление преимуществами (блок 5 — "что ты получишь")

### GET /api/admin/benefits

    curl -u admin:admin123 http://localhost:8080/api/admin/benefits

**Ответ (200 OK):**

    [
      {
        "id": 1,
        "icon_url": "https://example.com/skills.png",
        "title": "Новые навыки",
        "description": "Освоишь современные технологии",
        "hidden": false,
        "sort_order": 1
      }
    ]

### POST /api/admin/benefits

    curl -X POST -u admin:admin123 \
      -H "Content-Type: application/json" \
      -d '{"icon_url": "https://example.com/networking.png", "title": "Нетворкинг", "description": "Знакомства с профессионалами", "sort_order": 2}' \
      http://localhost:8080/api/admin/benefits

**Ответ (201 Created):** объект созданного преимущества.

### PUT /api/admin/benefits/:id

    curl -X PUT -u admin:admin123 \
      -H "Content-Type: application/json" \
      -d '{"icon_url": "https://example.com/new.png", "title": "Новый заголовок", "description": "Новое описание", "sort_order": 1}' \
      http://localhost:8080/api/admin/benefits/1

### DELETE /api/admin/benefits/:id

    curl -X DELETE -u admin:admin123 http://localhost:8080/api/admin/benefits/1

**Ответ:** 204 No Content

### PATCH /api/admin/benefits/:id/hidden

    curl -X PATCH -u admin:admin123 \
      -H "Content-Type: application/json" \
      -d '{"hidden": true}' \
      http://localhost:8080/api/admin/benefits/1/hidden

---

## 5. Управление новостями (блок 7)

### GET /api/admin/news

Получить список новостей (**включая** скрытые и отложенные — в отличие от публичного API).

    curl -u admin:admin123 http://localhost:8080/api/admin/news

**Ответ (200 OK):**

    [
      {
        "id": 1,
        "image_url": "https://example.com/news1.jpg",
        "title": "Открыт набор в клуб",
        "short_desc": "Мы набираем новых участников",
        "full_desc": "Полное описание новости с деталями...",
        "hidden": false,
        "publish_at": null,
        "created_at": "2026-08-01T10:00:00+03:00"
      },
      {
        "id": 2,
        "image_url": "https://example.com/news2.jpg",
        "title": "Предстоящее мероприятие",
        "short_desc": "Скоро будет встреча",
        "full_desc": "Подробности о мероприятии...",
        "hidden": false,
        "publish_at": "2026-08-10T18:00:00+03:00",
        "created_at": "2026-08-03T15:30:00+03:00"
      }
    ]

### POST /api/admin/news

Создать новость. Дата публикации (`created_at`) ставится автоматически.

    curl -X POST -u admin:admin123 \
      -H "Content-Type: application/json" \
      -d '{"image_url": "https://example.com/news3.jpg", "title": "Новая новость", "short_desc": "Краткое описание", "full_desc": "Полное описание"}' \
      http://localhost:8080/api/admin/news

Создать новость с отложенной публикацией:

    curl -X POST -u admin:admin123 \
      -H "Content-Type: application/json" \
      -d '{"image_url": "https://example.com/news4.jpg", "title": "Будущая новость", "short_desc": "Появится позже", "full_desc": "Полное описание", "publish_at": "2026-08-15T10:00:00+03:00"}' \
      http://localhost:8080/api/admin/news

**Ответ (201 Created):** объект созданной новости.

### PUT /api/admin/news/:id

Обновить новость.

    curl -X PUT -u admin:admin123 \
      -H "Content-Type: application/json" \
      -d '{"image_url": "https://example.com/news1-upd.jpg", "title": "Обновлённая новость", "short_desc": "Обновлённое краткое описание", "full_desc": "Обновлённое полное описание"}' \
      http://localhost:8080/api/admin/news/1

### DELETE /api/admin/news/:id

    curl -X DELETE -u admin:admin123 http://localhost:8080/api/admin/news/1

**Ответ:** 204 No Content

### PATCH /api/admin/news/:id/hidden

Скрыть / показать новость.

    curl -X PATCH -u admin:admin123 \
      -H "Content-Type: application/json" \
      -d '{"hidden": true}' \
      http://localhost:8080/api/admin/news/1/hidden

### PATCH /api/admin/news/:id/postpone

Отложить публикацию новости (новость появится в публичном API только после указанной даты).

    curl -X PATCH -u admin:admin123 \
      -H "Content-Type: application/json" \
      -d '{"publish_at": "2026-08-10T18:00:00+03:00"}' \
      http://localhost:8080/api/admin/news/1/postpone

**Ответ (200 OK):** объект новости с обновлённым `publish_at`.

Отменить отложенную публикацию (опубликовать сразу):

    curl -X PATCH -u admin:admin123 \
      -H "Content-Type: application/json" \
      -d '{"publish_at": ""}' \
      http://localhost:8080/api/admin/news/1/postpone

> Отложенные новости (publish_at в будущем) не отдаются в публичном `GET /api/site`, но видны в админке.

---

## 6. Управление выпадающими списками (школа / курс для формы заявки)

### GET /api/admin/dropdown-options

Получить список опций. Обязательный query-параметр `kind` = `school` или `course`.

    curl -u admin:admin123 "http://localhost:8080/api/admin/dropdown-options?kind=school"
    curl -u admin:admin123 "http://localhost:8080/api/admin/dropdown-options?kind=course"

**Ответ (200 OK):**

    [
      {
        "id": 1,
        "kind": "school",
        "value": "Высшая школа экономики и менеджмента",
        "sort_order": 1,
        "active": true
      }
    ]

### POST /api/admin/dropdown-options

Добавить новую опцию.

    curl -X POST -u admin:admin123 \
      -H "Content-Type: application/json" \
      -d '{"kind": "school", "value": "Новая высшая школа", "sort_order": 3, "active": true}' \
      http://localhost:8080/api/admin/dropdown-options

    # Для курса:
    curl -X POST -u admin:admin123 \
      -H "Content-Type: application/json" \
      -d '{"kind": "course", "value": "3 курс", "sort_order": 3, "active": true}' \
      http://localhost:8080/api/admin/dropdown-options

**Ответ (201 Created):** объект созданной опции.

### PUT /api/admin/dropdown-options/:id

Обновить опцию.

    curl -X PUT -u admin:admin123 \
      -H "Content-Type: application/json" \
      -d '{"kind": "school", "value": "Обновлённое название", "sort_order": 1, "active": true}' \
      http://localhost:8080/api/admin/dropdown-options/3

### DELETE /api/admin/dropdown-options/:id

    curl -X DELETE -u admin:admin123 http://localhost:8080/api/admin/dropdown-options/3

**Ответ:** 204 No Content

> Неактивные опции (active = false) не отдаются в публичном API и не проходят валидацию при отправке заявки.

---

## 7. Управление заявками

### GET /api/admin/applications

Получить список всех заявок (по убыванию даты).

    curl -u admin:admin123 http://localhost:8080/api/admin/applications

**Ответ (200 OK):**

    [
      {
        "id": 1,
        "full_name": "Иванов Иван Иванович",
        "school": "Высшая школа экономики и менеджмента",
        "course": "1 курс",
        "vk_link": "https://vk.com/ivanov",
        "agreement": true,
        "created_at": "2026-08-03T10:00:00+03:00"
      }
    ]

### DELETE /api/admin/applications/:id

Удалить заявку.

    curl -X DELETE -u admin:admin123 http://localhost:8080/api/admin/applications/1

**Ответ:** 204 No Content

---

## Коды ответов HTTP

| Код | Описание |
|-----|----------|
| 200 | Успех |
| 201 | Ресурс успешно создан |
| 204 | Успех, без содержимого (обычно для DELETE) |
| 400 | Ошибка валидации запроса |
| 401 | Не авторизован (неверный логин/пароль) |
| 404 | Ресурс не найден |
| 500 | Внутренняя ошибка сервера |