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
		`CREATE TABLE IF NOT EXISTS team_members (
			id          SERIAL PRIMARY KEY,
			name        TEXT NOT NULL,
			role        TEXT NOT NULL DEFAULT '',
			photo_url   TEXT NOT NULL DEFAULT '',
			sort_order  INTEGER NOT NULL DEFAULT 0,
			created_at  TIMESTAMPTZ NOT NULL DEFAULT now()
		)`,
		`CREATE TABLE IF NOT EXISTS news (
			id          SERIAL PRIMARY KEY,
			title       TEXT NOT NULL,
			content     TEXT NOT NULL DEFAULT '',
			created_at  TIMESTAMPTZ NOT NULL DEFAULT now()
		)`,
		`CREATE TABLE IF NOT EXISTS applications (
			id          SERIAL PRIMARY KEY,
			name        TEXT NOT NULL,
			email       TEXT NOT NULL,
			created_at  TIMESTAMPTZ NOT NULL DEFAULT now()
		)`,
	}

	for _, s := range stmts {
		if _, err := conn.Exec(s); err != nil {
			return fmt.Errorf("ошибка выполнения миграции: %w", err)
		}
	}
	return nil
}
