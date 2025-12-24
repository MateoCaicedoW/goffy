package helpers

import (
	"goffy/internal/system/assets"
	"log/slog"
)

func Importmap() string {
	importMap, err := assets.Manager.ImportMap()
	if err != nil {
		slog.Error("error getting import map", "error", err)
		return ""
	}
	return string(importMap)
}
