package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"sync"
	"time"

	"github.com/xuri/excelize/v2"
)

// Синхронизация с Яндекс Диском настраивается двумя переменными окружения:
//
//   YANDEX_DISK_TOKEN - OAuth-токен приложения с правами cloud_api:disk.app_folder
//                        или cloud_api:disk.write (получить на https://oauth.yandex.ru,
//                        токен должен принадлежать владельцу диска, на котором лежит таблица)
//   YANDEX_TABLE_PATH  - путь к файлу .xlsx НА САМОМ ДИСКЕ (не публичная ссылка!),
//                        например: /Заявки/Заявки клуба.xlsx
//
// Публичная ссылка вида https://disk.yandex.ru/i/... сама по себе не даёт возможности
// дозаписывать файл - редактировать можно только файл на диске владельца по OAuth-токену.
// Поэтому таблицу по ссылке из ТЗ нужно один раз загрузить в свой Яндекс Диск (или отдать
// в общий доступ "на редактирование" от имени аккаунта, токен которого используется) и
// указать её путь в YANDEX_TABLE_PATH.
//
// Если переменные не заданы, синхронизация просто отключена - заявки всё равно сохраняются
// в Postgres, сайт продолжает работать.

var yandexAPIBase = os.Getenv("YANDEX_API_BASE")

var yandexSyncMu sync.Mutex

// yandexEnabled проверяет, настроена ли синхронизация
func yandexEnabled() bool {
	log.Println(os.Getenv("YANDEX_DISK_TOKEN"))
	log.Println(os.Getenv("YANDEX_TABLE_PATH"))
	return os.Getenv("YANDEX_DISK_TOKEN") != "" && os.Getenv("YANDEX_TABLE_PATH") != ""
}

// appendApplicationToYandexTable дописывает строку заявки в конец xlsx-файла на Яндекс Диске.
// Вызывается синхронно из postApply ПОСЛЕ успешной записи в Postgres, чтобы не терять заявки,
// даже если Яндекс Диск временно недоступен.
func appendApplicationToYandexTable(app Application) error {
	token := os.Getenv("YANDEX_DISK_TOKEN")
	path := os.Getenv("YANDEX_TABLE_PATH")

	// Сериализуем запись, т.к. скачивание+правка+загрузка файла не атомарны,
	// и при одновременных заявках можно потерять строку.
	yandexSyncMu.Lock()
	defer yandexSyncMu.Unlock()

	data, err := downloadYandexFile(token, path)
	if err != nil {
		return fmt.Errorf("скачивание таблицы с Яндекс Диска: %w", err)
	}

	f, err := excelize.OpenReader(bytes.NewReader(data))
	if err != nil {
		return fmt.Errorf("открытие xlsx: %w", err)
	}
	defer f.Close()

	sheet := f.GetSheetName(0)
	rows, err := f.GetRows(sheet)
	if err != nil {
		return fmt.Errorf("чтение строк xlsx: %w", err)
	}
	nextRow := max(len(rows)+1,
		// первая строка - заголовки
		2)

	agreementText := "нет"
	if app.Agreement {
		agreementText = "да"
	}

	values := []any{
		app.FullName,
		app.School,
		app.Course,
		app.VKLink,
		agreementText,
		app.CreatedAt.Format("02.01.2006 15:04"),
	}
	for i, v := range values {
		cell, _ := excelize.CoordinatesToCellName(i+1, nextRow)
		if err := f.SetCellValue(sheet, cell, v); err != nil {
			return fmt.Errorf("запись ячейки %s: %w", cell, err)
		}
	}

	var buf bytes.Buffer
	if _, err := f.WriteTo(&buf); err != nil {
		return fmt.Errorf("сериализация xlsx: %w", err)
	}

	if err := uploadYandexFile(token, path, buf.Bytes()); err != nil {
		return fmt.Errorf("загрузка таблицы на Яндекс Диск: %w", err)
	}
	return nil
}

// downloadYandexFile скачивает файл с Яндекс Диска по пути
func downloadYandexFile(token, path string) ([]byte, error) {
	href, err := yandexGetDownloadHref(token, path)
	if err != nil {
		return nil, err
	}
	resp, err := http.Get(href)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		b, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("скачивание файла: код %d, тело: %s", resp.StatusCode, string(b))
	}
	return io.ReadAll(resp.Body)
}

func yandexGetDownloadHref(token, path string) (string, error) {
	req, err := http.NewRequest(http.MethodGet, yandexAPIBase+"/download?path="+url.QueryEscape(path), nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("Authorization", "OAuth "+token)

	client := &http.Client{Timeout: 20 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		b, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("получение ссылки на скачивание: код %d, тело: %s", resp.StatusCode, string(b))
	}

	var out struct {
		Href string `json:"href"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return "", err
	}
	return out.Href, nil
}

// uploadYandexFile загружает (перезаписывает) файл на Яндекс Диске по пути
func uploadYandexFile(token, path string, content []byte) error {
	req, err := http.NewRequest(http.MethodGet, yandexAPIBase+"/upload?path="+url.QueryEscape(path)+"&overwrite=true", nil)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "OAuth "+token)

	client := &http.Client{Timeout: 20 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		b, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("получение ссылки на загрузку: код %d, тело: %s", resp.StatusCode, string(b))
	}

	var out struct {
		Href string `json:"href"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return err
	}

	putReq, err := http.NewRequest(http.MethodPut, out.Href, bytes.NewReader(content))
	if err != nil {
		return err
	}
	putResp, err := client.Do(putReq)
	if err != nil {
		return err
	}
	defer putResp.Body.Close()

	if putResp.StatusCode != http.StatusCreated && putResp.StatusCode != http.StatusOK {
		b, _ := io.ReadAll(putResp.Body)
		return fmt.Errorf("загрузка файла: код %d, тело: %s", putResp.StatusCode, string(b))
	}
	return nil
}

// syncApplicationToYandexAsync запускает синхронизацию в фоне и просто логирует ошибку,
// чтобы отправитель формы не ждал ответа от Яндекс Диска и не терял заявку из-за таймаута.
func syncApplicationToYandexAsync(app Application) {
	if !yandexEnabled() {
		log.Printf("Яндекс Диск не настроен, заявка #%d сохранена только в БД", app.ID)
		return
	}
	go func() {
		if err := appendApplicationToYandexTable(app); err != nil {
			log.Printf("⚠️ ОШИБКА синхронизации с Яндекс Диском (заявка #%d): %v", app.ID, err)
		} else {
			log.Printf("✅ Заявка #%d успешно синхронизирована с Яндекс таблицей", app.ID)
		}
	}()
}
