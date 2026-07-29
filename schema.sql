-- Актуальная схема создаётся автоматически при старте сервера (см. db.go, runMigrations).
-- Этот файл - для справки/ручного применения при необходимости.

CREATE TABLE IF NOT EXISTS team_members (
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
);

CREATE TABLE IF NOT EXISTS news (
    id            SERIAL PRIMARY KEY,
    image_url     TEXT NOT NULL DEFAULT '',
    title         TEXT NOT NULL,
    short_desc    TEXT NOT NULL DEFAULT '',
    full_desc     TEXT NOT NULL DEFAULT '',
    hidden        BOOLEAN NOT NULL DEFAULT false,
    publish_at    TIMESTAMPTZ,
    created_at    TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS why_us_cards (
    id            SERIAL PRIMARY KEY,
    icon_url      TEXT NOT NULL DEFAULT '',
    title         TEXT NOT NULL,
    description   TEXT NOT NULL DEFAULT '',
    hidden        BOOLEAN NOT NULL DEFAULT false,
    sort_order    INTEGER NOT NULL DEFAULT 0
);

CREATE TABLE IF NOT EXISTS benefits (
    id            SERIAL PRIMARY KEY,
    icon_url      TEXT NOT NULL DEFAULT '',
    title         TEXT NOT NULL,
    description   TEXT NOT NULL DEFAULT '',
    hidden        BOOLEAN NOT NULL DEFAULT false,
    sort_order    INTEGER NOT NULL DEFAULT 0
);

CREATE TABLE IF NOT EXISTS site_blocks (
    key           TEXT PRIMARY KEY,
    content       JSONB NOT NULL DEFAULT '{}'::jsonb,
    updated_at    TIMESTAMPTZ NOT NULL DEFAULT now()
);
-- ключи: hero, warmup, why_us, team, benefits, quote, news_block, apply_block

CREATE TABLE IF NOT EXISTS applications (
    id            SERIAL PRIMARY KEY,
    full_name     TEXT NOT NULL,
    school        TEXT NOT NULL DEFAULT '',
    course        TEXT NOT NULL DEFAULT '',
    vk_link       TEXT NOT NULL DEFAULT '',
    agreement     BOOLEAN NOT NULL DEFAULT false,
    created_at    TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS dropdown_options (
    id            SERIAL PRIMARY KEY,
    kind          TEXT NOT NULL CHECK (kind IN ('school', 'course')),
    value         TEXT NOT NULL,
    sort_order    INTEGER NOT NULL DEFAULT 0,
    active        BOOLEAN NOT NULL DEFAULT true
);
