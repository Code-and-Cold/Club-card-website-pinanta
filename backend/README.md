# Сайт-визитка клуба (Go + Echo + Postgres)

Простой прототип: одна публичная страница (команда, новости, форма заявки) + админка для редактирования этих данных.

## Структура

```
main.go               - маршруты, запуск сервера, basic-auth для /api/admin/*
db.go                 - подключение к Postgres, авто-создание таблиц
models.go             - структуры TeamMember / News / Application
handlers_public.go    - GET /api/team, GET /api/news, POST /api/apply
handlers_admin.go     - CRUD для команды/новостей, просмотр заявок (под basic-auth)
static/index.html     - публичная страница
static/admin.html     - админка
schema.sql            - та же схема БД отдельным файлом (для справки)
```

## 1. Поднимаем Postgres в Docker

Здесь используется `docker run`, а не `docker compose`

**Базовый вариант (стандартный порт 5432):**

```bash
docker run -d \
  --name clubsite-postgres \
  -e POSTGRES_USER=clubuser \
  -e POSTGRES_PASSWORD=clubpass \
  -e POSTGRES_DB=clubdb \
  -p 5432:5432 \
  -v clubdata:/var/lib/postgresql/data \
  postgres:16
```

**Второй контейнер на нестандартном порту**. Здесь на хосте используется порт `5433`, внутри контейнера Postgres всё так же слушает `5432`:

```bash
docker run -d \
  --name clubsite-postgres-2 \
  -e POSTGRES_USER=clubuser \
  -e POSTGRES_PASSWORD=clubpass \
  -e POSTGRES_DB=clubdb \
  -p 5433:5432 \
  -v clubdata2:/var/lib/postgresql/data \
  postgres:16
```

Оба контейнера могут работать одновременно — у них разные имена (`clubsite-postgres` / `clubsite-postgres-2`), разные volume (`clubdata` / `clubdata2`) и разные порты на хосте (`5432` / `5433`).

**Полезные команды:**

```bash
docker ps                        # какие контейнеры сейчас работают
docker start clubsite-postgres   # запустить ранее созданный, но остановленный контейнер
docker stop clubsite-postgres    # остановить
docker rm -f clubsite-postgres   # удалить контейнер (данные в volume останутся)
docker volume rm clubdata        # удалить volume, если нужно начать с чистой БД
```

Если удобнее compose — `docker-compose.yml` в репозитории тоже рабочий:
```bash
docker compose up -d      # если есть плагин compose
# или
docker-compose up -d      # если стоит старая отдельная утилита
```

## 2. Настраиваем переменные окружения

Скопируй `.env.example` в `.env` и поправь при необходимости (логин/пароль админки, строку подключения к БД). Если поднял второй Postgres на порту `5433`, укажи это в `DATABASE_URL`.

## 3. Собираем и запускаем сервер

```bash
go mod tidy
go build -o clubsite .
```

Запуск с БД на стандартном порту 5432:
```bash
DATABASE_URL=postgres://clubuser:clubpass@localhost:5432/clubdb?sslmode=disable \
ADMIN_LOGIN=admin ADMIN_PASSWORD=admin123 PORT=8080 \
./clubsite
```

Запуск с БД на порту 5433 (второй контейнер):
```bash
DATABASE_URL=postgres://clubuser:clubpass@localhost:5433/clubdb?sslmode=disable \
ADMIN_LOGIN=admin ADMIN_PASSWORD=admin123 PORT=8081 \
./clubsite
```
(порт сервера `PORT=8081` тоже стоит поменять, если хочешь запустить второй экземпляр сервера параллельно с первым)

## 4. Открываем в браузере

- http://localhost:8080/            — публичная страница
- http://localhost:8080/admin.html  — админка (логин/пароль по умолчанию: admin / admin123)

## API

