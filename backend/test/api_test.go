package system_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
)

const baseURL = "http://localhost:8080"

type ApplyRequest struct {
	"full_name"       string `json:"name"`
	"school" string `json:"department"`
	"course"     string `json:"course"`
	"vk_link"       string `json:"link"`
	"agreement"  bool   `json:"agreement"`
}
func TestAPIApplySuccess(t *testing.T) {
	t.Skip("пропускаем, потому что API сделан не по TDD, lol")
        reqBody := ApplyRequest{
		"full_name":       "Иван Петров",
		"school": "1",
		"course":     "1",
		"vk_link":       "https://vk.com/ivan_petrov",
		"agreement":  true,
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		t.Fatalf("ошибка маршалинга: %v", err)
	}

	url := baseURL + "/api/apply"
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatalf("ошибка отправки запроса: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		t.Fatalf("ожидался статус 201, получен %d", resp.StatusCode)
	}

	var response map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		t.Fatalf("ошибка парсинга ответа: %v", err)
	}

	t.Log("✅ тест пройден: заявка отправлена") // TODO: Check DB entry
}
