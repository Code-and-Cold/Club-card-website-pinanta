package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var db *sql.DB

// connectDB открывает соединение с Postgres и создаёт таблицы, если их нет.
func connectDB() *sql.DB {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		// Значения по умолчанию для локальной разработки
		dsn = "postgres://clubuser:clubpass@localhost:5432/clubdb?sslmode=disable"
	}

	conn, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatalf("не удалось открыть соединение с БД: %v", err)
	}

	if err := conn.Ping(); err != nil {
		log.Fatalf("не удалось подключиться к БД: %v", err)
	}

	if err := runMigrations(conn); err != nil {
		log.Fatalf("не удалось применить миграции: %v", err)
	}

	log.Println("подключение к БД установлено, таблицы готовы")
	return conn
}

func runMigrations(conn *sql.DB) error {
	stmts := []string{
		// ---------- Команда ----------
		`CREATE TABLE IF NOT EXISTS team_members (
			id            SERIAL PRIMARY KEY,
			name          TEXT NOT NULL,
			role          TEXT NOT NULL DEFAULT '',
			photo_url     TEXT NOT NULL DEFAULT '',
			hover_info    TEXT NOT NULL DEFAULT '',
			description   TEXT NOT NULL DEFAULT '',
			contacts      TEXT NOT NULL DEFAULT '',
			hidden        BOOLEAN NOT NULL DEFAULT false,
			sort_order    INTEGER NOT NULL DEFAULT 0,
			created_at    TIMESTAMPTZ NOT NULL DEFAULT now()
		)`,
		`ALTER TABLE team_members ADD COLUMN IF NOT EXISTS hover_info TEXT NOT NULL DEFAULT ''`,
		`ALTER TABLE team_members ADD COLUMN IF NOT EXISTS description TEXT NOT NULL DEFAULT ''`,
		`ALTER TABLE team_members ADD COLUMN IF NOT EXISTS contacts TEXT NOT NULL DEFAULT ''`,
		`ALTER TABLE team_members ADD COLUMN IF NOT EXISTS hidden BOOLEAN NOT NULL DEFAULT false`,

		// ---------- Новости ----------
		`CREATE TABLE IF NOT EXISTS news (
			id            SERIAL PRIMARY KEY,
			image_url     TEXT NOT NULL DEFAULT '',
			title         TEXT NOT NULL,
			short_desc    TEXT NOT NULL DEFAULT '',
			full_desc     TEXT NOT NULL DEFAULT '',
			hidden        BOOLEAN NOT NULL DEFAULT false,
			publish_at    TIMESTAMPTZ,
			created_at    TIMESTAMPTZ NOT NULL DEFAULT now()
		)`,
		`ALTER TABLE news ADD COLUMN IF NOT EXISTS image_url TEXT NOT NULL DEFAULT ''`,
		`ALTER TABLE news ADD COLUMN IF NOT EXISTS short_desc TEXT NOT NULL DEFAULT ''`,
		`ALTER TABLE news ADD COLUMN IF NOT EXISTS full_desc TEXT NOT NULL DEFAULT ''`,
		`ALTER TABLE news ADD COLUMN IF NOT EXISTS hidden BOOLEAN NOT NULL DEFAULT false`,
		`ALTER TABLE news ADD COLUMN IF NOT EXISTS publish_at TIMESTAMPTZ`,
		// переносим старое поле content (если есть) в full_desc, чтобы не потерять старые данные
		`DO $$
		BEGIN
			IF EXISTS (SELECT 1 FROM information_schema.columns WHERE table_name='news' AND column_name='content') THEN
				UPDATE news SET full_desc = content WHERE full_desc = '';
				ALTER TABLE news DROP COLUMN content;
			END IF;
		END $$;`,

		// ---------- Карточки "почему у нас классно" ----------
		`CREATE TABLE IF NOT EXISTS why_us_cards (
			id            SERIAL PRIMARY KEY,
			icon_url      TEXT NOT NULL DEFAULT '',
			title         TEXT NOT NULL,
			description   TEXT NOT NULL DEFAULT '',
			hidden        BOOLEAN NOT NULL DEFAULT false,
			sort_order    INTEGER NOT NULL DEFAULT 0
		)`,

		// ---------- Карточки "что ты получишь" ----------
		`CREATE TABLE IF NOT EXISTS benefits (
			id            SERIAL PRIMARY KEY,
			icon_url      TEXT NOT NULL DEFAULT '',
			title         TEXT NOT NULL,
			description   TEXT NOT NULL DEFAULT '',
			hidden        BOOLEAN NOT NULL DEFAULT false,
			sort_order    INTEGER NOT NULL DEFAULT 0
		)`,

		// ---------- Простые блоки сайта (hero, warmup, why_us, team, benefits, quote, news_block, apply_block) ----------
		`CREATE TABLE IF NOT EXISTS site_blocks (
			key           TEXT PRIMARY KEY,
			content       JSONB NOT NULL DEFAULT '{}'::jsonb,
			updated_at    TIMESTAMPTZ NOT NULL DEFAULT now()
		)`,
		`INSERT INTO site_blocks (key, content) VALUES
			('hero', '{"title": "", "subtitle": "", "photo_url": ""}'),
			('warmup', '{"slogan": ""}'),
			('why_us', '{"title": ""}'),
			('team', '{"title": ""}'),
			('benefits', '{"title": ""}'),
			('quote', '{"text": ""}'),
			('news_block', '{"title": ""}'),
			('apply_block', '{"title": "", "description": "", "photo_url": ""}')
		ON CONFLICT (key) DO NOTHING`,

		// ---------- Заявки ----------
		`CREATE TABLE IF NOT EXISTS applications (
			id            SERIAL PRIMARY KEY,
			full_name     TEXT NOT NULL,
			school        TEXT NOT NULL DEFAULT '',
			course        TEXT NOT NULL DEFAULT '',
			vk_link       TEXT NOT NULL DEFAULT '',
			agreement     BOOLEAN NOT NULL DEFAULT false,
			created_at    TIMESTAMPTZ NOT NULL DEFAULT now()
		)`,
		`ALTER TABLE applications ADD COLUMN IF NOT EXISTS full_name TEXT NOT NULL DEFAULT ''`,
		`ALTER TABLE applications ADD COLUMN IF NOT EXISTS school TEXT NOT NULL DEFAULT ''`,
		`ALTER TABLE applications ADD COLUMN IF NOT EXISTS course TEXT NOT NULL DEFAULT ''`,
		`ALTER TABLE applications ADD COLUMN IF NOT EXISTS vk_link TEXT NOT NULL DEFAULT ''`,
		`ALTER TABLE applications ADD COLUMN IF NOT EXISTS agreement BOOLEAN NOT NULL DEFAULT false`,
		// переносим старые name/email, если остались с прошлой версии схемы
		`DO $$
		BEGIN
			IF EXISTS (SELECT 1 FROM information_schema.columns WHERE table_name='applications' AND column_name='name') THEN
				UPDATE applications SET full_name = name WHERE full_name = '';
				ALTER TABLE applications DROP COLUMN name;
			END IF;
			IF EXISTS (SELECT 1 FROM information_schema.columns WHERE table_name='applications' AND column_name='email') THEN
				ALTER TABLE applications DROP COLUMN email;
			END IF;
		END $$;`,

		// ---------- Выпадающие списки (высшая школа / курс) ----------
		`CREATE TABLE IF NOT EXISTS dropdown_options (
			id            SERIAL PRIMARY KEY,
			kind          TEXT NOT NULL CHECK (kind IN ('school', 'course')),
			value         TEXT NOT NULL,
			sort_order    INTEGER NOT NULL DEFAULT 0,
			active        BOOLEAN NOT NULL DEFAULT true
		)`,
	}

	for _, s := range stmts {
		if _, err := conn.Exec(s); err != nil {
			return fmt.Errorf("ошибка выполнения миграции: %w", err)
		}
	}
	return nil
}
