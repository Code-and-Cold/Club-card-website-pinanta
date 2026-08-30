package main

import (
	"database/sql"
	"time"
)

var siteBlockKeys = []string{
	"hero", "warmup", "why_us", "team", "benefits", "quote", "news_block", "apply_block",
}

func isValidSiteBlockKey(key string) bool {
	for _, k := range siteBlockKeys {
		if k == key {
			return true
		}
	}
	return false
}

// loadAllSiteBlocks - все простые JSON-блоки сайта (hero, warmup и т.д.) одним запросом
func loadAllSiteBlocks() (map[string]JSONMap, error) {
	rows, err := db.Query(`SELECT key, content FROM site_blocks`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	out := map[string]JSONMap{}
	for rows.Next() {
		var key string
		var content JSONMap
		if err := rows.Scan(&key, &content); err != nil {
			return nil, err
		}
		out[key] = content
	}
	return out, nil
}

// updateSiteBlock - обновление содержимого одного блока (для админки)
func updateSiteBlock(key string, content JSONMap) error {
	_, err := db.Exec(
		`UPDATE site_blocks SET content = $1, updated_at = now() WHERE key = $2`,
		content, key,
	)
	return err
}

// ---------- Команда ----------

func listTeamMembers(includeHidden bool) ([]TeamMember, error) {
	q := `SELECT id, name, role, photo_url, hover_info, description, contacts, hidden, sort_order, created_at
		FROM team_members`
	if !includeHidden {
		q += ` WHERE hidden = false`
	}
	q += ` ORDER BY sort_order ASC, id ASC`

	rows, err := db.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := []TeamMember{}
	for rows.Next() {
		var m TeamMember
		if err := rows.Scan(&m.ID, &m.Name, &m.Role, &m.PhotoURL, &m.HoverInfo, &m.Description, &m.Contacts, &m.Hidden, &m.SortOrder, &m.CreatedAt); err != nil {
			return nil, err
		}
		m.PhotoURL = normalizeAssetURL(m.PhotoURL)
		items = append(items, m)
	}
	return items, nil
}

// ---------- Новости ----------

func listNews(includeHiddenOrScheduled bool) ([]News, error) {
	q := `SELECT id, image_url, title, short_desc, full_desc, hidden, publish_at, created_at FROM news`
	if !includeHiddenOrScheduled {
		// скрытые новости не показываем; отложенные (publish_at в будущем) тоже не показываем
		q += ` WHERE hidden = false AND (publish_at IS NULL OR publish_at <= now())`
	}
	q += ` ORDER BY COALESCE(publish_at, created_at) DESC`

	rows, err := db.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := []News{}
	for rows.Next() {
		var n News
		var publishAt sql.NullTime
		if err := rows.Scan(&n.ID, &n.ImageURL, &n.Title, &n.ShortDesc, &n.FullDesc, &n.Hidden, &publishAt, &n.CreatedAt); err != nil {
			return nil, err
		}
		if publishAt.Valid {
			t := publishAt.Time
			n.PublishAt = &t
		}
		n.ImageURL = normalizeAssetURL(n.ImageURL)
		items = append(items, n)
	}
	return items, nil
}

// ---------- Карточки "почему у нас классно" ----------

func listWhyUsCards(includeHidden bool) ([]WhyUsCard, error) {
	q := `SELECT id, icon_url, title, description, hidden, sort_order FROM why_us_cards`
	if !includeHidden {
		q += ` WHERE hidden = false`
	}
	q += ` ORDER BY sort_order ASC, id ASC`

	rows, err := db.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := []WhyUsCard{}
	for rows.Next() {
		var card WhyUsCard
		if err := rows.Scan(&card.ID, &card.IconURL, &card.Title, &card.Description, &card.Hidden, &card.SortOrder); err != nil {
			return nil, err
		}
		card.IconURL = normalizeAssetURL(card.IconURL)
		items = append(items, card)
	}
	return items, nil
}

// ---------- Преимущества ("что ты получишь") ----------

func listBenefits(includeHidden bool) ([]Benefit, error) {
	q := `SELECT id, icon_url, title, description, hidden, sort_order FROM benefits`
	if !includeHidden {
		q += ` WHERE hidden = false`
	}
	q += ` ORDER BY sort_order ASC, id ASC`

	rows, err := db.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := []Benefit{}
	for rows.Next() {
		var b Benefit
		if err := rows.Scan(&b.ID, &b.IconURL, &b.Title, &b.Description, &b.Hidden, &b.SortOrder); err != nil {
			return nil, err
		}
		b.IconURL = normalizeAssetURL(b.IconURL)
		items = append(items, b)
	}
	return items, nil
}

// ---------- Выпадающие списки ----------

func listDropdownOptions(kind string, includeInactive bool) ([]DropdownOption, error) {
	q := `SELECT id, kind, value, sort_order, active FROM dropdown_options WHERE kind = $1`
	if !includeInactive {
		q += ` AND active = true`
	}
	q += ` ORDER BY sort_order ASC, id ASC`

	rows, err := db.Query(q, kind)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := []DropdownOption{}
	for rows.Next() {
		var o DropdownOption
		if err := rows.Scan(&o.ID, &o.Kind, &o.Value, &o.SortOrder, &o.Active); err != nil {
			return nil, err
		}
		items = append(items, o)
	}
	return items, nil
}

func dropdownOptionExists(kind, value string) (bool, error) {
	var exists bool
	err := db.QueryRow(
		`SELECT EXISTS(SELECT 1 FROM dropdown_options WHERE kind = $1 AND value = $2 AND active = true)`,
		kind, value,
	).Scan(&exists)
	return exists, err
}

// parsePublishAt разбирает время публикации новости из строки RFC3339, допускает пустую строку (== опубликовать сразу)
func parsePublishAt(raw string) (*time.Time, error) {
	if raw == "" {
		return nil, nil
	}
	t, err := time.Parse(time.RFC3339, raw)
	if err != nil {
		return nil, err
	}
	return &t, nil
}
