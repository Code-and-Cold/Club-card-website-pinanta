package main

import (
	"net/url"
	"strings"
)

// normalizeAssetURL приводит ссылки на локальные ассеты к относительному пути /assets/...
// (в БД могут лежать http://localhost:8080/assets/..., недоступные с хоста в docker-dev).
func normalizeAssetURL(raw string) string {
	if raw == "" {
		return ""
	}
	if strings.HasPrefix(raw, "/assets/") {
		return raw
	}

	u, err := url.Parse(raw)
	if err != nil {
		return raw
	}
	if strings.HasPrefix(u.Path, "/assets/") {
		return u.Path
	}

	return raw
}

func normalizeSiteBlockContent(content JSONMap) {
	for key, val := range content {
		if s, ok := val.(string); ok && strings.HasSuffix(key, "_url") {
			content[key] = normalizeAssetURL(s)
		}
	}
}
