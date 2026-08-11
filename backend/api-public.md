# Публичный API

Базовый URL: `http://your-domain.com`

## 1. Получить данные сайта

### `GET /api/site`

Возвращает все данные для главной страницы одним запросом.

**Ответ (200 OK):**
```json
{
  "hero": {
    "title": "Код и Холод",
    "subtitle": "Клуб разработчиков",
    "photo_url": "https://example.com/hero.jpg"
  },
  "warmup": {
    "text": "Присоединяйся к нашему клубу разработчиков!"
  },
  "why_us": {
    "title": "Почему у нас классно"
  },
  "why_us_cards": [
    {
      "id": 1,
      "icon_url": "https://example.com/mentors.png",
      "title": "Опытные менторы",
      "description": "Наши менторы помогут вам в развитии"
    },
    {
      "id": 2,
      "icon_url": "https://example.com/projects.png",
      "title": "Реальные проекты",
      "description": "Работа над настоящими проектами"
    }
  ],
  "team": {
    "title": "Наша команда"
  },
  "team_members": [
    {
      "id": 1,
      "name": "Иван Иванов",
      "role": "Руководитель клуба",
      "photo_url": "https://example.com/ivan.jpg",
      "hover_info": "Люблю программировать",
      "description": "Руководитель клуба с 5-летним опытом",
      "contacts": "@ivanov | ivan@example.com"
    }
  ],
  "benefits": {
    "title": "Что ты получишь"
  },
  "benefits_items": [
    {
      "id": 1,
      "icon_url": "https://example.com/skills.png",
      "title": "Новые навыки",
      "description": "Освоишь современные технологии"
    }
  ],
  "quote": {
    "text": "Программирование - это искусство"
  },
  "news_block": {
    "title": "Новости"
  },
  "news": [
    {
      "id": 1,
      "image_url": "https://example.com/news1.jpg",
      "title": "Открыт набор в клуб",
      "short_desc": "Мы набираем новых участников",
      "full_desc": "Полное описание новости с деталями...",
      "created_at": "2026-08-01T10:00:00+03:00"
    }
  ],
  "apply_block": {
    "title": "Подать заявку",
    "description": "Присоединяйся к нашему клубу!",
    "photo_url": "https://example.com/apply.jpg"
  },
  "schools": [
    {
      "id": 1,
      "value": "Высшая школа экономики и менеджмента",
      "sort_order": 1,
      "active": true
    },
    {
      "id": 2,
      "value": "Институт компьютерных технологий",
      "sort_order": 2,
      "active": true
    }
  ],
  "courses": [
    {
      "id": 1,
      "value": "1 курс",
      "sort_order": 1,
      "active": true
    },
    {
      "id": 2,
      "value": "2 курс",
      "sort_order": 2,
      "active": true
    }
  ]
}
```