Публичное:
- `GET /api/site` — одним запросом все данные для главной страницы: `hero`, `warmup`, `why_us` (+`why_us_cards`), `team` (+`team_members`), `benefits`, `quote`, `news_block` (+`news`), `apply_block`, `schools`, `courses`. Скрытые карточки/участники/новости и отложенные новости (publish_at в будущем) в ответ не попадают.
- `POST /api/apply` — заявка на вступление, body:
  ```json
  {
    "full_name": "Иванов Иван Иванович",
    "school": "Высшая школа экономики и менеджмента",
    "course": "1 курс",
    "vk_link": "https://vk.com/ivanov",
    "agreement": true
  }
  ```
  Валидация: `full_name` — только буквы/пробелы; `school`/`course` — должны совпадать со значением из активных `dropdown_options`; `vk_link` — ссылка на vk.com; `agreement` — обязательно `true`. При ошибке возвращается `400` с полями `error` (код) и `message` (текст для показа пользователю на русском). При успехе — `201` и `{"status":"ok","message":"Заявка успешно отправлена!"}`. Заявка сохраняется в Postgres и (если настроено) асинхронно синхронизируется со строкой в xlsx-таблице на Яндекс Диске — см. раздел ниже.

Админское (Basic Auth):
- `GET/PUT /api/admin/content/:key` — простые блоки: `hero`, `warmup`, `why_us`, `team`, `benefits`, `quote`, `news_block`, `apply_block`. Тело PUT: `{"content": {...}}`, содержимое произвольное (например для `hero`: `{"title":"...","subtitle":"...","photo_url":"..."}`).
- `GET/POST/PUT/DELETE /api/admin/team[/:id]`, `PATCH /api/admin/team/:id/hidden` `{"hidden": true}`
- `GET/POST/PUT/DELETE /api/admin/why-us-cards[/:id]`, `PATCH .../:id/hidden`
- `GET/POST/PUT/DELETE /api/admin/benefits[/:id]`, `PATCH .../:id/hidden`
- `GET/POST/PUT/DELETE /api/admin/news[/:id]`, `PATCH .../:id/hidden`, `PATCH .../:id/postpone` `{"publish_at": "2026-08-01T10:00:00+03:00"}` — отложить публикацию (пустая строка = опубликовать сразу)
- `GET/POST/PUT/DELETE /api/admin/dropdown-options[/:id]` — пункты выпадающих списков формы заявки, `GET` принимает `?kind=school` или `?kind=course`
- `GET /api/admin/applications`, `DELETE /api/admin/applications/:id`

### Синхронизация заявок с таблицей на Яндекс Диске

Реализована в `yandex.go`: при успешной заявке сервер скачивает xlsx-файл с Яндекс Диска, дописывает строку (ФИО, школа, курс, ссылка ВК, согласие, дата) и загружает файл обратно.

Важно: публичная ссылка вида `https://disk.yandex.ru/i/...` сама по себе не позволяет что-либо в файл дозаписывать — это просто "посмотреть/скачать". Чтобы бек мог писать в таблицу, нужно:
1. Получить OAuth-токен приложения на диске, которому принадлежит файл: https://oauth.yandex.ru (права `cloud_api:disk.write`).
2. Убедиться, что таблица (тот самый файл) лежит именно на диске владельца этого токена, и узнать её путь на диске (например `/Заявки/Заявки клуба.xlsx`).
3. Указать в `.env`: `YANDEX_DISK_TOKEN` и `YANDEX_TABLE_PATH`.

Если переменные не заданы — синхронизация просто отключена, заявки всё равно сохраняются в Postgres (это основной источник данных).

## Что упрощено намеренно

- Авторизация админки — HTTP Basic Auth с логином/паролем из переменных окружения (без сессий, JWT, ролей).
- Фото/иконки — просто URL-строка (загрузки файлов на сервер нет; ожидается, что картинки заливаются куда-то отдельно, например тот же Яндекс Диск/CDN, а в блок передаётся ссылка).
- Нет пагинации и поиска.
- Списки "высшая школа" и "курс" в форме заявки хранятся в таблице `dropdown_options` и управляются через `/api/admin/dropdown-options` — перед первым использованием формы их нужно один раз заполнить (конкретные значения из макета в ТЗ не были доступны на момент разработки бэка).
- Скрытие карточек/участников/новостей — soft-флаг `hidden`, запись не удаляется из БД.
- "Отложить новость" — поле `publish_at`; новость с датой в будущем не отдаётся в публичном `/api/site`, но видна в админском `/api/admin/news`.

## Не забыть перед запуском

```bash
go mod tidy   # подтянет новую зависимость github.com/xuri/excelize/v2
```

Этого достаточно для тестового прототипа; для продакшена стоит добавить нормальную аутентификацию, HTTPS, валидацию и, возможно, миграции через отдельный инструмент (goose/golang-migrate).
# empty
