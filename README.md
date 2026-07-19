# Сайт-визитка клуба

Сайт пиньята для клуба программистов. Содержит одну публичную страницу (команда, новости, форма заявки) + админку для редактирования этих данных.
Фронт залит на GitHub pages (TODO: Здесь ссылка).

## Стек

Backend:

- Go: язык программирования.
- Echo: быстрый и простой Go web framework
- Postgres 16: современная БД.

Frontend:

- Vue: JS framework для создания пользовательских интерфейсов.
  Пока используется вместе с набором инструментов Vite.

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

## Быстрый старт

TODO: полное развертывание back+front в dev режиме (docker compose?)

## Развертывание backend

### 1. Поднимаем Postgres в Docker

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

### 2. Настраиваем переменные окружения

Скопируй `.env.example` в `.env` и поправь при необходимости (логин/пароль админки, строку подключения к БД). Если поднял второй Postgres на порту `5433`, укажи это в `DATABASE_URL`.

### 3. Собираем и запускаем сервер

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

### 4. Открываем в браузере

- http://localhost:8080/            — публичная страница
- http://localhost:8080/admin.html  — админка (логин/пароль по умолчанию: admin / admin123)

## API

Публичное:
- `GET /api/team` — список команды
- `GET /api/news` — список новостей
- `POST /api/apply` — заявка на вступление, body: `{"name": "...", "email": "..."}`

Админское (Basic Auth):
- `POST/PUT/DELETE /api/admin/team[/:id]`
- `POST/PUT/DELETE /api/admin/news[/:id]`
- `GET /api/admin/applications`
- `DELETE /api/admin/applications/:id`

## Что упрощено намеренно

- Авторизация админки — HTTP Basic Auth с логином/паролем из переменных окружения (без сессий, JWT, ролей).
- Верификации почты в заявке нет.
- Фото участников — просто URL-строка (загрузки файлов нет).
- Нет пагинации, поиска, валидации помимо базовой проверки "поле не пустое".

Этого достаточно для тестового прототипа; для продакшена стоит добавить нормальную аутентификацию, HTTPS, валидацию и, возможно, миграции через отдельный инструмент (goose/golang-migrate).
# empty

## Развертывание frontend

Деплой будем делать на GitHub pages.

```bash
cd frontend
# Устанавливаем зависимости:
npm install
# Build:
npm run predeploy
# Deploy на GitHub pages (обновляет ветку gh-pages)
npm run deploy
```
