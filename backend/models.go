package main

import "time"

// ---------- Заявка на вступление ----------

// Application - заявка на вступление в клуб
type Application struct {
	ID        int       `json:"id"`
	FullName  string    `json:"full_name"`
	School    string    `json:"school"`
	Course    string    `json:"course"`
	VKLink    string    `json:"vk_link"`
	Agreement bool      `json:"agreement"`
	CreatedAt time.Time `json:"created_at"`
}

// ApplyRequest - тело POST /api/apply
type ApplyRequest struct {
	FullName  string `json:"full_name"`
	School    string `json:"school"`
	Course    string `json:"course"`
	VKLink    string `json:"vk_link"`
	Agreement bool   `json:"agreement"`
}

// DropdownOption - пункт выпадающего списка (высшая школа / курс)
type DropdownOption struct {
	ID        int    `json:"id"`
	Kind      string `json:"kind"` // "school" | "course"
	Value     string `json:"value"`
	SortOrder int    `json:"sort_order"`
	Active    bool   `json:"active"`
}

// ---------- Команда ----------

// TeamMember - участник команды (блок "наша команда")
type TeamMember struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Role        string    `json:"role"`
	PhotoURL    string    `json:"photo_url"`
	HoverInfo   string    `json:"hover_info"` // необязательное поле, показывается при наведении
	Description string    `json:"description"`
	Contacts    string    `json:"contacts"`
	Hidden      bool      `json:"hidden"`
	SortOrder   int       `json:"sort_order"`
	CreatedAt   time.Time `json:"created_at"`
}

// ---------- Новости ----------

// News - карточка новости (блок "новости")
type News struct {
	ID        int        `json:"id"`
	ImageURL  string     `json:"image_url"`
	Title     string     `json:"title"`
	ShortDesc string     `json:"short_desc"`
	FullDesc  string     `json:"full_desc"`
	Hidden    bool       `json:"hidden"`
	PublishAt *time.Time `json:"publish_at"` // если в будущем - новость "отложена"
	CreatedAt time.Time  `json:"created_at"`
}

// ---------- Карточки "почему у нас классно" ----------

type WhyUsCard struct {
	ID          int    `json:"id"`
	IconURL     string `json:"icon_url"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Hidden      bool   `json:"hidden"`
	SortOrder   int    `json:"sort_order"`
}

// ---------- Карточки "что ты получишь" ----------

type Benefit struct {
	ID          int    `json:"id"`
	IconURL     string `json:"icon_url"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Hidden      bool   `json:"hidden"`
	SortOrder   int    `json:"sort_order"`
}

// ---------- Простые текстовые/JSON-блоки сайта ----------
// hero, warmup, why_us, team, benefits, quote, news_block, apply_block
// хранятся как произвольный JSON в одной таблице site_blocks

type SiteBlock struct {
	Key       string    `json:"key"`
	Content   JSONMap   `json:"content"`
	UpdatedAt time.Time `json:"updated_at"`
}

// JSONMap - обёртка для хранения произвольного JSON в Postgres (jsonb)
type JSONMap map[string]interface{}
