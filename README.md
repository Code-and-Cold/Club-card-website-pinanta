
[![Node.js CI](https://github.com/Code-and-Cold/Club-card-website-pinanta/actions/workflows/node.js.yml/badge.svg)](https://github.com/Code-and-Cold/Club-card-website-pinanta/actions/workflows/node.js.yml)
[![Vue](https://img.shields.io/badge/vue-3.x-4FC08D?logo=vue.js)](https://vuejs.org/)
[![Vite](https://img.shields.io/badge/vite-6.x-646CFF?logo=vite)](https://vitejs.dev/)
[![Go](https://img.shields.io/badge/go-1.26.5-00ADD8?logo=go)](https://golang.org/)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

# Сайт-визитка клуба программистов "Код и Холод"

Образовательный проект для студентов, желающих получить практику в области разработки программного обеспечения.

В данном репозитории находится SPA сайт-визитка клуба со всеми зависимостями.
Проект находится в разработке, текущий статус можно узнать:

- На хостинге [GitHub pages](https://code-and-cold.github.io/Club-card-website-pinanta/).
- В макете [Figma](https://www.figma.com/design/QO5qKvgzRQCtMwVDZXz4md/%D0%9A%D0%BE%D0%B4-%D0%B8-%D0%A5%D0%BE%D0%BB%D0%BE%D0%B4?node-id=52-2).
- В системе управления проектами [Битрикс 24](https://code-and-cold.bitrix24.ru).

# Требования

Для работы проекта необходимы следующие компоненты:

- [Node.js 24+](https://nodejs.org/en/download)
- [Go 1.26.5+](https://go.dev/doc/install)
- [Docker 26.1.5+](https://docs.docker.com/engine/install/)

# Быстрый старт

## Полный деплой

Развертывание проекта через Docker Compose:

```
cp .env.dev .env
docker compose up -d
```

## Фронтенд

Проект является одностраничником (SPA) на Node.js реализованным с применением JS фреймворка Vue со стандартным композитором Vite.

Запуск сайта без БД и API:

```
cd frontend
npm install
npm run dev
```

Сайт доступен по адресу http://localhost:5173/Club-card-website-pinanta/. Порт может быть другим, если исходный занят.

## Бекенд

Работа с базой данных реализована через API, написанный на Go.

Запуск БД, API и панели администратора:

```
docker run -d --name clubsite-postgres \
  -e POSTGRES_USER=clubuser \
  -e POSTGRES_PASSWORD=clubpass \
  -e POSTGRES_DB=clubdb \
  -p 5432:5432 postgres:17-alpine

cd backend
go run . -env=../.env.dev
```

# Сценарий разработки

Типичный день frontend разработчика:

1. Запустим сайт для быстрой обратной связи, тесты для реализации TDD и защиты от регрессии, откроем фигму, битрикс и вк:

```
cd frontend
npm run dev
npm run test:system
```

2. Выберем задачу и решим её (опционально через TDD, см. тестирование в проекте):

```
git checkout -b feat/frontend/very-important-feature
git add -A
git commit -m "feat(frontend): very important section of right-most pixel"
```

3. Вовремя вспомним о том, что линтинг и форматирование при CI/CD не пройдут,
а кто-то очень ленивый не настроил `pre-commit-hooks` (TODO: настроить `pre-commit-hooks`):

```
npm run lint && npm run format
npm run test:ci
```

4. Сохраним изменения, корректируя историю при необходимости:

```
git add -A
git commit --amend --no-edit

git checkout develop
git rebase feat/frontend/very-important-feature
git push
```

4. (Опционально) Выложим наше творение ~~в даркнет~~ на GitHub Pages:

```
npm run deploy
```

TODO: Пнуть backend-ра для получения глубинного знания и пополнения инструкций.

# Дополнительные материалы (WIP)

- Принципы Frontend разработки в проекте
- Принципы Backend разработки в проекте
- Спецификация API
- Тестирование в проекте
- Методики управление проектом
- Принципы обеспечения безопасности
