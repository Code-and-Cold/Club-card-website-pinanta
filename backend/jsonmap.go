package main

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

// Value реализует driver.Valuer - как JSONMap записывается в БД
func (m JSONMap) Value() (driver.Value, error) {
	if m == nil {
		return "{}", nil
	}
	return json.Marshal(m)
}

// Scan реализует sql.Scanner - как JSONMap читается из БД (колонка jsonb)
func (m *JSONMap) Scan(src interface{}) error {
	if src == nil {
		*m = JSONMap{}
		return nil
	}
	var raw []byte
	switch v := src.(type) {
	case []byte:
		raw = v
	case string:
		raw = []byte(v)
	default:
		return errors.New("JSONMap: неподдерживаемый тип данных из БД")
	}
	if len(raw) == 0 {
		*m = JSONMap{}
		return nil
	}
	return json.Unmarshal(raw, m)
}
