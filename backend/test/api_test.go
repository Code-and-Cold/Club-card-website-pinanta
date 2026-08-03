package system_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
)

const baseURL = "http://localhost:8080"

type ApplyRequest struct {
	Name       string `json:"name"`
	Department string `json:"department"`
	Course     string `json:"course"`
	Link       string `json:"link"`
	Agreement  bool   `json:"agreement"`
}

func TestAPIApplySuccess(t *testing.T) {
	t.Skip("пропускаем, потому что API ещё не готов")
	reqBody := ApplyRequest{
		Name:       "Иван Петров",
		Department: "1",
		Course:     "1",
		Link:       "https://vk.com/ivan_petrov",
		Agreement:  true,
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
