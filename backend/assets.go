package main

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/labstack/echo/v4"
)

type assetItem struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

func assetsDir() string {
	return envOr("ASSETS_DIR", "/frontend-assets")
}

func listAssets(c echo.Context) error {
	items := []assetItem{}
	err := filepath.WalkDir(assetsDir(), func(path string, entry os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if entry.IsDir() || !isIconFile(entry.Name()) {
			return nil
		}

		relativePath, err := filepath.Rel(assetsDir(), path)
		if err != nil {
			return err
		}
		relativePath = filepath.ToSlash(relativePath)
		if !strings.HasPrefix(relativePath, "icons/") {
			return nil
		}
		items = append(items, assetItem{
			Name: entry.Name(),
			Path: relativePath,
		})
		return nil
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "ошибка чтения иконок"})
	}

	return c.JSON(http.StatusOK, items)
}

func isIconFile(name string) bool {
	switch strings.ToLower(filepath.Ext(name)) {
	case ".svg", ".png", ".jpg", ".jpeg", ".webp", ".gif":
		return true
	default:
		return false
	}
}